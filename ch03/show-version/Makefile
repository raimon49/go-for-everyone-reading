GIT_VER :=   $(shell git describe --tags)
.DEFAULT_GOAL:= build
.PHONY: build
build:
	go build -ldflags "-X main.version=${GIT_VER}" -o main
