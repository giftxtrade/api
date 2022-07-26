package app

import (
	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppBase struct {
	types.AppContext
	Service services.Service
}

type IAppBase interface {
	NewBaseHandler() *AppBase
	// database
	AutoMigrate(db *gorm.DB) error
	CreateSchemas() *AppBase
}

// Given app.AppBase.DB, and app.AppBase.Router
// creates db migrations, db services, oauth, and routes
func (app *AppBase) NewBaseHandler() *AppBase {
	tokens, tokens_err := utils.ParseTokens()
	if tokens_err != nil {
		panic(tokens_err)
	}
	app.Tokens = &tokens

	app.CreateSchemas() // create schemas
	app.Service = services.New(app.DB) // create services
	utils.SetupOauthProviders(tokens) // oauth providers
	controllers.New(app.AppContext, app.Service)
	return app
}

func New(conn *gorm.DB, server *fiber.App) *AppBase {
	app := AppBase{}
	app.DB = conn
	app.Server = server
	return app.NewBaseHandler()
}