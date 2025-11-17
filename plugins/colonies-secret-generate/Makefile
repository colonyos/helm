.DEFAULT: build
.PHONY: build test clean

build:
	@tinygo build -target wasip1 -buildmode=c-shared -o ./plugin.wasm plugin.go

test: build
	@go test -v ./tests/...

clean:
	rm ./plugin.wasm
