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
	UserServices *services.UserService
	CategoryServices *services.CategoryService
	ProductServices *services.ProductService
}

type IAppBase interface {
	NewBaseHandler(conn *sql.DB)
	CreateRoutes(router *mux.Router)
}

func (app *AppBase) NewBaseHandler(conn *gorm.DB, router *mux.Router) *AppBase {
	app.DB = conn
	app.Router = router
	app.UserServices = &services.UserService{
		Service: services.New(conn, "users"),
	}
	app.CategoryServices = &services.CategoryService{
		Service: services.New(conn, "categories"),
	}
	app.ProductServices = &services.ProductService{
		Service: services.New(conn, "products"),
		CategoryService: app.CategoryServices,
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