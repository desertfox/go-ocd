#Docker shortcuts
#Go shortcuts
PWD=$(shell pwd)

.PHONY:all
all: fmt lint build

.PHONY:build
build:
	go build .

.PHONY:fmt
fmt:
	gofmt -s -w .

.PHONY:lint
lint:
	golangci-lint run ./...

.PHONY:test
test:
	go test ./...

.PHONY:run
run:
	go run . 
