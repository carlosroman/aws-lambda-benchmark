.PHONY: setup install build clean

clean:
	@rm -f $(shell go env GOPATH)/src/lambda/main

setup:
	@echo "Installing deps version '${DEP_VERSION}'"
	@mkdir -p $(shell go env GOPATH)/bin
	@curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o $(shell go env GOPATH)/bin/dep
	@chmod +x $(shell go env GOPATH)/bin/dep

install:
	@(cd $(shell go env GOPATH)/src/lambda; dep ensure)

build:
	@go build -o $(shell go env GOPATH)/src/lambda/main lambda
