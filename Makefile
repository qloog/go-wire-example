SHELL := /bin/bash
BASEDIR = $(shell pwd)

SERVICE_NAME := go-wire-example

# Detect OS and Architecture
OS := $(shell uname -s)
ARCH := $(shell uname -m)

.PHONY: all
all: run

.PHONY: init
init:
	@go install github.com/google/wire/cmd/wire@latest

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: build
build:
	@echo "Detected OS: $(OS)"
	@echo "Detected ARCH: $(ARCH)"
ifeq ($(OS),Linux)
	@echo "Building for Linux..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/$(SERVICE_NAME) cmd/main.go cmd/wire_gen.go
else ifeq ($(OS),Darwin)
ifeq ($(ARCH),x86_64)
	@echo "Building for macOS (Intel)..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o bin/$(SERVICE_NAME) cmd/main.go cmd/wire_gen.go
else ifeq ($(ARCH),arm64)
	@echo "Building for macOS (M-series)..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -o bin/$(SERVICE_NAME) cmd/main.go cmd/wire_gen.go
else
	@echo "Unsupported macOS architecture: $(ARCH)"
	exit 1
endif
else
	@echo "Unsupported OS: $(OS)"
	exit 1
endif

.PHONY: run
# make run, run current project
run: build
	./bin/$(SERVICE_NAME)

.PHONY: test
# make test, run test
test:
	go test ./...

.PHONY: clean
# make clean, clean bin dir
clean:
	rm -rf bin/*

.PHONY: wire
# make wire, run wire
wire:
	wire ./...

.PHONY: help
# make help, show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m  %-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := all