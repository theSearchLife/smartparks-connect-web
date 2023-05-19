.PHONY: all v build test lint

VERSION := $(shell git describe --tags --always --dirty="-dev")

all: build

v:
	@echo "Version: ${VERSION}"

build:
	go build -trimpath -ldflags "-s -X main.version=${VERSION}" -v

test:
	go test -race ./... 

lint:
	golangci-lint run

mod:
	go mod tidy
	git status # for some reason without this the next step shows some cached result.
	git diff-index HEAD
	git diff-index --quiet HEAD