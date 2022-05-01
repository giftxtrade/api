package app

import (
	"log"

	"github.com/giftxtrade/api/src/types"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&types.User{},
		&types.Category{},
		&types.Product{},
		&types.Event{},
	)
}

func (app *AppBase) CreateSchemas() {
	if err := AutoMigrate(app.DB); err != nil {
		log.Fatal("Could not generate schema.\n")
		panic(err)
	}
}