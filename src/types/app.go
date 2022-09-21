package types

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppContext struct {
	DB *gorm.DB
	Tokens *Tokens
	Server *fiber.App
	Validator *validator.Validate
}

type DbConnectionOptions struct {
	Host string
	User string
	Password string
	DbName string
	Port string
	SslMode bool
	DisableLogger bool
}