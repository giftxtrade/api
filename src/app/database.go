package app

import (
	"log"

	"github.com/giftxtrade/api/src/types"
)

func (app *AppBase) CreateSchemas() {
	err := app.DB.AutoMigrate(
		&types.User{},
	)
	if err != nil {
		log.Fatal("Could not generate schema.\n")
		panic(err)
	}
}