package types

import (
	"gorm.io/gorm"
)

type AppContext struct {
	DB *gorm.DB
	Tokens Tokens
}