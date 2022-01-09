build:
	go build src/main.go

run:
	go run src/main.go

watch:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run src/main.go