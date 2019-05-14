VERSION :=   $(shell git describe --tags --abbrev=0)
REVISION :=  $(shell git rev-parse --short HEAD)
GOVERSION := $(shell go version)
BUILDDATE := $(shell date '+%Y/%m/%d %H:%M:%S %Z')
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)' \
           -X 'main.goversion=$(GOVERSION)' \
           -X 'main.builddate=$(BUILDDATE)'

.DEFAULT_GOAL:= build
.PHONY: build
build:
	go build -ldflags "$(LDFLAGS)" -o main
