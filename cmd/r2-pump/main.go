package main

import (
	"bytes"
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/pkg/errors"

	"github.com/cdnjs/tools/audit"
	"github.com/cdnjs/tools/gcp"
	"github.com/cdnjs/tools/metrics"
	"github.com/cdnjs/tools/packages"
	"github.com/cdnjs/tools/sentry"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"cloud.google.com/go/pubsub"
)

var (
	PROJECT      = os.Getenv("PROJECT")
	SUBSCRIPTION = os.Getenv("SUBSCRIPTION")

	R2_BUCKET     = os.Getenv("R2_BUCKET")
	R2_KEY_ID     = os.Getenv("R2_KEY_ID")
	R2_KEY_SECRET = os.Getenv("R2_KEY_SECRET")
	R2_ENDPOINT   = os.Getenv("R2_ENDPOINT")
)

func init() {
	sentry.Init()
}

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, PROJECT)
	if err != nil {
		log.Fatalf("could not create pubsub Client: %v", err)
	}
	sub := client.Subscription(SUBSCRIPTION)
	sub.ReceiveSettings.Synchronous = true
	sub.ReceiveSettings.MaxOutstandingMessages = 5
	sub.ReceiveSettings.NumGoroutines = runtime.NumCPU()

	for {
		log.Printf("started consuming messages\n")
		if err := consume(ctx, client, sub); err != nil {
			log.Fatalf("could not pull messages: %s", err)
		}
	}
}

type Message struct {
	GCSEvent gcp.GCSEvent `json:"gcsEvent"`
}

func consume(ctx context.Context, client *pubsub.Client, sub *pubsub.Subscription) error {
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		log.Printf("received message: %s\n", msg.Data)

		msg.Ack()
		if err := processMessage(ctx, msg.Data); err != nil {
			log.Printf("failed to process message: %s\n", err)
		}
	})
	if err != nil {
		return errors.Wrap(err, "could not receive from subscription")
	}
	return nil
}

func processMessage(ctx context.Context, data []byte) error {
	var message Message
	if err := json.Unmarshal(data, &message); err != nil {
		return errors.Wrap(err, "failed to parse")
	}

	return invoke(ctx, message.GCSEvent)
}

func invoke(ctx context.Context, e gcp.GCSEvent) error {
	sentry.Init()
	defer sentry.PanicHandler()

	pkgName := e.Metadata["package"].(string)
	version := e.Metadata["version"].(string)
	log.Printf("Invoke %s %s\n", pkgName, version)

	configStr, err := b64.StdEncoding.DecodeString(e.Metadata["config"].(string))
	if err != nil {
		return fmt.Errorf("could not decode config: %v", err)
	}

	archive, err := gcp.ReadObject(ctx, e.Bucket, e.Name)
	if err != nil {
		return fmt.Errorf("could not read object: %v", err)
	}

	bucket := aws.String(R2_BUCKET)

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: R2_ENDPOINT,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(R2_KEY_ID, R2_KEY_SECRET, "")),
	)
	if err != nil {
		return fmt.Errorf("could not load config: %s", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	keys := make([]string, 0)

	onFile := func(name string, r io.Reader) error {
		// remove leading slash
		name = name[1:]
		key := fmt.Sprintf("%s/%s/%s", pkgName, version, name)

		content, err := ioutil.ReadAll(r)
		if err != nil {
			return errors.Wrap(err, "could not read file")
		}

		keys = append(keys, key)

		meta := newMetadata(len(content))

		s3Object := s3.PutObjectInput{
			Body:     bytes.NewReader(content),
			Bucket:   bucket,
			Key:      aws.String(key),
			Metadata: meta,
		}
		if err := uploadFile(ctx, s3Client, &s3Object); err != nil {
			return errors.Wrap(err, "failed to upload file")
		}
		return nil
	}
	if err := gcp.Inflate(bytes.NewReader(archive), onFile); err != nil {
		return fmt.Errorf("could not inflate archive: %s", err)
	}

	if len(keys) == 0 {
		log.Printf("%s: no files to publish\n", pkgName)
	}

	pkg := new(packages.Package)
	if err := json.Unmarshal([]byte(configStr), &pkg); err != nil {
		return fmt.Errorf("failed to parse config: %s", err)
	}

	if err := audit.WroteR2(ctx, pkgName, version, keys); err != nil {
		log.Printf("failed to audit: %s\n", err)
	}
	if err := metrics.NewUpdatePublishedR2(); err != nil {
		return errors.Wrap(err, "could not report metrics")
	}

	return nil
}

func newMetadata(size int) map[string]string {
	lastModifiedTime := time.Now()
	lastModifiedSeconds := lastModifiedTime.UnixNano() / int64(time.Second)
	lastModifiedStr := lastModifiedTime.Format(http.TimeFormat)
	etag := fmt.Sprintf("%x-%x", lastModifiedSeconds, size)

	meta := make(map[string]string)

	// https://github.com/cdnjs/origin-worker/blob/ff91d30586c9e924ff919407401dff6f52826b4d/src/index.js#L212-L213
	meta["etag"] = etag
	meta["last_modified"] = lastModifiedStr

	return meta
}

func uploadFile(ctx context.Context, s3Client *s3.Client, obj *s3.PutObjectInput) error {
	if _, err := s3Client.PutObject(ctx, obj); err != nil {
		return errors.Wrapf(err, "failed to put Object %s", *obj.Key)
	}

	return nil
}
