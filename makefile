sqlc:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc generate

create-migration:
	go run github.com/ayaanqui/go-migration-tool --directory "./src/database/migrations" create-migration $(fileName)

docker-container:
	docker compose -f ./docker-compose.yml up

build:
	make sqlc && go build src/server.go

run:
	make sqlc && go run src/server.go

watch:
	go run github.com/cosmtrek/air -c .air.toml -- -h

test:
	make sqlc && go run github.com/rakyll/gotest -v ./src/tests

ts:
	go run github.com/gzuidhof/tygo generate

ts_publish:
	npm publish --access public
