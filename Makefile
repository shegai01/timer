.phony: run

build:
	go build main.go -o api_simple

run: build
	./main

PORT=8080
stop:
	fuser -k $(PORT)/tcp

