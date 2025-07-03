.phony: build
build:
	go build main.go
run:
	./main
PORT=8080
stop:
	@fuser -k $(PORT)/tcp