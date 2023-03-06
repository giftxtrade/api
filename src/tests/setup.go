package tests

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/gorm"

	"github.com/giftxtrade/api/src/app"
	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func NewMockDB(t *testing.T) (*gorm.DB, error) {
	test_db := os.Getenv("TEST_DB")
	if (test_db == "") {
		test_db = "giftxtrade_test_db"
	}
	db, err := utils.CreateDbConnection(types.DbConnectionOptions{
		Host: "localhost", 
		User: "postgres", 
		Password: "postgres", 
		DbName: test_db, 
		Port: "5432", 
		SslMode: false, 
		DisableLogger: true,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return nil, err
	}
	return db, nil
}

func MockMigration(t *testing.T) *gorm.DB {
	db, err := NewMockDB(t)
	if err != nil {
		t.FailNow()
	}
	return db
}

func New(t *testing.T) *app.AppBase {
	db := MockMigration(t)
	app := app.NewMock(db, fiber.New())
	err := db.Exec(`
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
	`).Error
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