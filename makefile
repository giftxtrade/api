sqlc:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc generate

jet:
	go run github.com/go-jet/jet/v2/cmd/jet -dsn=postgresql://postgres:postgres@localhost:5433/postgres?sslmode=disable -path=./src/database/jet

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
	go run github.com/gzuidhof/tygo generate && go run ts_types/version.go

ts_publish:
	npm publish --access public
