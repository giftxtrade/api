package types

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppContext struct {
	DB *gorm.DB
	Tokens *Tokens
	Server *fiber.App
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