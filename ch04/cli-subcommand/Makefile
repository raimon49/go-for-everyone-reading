VERSION :=   $(shell git describe --tags --abbrev=0)
REVISION :=  $(shell git rev-parse --short HEAD)
GOVERSION := $(shell go version)
BUILDDATE := $(shell date '+%Y-%m-%d %H:%M:%S %Z')
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)' \
           -X 'main.goversion=$(GOVERSION)' \
           -X 'main.builddate=$(BUILDDATE)'
BIN_NAME := "todo"

.DEFAULT_GOAL:= build
.PHONY: setup
setup:
	go get "github.com/mitchellh/cli"

.PHONY: build
build: setup
	go build -ldflags "$(LDFLAGS)" -o "$(BIN_NAME)"

.PHONY: lint
lint:
	go vet main.go
	golint -set_exit_status main.go

