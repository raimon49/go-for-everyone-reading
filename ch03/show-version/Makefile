GIT_VER :=   $(shell git describe --tags)
.DEFAULT_GOAL:= build
.PHONY: build
build: setup
	go build -ldflags "-X main.version=${GIT_VER}" -o main

.PHONY: init
init:
	mkdir vendor
	touch vendor/.gitkeep
	go mod init main

.PHONY: setup
setup:
	go mod vendor
	touch vendor/.gitkeep
