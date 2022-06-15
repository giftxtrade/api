package app

import (
	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AppBase struct {
	types.AppContext
	UserService *services.UserService
	CategoryService *services.CategoryService
	ProductService *services.ProductService
}

type IAppBase interface {
	NewBaseHandler() *AppBase
	// database
	AutoMigrate(db *gorm.DB) error
	CreateSchemas() *AppBase
	// routes
	CreateRoutes() *AppBase
	// services
	CreateServices() *AppBase
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
	app.CreateServices() // create services
	utils.SetupOauthProviders(tokens) // oauth providers
	app.CreateRoutes()
	return app
}

func New(conn *gorm.DB, router *mux.Router) *AppBase {
	app := AppBase{}
	app.DB = conn
	app.Router = router
	return app.NewBaseHandler()
}