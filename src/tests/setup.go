package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/giftxtrade/api/src/app"
	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
	"github.com/gofiber/fiber/v2"
)

func NewMockDB(t *testing.T) (*sql.DB, error) {
	test_db := os.Getenv("TEST_DB")
	test_password := "password"
	if (test_db == "") {
		test_db = "giftxtrade_test_db"
		test_password = "postgres"
	}
	db, err := database.CreateDbConnection(types.DbConnection{
		Host: "localhost", 
		Username: "postgres", 
		Password: test_password, 
		DbName: test_db, 
		Port: 5432, 
		SslMode: false,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow() 
		return nil, err
	}
	return db, nil
}

func MockMigration(t *testing.T) *sql.DB {
	db, err := NewMockDB(t)
	if err != nil {
		t.FailNow()
	}
	return db
}

func New(t *testing.T) *app.AppBase {
	db := MockMigration(t)
	app := app.NewMock(db, fiber.New())
	_, err := db.Exec(`
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
	`)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return nil
	}
	return app
}

func SetupMockController(app *app.AppBase) controllers.Controller {
	return controllers.Controller{
		AppContext: app.AppContext,
		Service: app.Service,
	}
}