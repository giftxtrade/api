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
		Password: "password", 
		DbName: test_db, 
		Port: "5432", 
		SslMode: false, 
		DisableLogger: true,
	})
	if (err != nil) {
		fmt.Print(err)
		t.FailNow()
	}
	return db, nil
}

func MockMigration(t *testing.T) *gorm.DB {
	db, err := NewMockDB(t)
	if err != nil {
		t.FailNow()
	}

	db.Exec("drop table participants, events, users, products, categories")

	if err = app.AutoMigrate(db); err != nil {
		t.Fatal("migration failed", err)
	}
	return db
}

func New(t *testing.T) *app.AppBase {
	db := MockMigration(t)
	return app.NewMock(db, fiber.New())
}

func SetupMockController(app *app.AppBase) controllers.Controller {
	return controllers.Controller{
		AppContext: app.AppContext,
		Service: app.Service,
	}
}