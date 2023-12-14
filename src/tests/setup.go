package tests

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/giftxtrade/api/src/app"
	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/database"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func NewMockDB(t *testing.T) *sql.DB {
	db, err := database.CreateDbConnection(database.DbConnection{
		Host: "localhost", 
		Username: "postgres", 
		Password: "postgres", 
		DbName: "postgres", 
		Port: 54322,
		SslMode: false,
	})
	if err != nil {
		fmt.Println("could not establish connection with test db", err)
		t.FailNow() 
		return nil
	}
	return db
}

func New(t *testing.T) *app.AppBase {
	db := NewMockDB(t)
	return app.NewMock(db, fiber.New())
}

func SetupMockController(app *app.AppBase) controllers.Controller {
	return controllers.Controller{
		AppContext: app.AppContext,
		Service: app.Service,
	}
}