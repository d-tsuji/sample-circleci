.PHONY: all build test lint clean deps devel-deps

BIN := sample-circleci
BUILD_LDFLAGS := "-s -w"
export GO111MODULE=on

all: clean build

deps:
	go mod tidy

devel-deps: deps
	go get -u \
	  github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0 \
	  github.com/reviewdog/reviewdog/cmd/reviewdog

build:
	go build -ldflags=$(BUILD_LDFLAGS) -o $(BIN)

test: deps
	go test -v -count=1 ./...

test-cover: deps
	go test -v -count=1 ./... -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html

lint: devel-deps
	go vet ./...
	$(GOBIN)/golint -set_exit_status ./...

clean:
	rm -rf $(BIN)
	go clean
