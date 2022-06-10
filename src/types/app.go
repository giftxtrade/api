package types

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AppContext struct {
	DB *gorm.DB
	Tokens *Tokens
	Router *mux.Router
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