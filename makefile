build:
	go build src/server.go

run:
	go run src/server.go

watch:
	go run github.com/go-playground/justdoit -build="make build" -run="./server"

test:
	go test -v ./src/tests