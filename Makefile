.phony: build run test restart stop
build:
	go build -o bin/app cmd/app/main.go
run:
	go run -race cmd/app/main.go
PORT=8080
stop:
	@fuser -k ${PORT}/tcp
restart: stop run
test:
	$(shell sh -c "for ((i = 0 ; i < 100 ; i++ )); do curl http://localhost:${PORT}/show > /dev/null &; done")
lint:
	golangci-lint run
