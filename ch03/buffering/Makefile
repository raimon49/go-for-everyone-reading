.DEFAULT_GOAL:= build
.PHONY: build
build: setup
	go build -o main

.PHONY: init
init:
	mkdir vendor
	touch vendor/.gitkeep
	go mod init main

.PHONY: setup
setup:
	go mod vendor
