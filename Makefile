GO_BUILD_ARGS = -mod=readonly -v -ldflags="-s -w"
CLOUD_FUNCTIONS = process-version check-pkg-updates kv-pump algolia-pump force-update r2-pump

define generate-func-make
	make -C ./functions/$1 $1.zip

endef

.PHONY: all
all: bin/process-version-host bin/git-sync bin/checker bin/r2-pump \
   ;$(foreach n,${CLOUD_FUNCTIONS},$(call generate-func-make,$n))

bin/checker:
	go build $(GO_BUILD_ARGS) -o bin/checker ./cmd/checker

bin/git-sync:
	go build $(GO_BUILD_ARGS) -o bin/git-sync ./cmd/git-sync

bin/process-version-host:
	go build $(GO_BUILD_ARGS) -o bin/process-version-host ./cmd/process-version-host

bin/process-version:
	go build $(GO_BUILD_ARGS) -o bin/process-version ./cmd/process-version

bin/r2-pump:
	go build $(GO_BUILD_ARGS) -o bin/r2-pump ./cmd/r2-pump

.PHONY: schema
schema:
	./bin/packages human > schema_human.json
	./bin/packages non-human > schema_non_human.json

.PHONY: clean
clean:
	rm -rfv bin/*
	rm -rfv functions/*/*.zip

.PHONY: test
test: clean bin/checker
	go test -v ./test/...

.PHONY: lint
lint:
	go get -u golang.org/x/lint/golint
	$(GOPATH)/bin/golint ./...

.PHONY: dev
dev: autoupdate
	docker build -t cdnjs-dev -f ./dev/Dockerfile .
	docker run -it cdnjs-dev

.PHONY: process-version-sandbox
process-version-sandbox:
	docker build -t $@ -f ./docker/process-version/Dockerfile .
