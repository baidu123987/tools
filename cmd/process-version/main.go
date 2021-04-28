package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"cloud.google.com/go/pubsub"
	"github.com/pkg/errors"
)

var (
	PROJECT      = os.Getenv("PROJECT")
	SUBSCRIPTION = os.Getenv("SUBSCRIPTION")
	DOCKER_IMAGE = os.Getenv("DOCKER_IMAGE")
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, PROJECT)
	if err != nil {
		log.Fatalf("could not create pubsub Client: %v", err)
	}
	sub := client.Subscription(SUBSCRIPTION)

	for {
		if err := consume(client, sub); err != nil {
			log.Fatalf("could not pull messages: %s", err)
		}
	}
}

// FIXME: share with process-version function
type Message struct {
	OutgoingSignedURL string `json:"outgoingSignedURL"`
	Tar               string `json:"tar"`
	Pkg               string `json:"package"`
	Version           string `json:"version"`
}

func consume(client *pubsub.Client, sub *pubsub.Subscription) error {
	ctx := context.Background()
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		log.Printf("received message: %s\n", msg.Data)

		if err := processMessage(msg.Data); err != nil {
			log.Fatalf("failed to process message: %s", err)
		}
		msg.Ack()
	})
	if err != nil {
		return errors.Wrap(err, "could not receive from subscription")
	}
	return nil
}

func processMessage(data []byte) error {
	var message Message
	if err := json.Unmarshal(data, &message); err != nil {
		return errors.Wrap(err, "failed to parse")
	}

	inDir, outDir, err := setupSandbox()
	if err != nil {
		return errors.Wrap(err, "failed to setup sandbox")
	}

	if err := download(inDir, message.Tar); err != nil {
		return errors.Wrapf(err, "failed to download: %s", message.Tar)
	}

	if err := runSandbox(inDir, outDir); err != nil {
		return errors.Wrap(err, "failed to run sandbox")
	}

	if err := uploadToOutgoing(outDir, message.OutgoingSignedURL); err != nil {
		return errors.Wrap(err, "failed to upload to outgoing bucket")
	}

	defer os.RemoveAll(inDir)
	defer os.RemoveAll(outDir)
	return nil
}

func uploadToOutgoing(dir, url string) error {
	content := bytes.NewReader(bytes.Repeat([]byte("a"), 100))

	req, err := http.NewRequest("PUT", url, content)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "request failed")
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "could not read response body")
	}

	log.Printf("returned %s: %s\n", res.Status, string(bodyBytes))
	return nil
}

func download(dstDir string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dst, err := os.Create(path.Join(dstDir, "new-version.tgz"))
	if err != nil {
		return errors.Wrap(err, "could not write tmp file")
	}
	defer dst.Close()

	_, err = io.Copy(dst, resp.Body)
	return err
}
