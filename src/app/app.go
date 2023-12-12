package app

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/ayaanqui/go-migration-tool/migration_tool"
	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppBase struct {
	types.AppContext
	Service services.Service
	MigrationDirectory string
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
	app.Validator = validator.New()
	app.Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	m := migration_tool.New(app.DB, &migration_tool.Config{
		TableName: "migration",
		Directory: app.MigrationDirectory,
	})
	m.RunMigration()

	app.Service = services.New(app.DB, app.Validator) // create services
	utils.SetupOauthProviders(*app.Tokens) // oauth providers
	controllers.New(app.AppContext, app.Service)
	return app
}

func New(conn *sql.DB, server *fiber.App) *AppBase {
	app := AppBase{}
	app.DB = conn
	app.Server = server
	// initialize tokens
	tokens, tokens_err := utils.ParseTokens()
	if tokens_err != nil {
		panic(tokens_err)
	}
	app.Tokens = &tokens
	app.MigrationDirectory = "./migrations"
	return app.NewBaseHandler()
}

func NewMock(conn *sql.DB, server *fiber.App) *AppBase {
	app := AppBase{}
	app.DB = conn
	app.Server = server
	app.Tokens = &types.Tokens{
		JwtKey: "my-secret-jwt-token",
	}
	app.MigrationDirectory = "../../migrations"
	return app.NewBaseHandler()
}