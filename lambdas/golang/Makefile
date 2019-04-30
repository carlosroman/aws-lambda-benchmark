.PHONY: build clean
GOOS ?= linux
GOARCH ?= amd64
CGO_ENABLED ?= 0 

clean:
	@(echo Removing \`target\` dir)
	@(rm -rf ./target)

build: clean
	@(echo Building Golang lambda...)
	@(GOOS=$(GOOS) \
		GOARCH=$(GOARCH) \
		CGO_ENABLED=$(CGO_ENABLED) \
		go build \
		-v \
		-o ./target/main main.go)

fmt:
	@go fmt

test:
	@go test \
		-v \
		-race \
		./...
