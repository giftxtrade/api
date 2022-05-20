package types

import (
	"gorm.io/gorm"
)

type AppContext struct {
	DB *gorm.DB
	Tokens *Tokens
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