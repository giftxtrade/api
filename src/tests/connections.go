package tests

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/gorm"

	"github.com/giftxtrade/api/src/app"
	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
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

func MockMigration(t *testing.T, callback func(db *gorm.DB)) *gorm.DB {
	db, err := NewMockDB(t)
	if err != nil {
		t.FailNow()
	}

	if app.AutoMigrate(db) != nil {
		t.Fatal("migration failed")
	}

	callback(db)
	return db
}

func SetupMockController(db *gorm.DB) *controllers.Controller {
	return &controllers.Controller{
		AppContext: types.AppContext{
			DB: db,
			Tokens: &types.Tokens{
				JwtKey: "my-secret-jwt-token",
			},
		},
	}
}

func SetupMockUserServices(t *testing.T) (*services.UserService) {
	db := MockMigration(t, func(db *gorm.DB) {
		db.Exec("delete from users")
	})

	return &services.UserService{
		Service: services.Service{
			DB: db,
			TABLE: "users",
		},
	}
}

func SetupMockCategoryServices(t *testing.T) (*services.CategoryServices) {
	db := MockMigration(t, func(db *gorm.DB) {
		db.Exec("delete from categories")
	})

	return &services.CategoryServices{
		Service: services.Service{
			DB: db,
			TABLE: "categories",
		},
	}
}