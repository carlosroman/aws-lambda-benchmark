.PHONY: setup install build clean

clean:
	@rm -rf $(shell go env GOPATH)/src/lambda/target

setup:
	@echo "Installing deps version '${DEP_VERSION}'"
	@mkdir -p $(shell go env GOPATH)/bin
	@curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o $(shell go env GOPATH)/bin/dep
	@chmod +x $(shell go env GOPATH)/bin/dep

install:
	@(cd ./src/lambda; dep ensure)

build:
	@go build -v -o ./src/lambda/target/main lambda
