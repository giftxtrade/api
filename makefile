build:
	go build src/server.go

run:
	go run src/server.go

watch:
	nodemon --watch './src/**/*.go' -e go --signal SIGTERM --exec 'make' run

test:
	go test -v ./src/tests