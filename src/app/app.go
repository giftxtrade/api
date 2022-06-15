package app

import (
	"database/sql"

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
	NewBaseHandler(conn *sql.DB)
	CreateRoutes(router *mux.Router)
}

func (app *AppBase) NewBaseHandler(conn *gorm.DB, router *mux.Router) *AppBase {
	app.DB = conn
	app.Router = router
	app.UserService = &services.UserService{
		Service: services.New(conn, "users"),
	}
	app.CategoryService = &services.CategoryService{
		Service: services.New(conn, "categories"),
	}
	app.ProductService = &services.ProductService{
		Service: services.New(conn, "products"),
		CategoryService: app.CategoryService,
	}
	
	tokens, tokens_err := utils.ParseTokens()
	if tokens_err != nil {
		panic(tokens_err)
	}
	app.Tokens = &tokens

	app.CreateSchemas() // create schemas
	utils.SetupOauthProviders(tokens) // oauth providers
	app.CreateRoutes()
	return app
}

func New(conn *gorm.DB, router *mux.Router) *AppBase {
	app := AppBase{}
	return app.NewBaseHandler(conn, router)
}