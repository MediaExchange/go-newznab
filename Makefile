.PHONY: build test

build:
	go build -o newznab ./main/main.go

test:
	go test ./...
