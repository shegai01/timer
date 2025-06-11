.phony: run

build:
	go build main.go

run: build
	./main

PORT=8080
stop:
	@fuser -k $(PORT)/tcp

# restart: build run