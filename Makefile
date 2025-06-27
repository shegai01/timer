.phony: run

build:
	go build .

run: build
	./main

PORT=8080
stop:
	fuser -k $(PORT)/tcp

docker_db:
	docker exec -it timerdb psql -U alex01 
