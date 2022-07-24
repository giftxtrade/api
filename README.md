<p align="center">
    <a href="http://giftxtrade.com/" target="blank">
        <img src="https://raw.githubusercontent.com/giftxtrade/logos/master/logotype_rounded_color.svg" width='250' alt="GiftTrade Logo" />
    </a>
</p>

<p align="center">
    The platform that aims to simplify your online gift exchange and secret santa for <i>free</i>.
</p>

<br />

## Tech stack
- [PostgreSQL](https://www.postgresql.org/) - Primary database
- [GORM](https://gorm.io) - ORM to interact with the database programmatically
- [Postgres for GORM](https://github.com/go-gorm/postgres) - Postgres Driver for GORM
- [Google UUID](https://pkg.go.dev/github.com/google/uuid@v1.3.0) - Generates UUID before inserts
- [Fiber](https://github.com/gofiber/fiber) - Express.js inspired framework for Go (uses [Fasthttp](https://github.com/valyala/fasthttp))
- [Goth](https://github.com/markbates/goth) - OAuth support for multiple platforms
- [Goth Fiber](https://github.com/Shareed2k/goth_fiber) - Goth implementation for Fiber

## Instructions

### Set up config files

#### `db_config.json`
This project also uses a PostgreSQL database in order to run. To start, create a file called `db_config.json` in the project root and place the following in the file, replacing all content within the `[]` with the correct database values:

```json
{
    "host": "localhost",
    "dbName": "[database name]",
    "username": "[database username]",
    "password": "[database password]",
    "port": "5432"
}
```

***Note:*** All database table and column names are represented in `snake_case`. While all json field names are represented using `camelCase`.

#### `tokens.json`

In addition to the `db_config.json`, you will also need to create a `tokens.json` file which will hold the JWT secret, note that this token should be a randomly generated value and must not be made public. The `token.json` file should contain the following:
```json
{
    "jwtKey": "[YOUR SECRET TOKEN]",
    "twitter": {
        "apiKey": "[Twitter OAuth 1.0 API Key]",
        "apiKeySecret": "[Twitter OAuth 1.0 API Secret]",
        "bearerToken": "[Twitter OAuth Bearer Token]",
        "callbackUrl": "http://localhost:8080/auth/twitter/callback"
    },
    "google": {
        "clientId": "[Google Client Id]",
        "clientSecret": "[Google Secret Key]",
        "callbackUrl": "http://localhost:8080/auth/twitter/callback"
    }
}
```

### Generate Binary

```
$ make build
```

or 

```
$ go build src/server.go
```

Creates an executable binary file called `server`. To run this file call `./server`, like so:

```
$ ./server
```

This should start the server on port `8080`.

### Run without Binary

Another way to run the server is by using the `make run` command.

```
$ make run
```

Running the command should also start the server on port `8080`. This command is equivalent to running `go run src/server.go`.
