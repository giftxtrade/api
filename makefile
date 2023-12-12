build:
	go build src/server.go

run:
	go run src/server.go

watch:
	go run github.com/go-playground/justdoit -build="make build" -run="./server"

test:
	go run github.com/rakyll/gotest -v ./src/tests

typegen:
	go run github.com/tkrajina/typescriptify-golang-structs/tscriptify -package=github.com/giftxtrade/api/src/types -target=typescript/types.ts -interface Product Event

sqlc:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc generate

create-migration:
	go run github.com/ayaanqui/go-migration-tool --directory "./src/database/migrations" create-migration $(fileName)
