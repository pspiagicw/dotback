all: build test

build:
	go build .
test: 
	go test -v ./...
format:
	go format ./...
.PHONY: build test format all

