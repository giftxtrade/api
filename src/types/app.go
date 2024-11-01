package types

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AppContext struct {
	DB *sql.DB
	Tokens *Tokens
	Server *fiber.App
	Validator *validator.Validate
}
