.PHONY: build test

build:
	go build -o newznab ./cmd/main.go

test:
	go test ./...
