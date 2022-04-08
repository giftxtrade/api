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
	DB *gorm.DB
	Tokens types.Tokens
	UserServices services.UserService
	CategoryServices services.CategoryService
}

type IAppBase interface {
	NewBaseHandler(conn *sql.DB)
	CreateRoutes(router *mux.Router)
}

func (app *AppBase) NewBaseHandler(conn *gorm.DB) *AppBase {
	app.DB = conn
	app.UserServices = services.UserService{
		Service: services.Service{
			DB: conn,
			TABLE: "users",
		},
	}
	app.CategoryServices = services.CategoryService{
		Service: services.Service{
			DB: conn,
			TABLE: "categories",
		},
	}
	
	tokens, tokens_err := utils.ParseTokens()
	if tokens_err != nil {
		panic(tokens_err)
	}
	app.Tokens = tokens

	app.CreateSchemas() // create schemas
	app.SetupOauthProviders() // oauth providers
	return app
}

func New(conn *gorm.DB) *AppBase {
	app := AppBase{}
	return app.NewBaseHandler(conn)
}