.phony: build
build:
	go build -o bin/app cmd/app/main.go
run:
	go run -race cmd/app/main.go
PORT=8080
stop:
	@fuser -k ${PORT}/tcp
restart: stop run
test:
	for ((i = 0 ; i < 100 ; i++ )); do curl http://localhost:${PORT}/show > /dev/null &; done