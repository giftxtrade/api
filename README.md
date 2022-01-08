# 🖥 Go REST API

A simple REST API made using the standard Go library [`net/http`](https://pkg.go.dev/net/http) and PostgreSQL as the database. 

## Tech stack
- [PostgreSQL](https://www.postgresql.org/) - Primary database
- [GORM](https://gorm.io) - ORM to interact with the database programmatically
- [Postgres for GORM](https://github.com/go-gorm/postgres) - Postgres Driver for GORM
- [Google UUID](https://pkg.go.dev/github.com/google/uuid@v1.3.0) - Generates UUID before inserts
- [gorilla/mux](https://github.com/gorilla/mux) - Router built using the standard Go `http.Handler` interface
- [Goth](https://github.com/markbates/goth) - OAuth support for multiple platforms

## Instructions

### Set up config files

#### `db_config.json`
This project also uses a PostgreSQL database in order to run. To start, create a file called `db_config.json` in the project root and place the following in the file, replacing all content within the `[]` with the correct database values:

```json
{
    "db_name": "[database name]",
    "username": "[database username]",
    "password": "[database password]"
}
```

#### `tokens.json`

In addition to the `db_config.json`, you will also need to create a `tokens.json` file which will hold the JWT secret, note that this token should be a randomly generated value and must not be made public. The `token.json` file should contain the following:
```json
{
    "jwt_key": "[YOUR_SECRET_TOKEN]",
    "twitter": {
        "api_key": "[Twitter OAuth 1.0 API Key]",
        "api_key_secret": "[Twitter OAuth 1.0 API Secret]",
        "bearer_token": "[Twitter OAuth Bearer Token]",
        "access_token": "[Twitter OAuth Access Token]",
        "access_token_secret": "[Twitter OAuth Access Token Secret]"
    }
}
```

### Generate Binary

```
$ make build
```

or 

```
$ go build src/main.go
```

Creates an executable binary file called `main`. To run this file call `./main`, like so:

```
$ ./main
```

This should start the server on port `3001`.

### Run without Binary

Another way to run the server is by using the `make run` command.

```
$ make run
```

Running the command should also start the server on port `3001`. This command is equivalent to running `go run src/main.go`.
