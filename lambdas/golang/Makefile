.PHONY: build clean

clean:
	@(echo Removing \`target\` dir)
	@(rm -rf ./target)

build: clean
	@(echo Building Golang lambda...)
	@(go build -v -o ./target/main main.go)
