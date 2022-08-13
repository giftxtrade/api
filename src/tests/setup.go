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

	if err = app.AutoMigrate(db); err != nil {
		t.Fatal("migration failed", err)
	}

	callback(db)
	return db
}

func SetupMockController(db *gorm.DB) controllers.Controller {
	return controllers.Controller{
		AppContext: types.AppContext{
			DB: db,
			Tokens: &types.Tokens{
				JwtKey: "my-secret-jwt-token",
			},
		},
		Service: services.New(db),
	}
}

func SetupMockUserService(t *testing.T) *gorm.DB {
	db := MockMigration(t, func(db *gorm.DB) {
		db.Exec("delete from users")
	})
	if db.Error != nil {
		t.FailNow()
	}
	return db
}

func SetupMockCategoryService(t *testing.T) *gorm.DB {
	db := MockMigration(t, func(db *gorm.DB) {
		db.Exec("delete from categories")
	})
	if db.Error != nil {
		t.FailNow()
	}
	return db
}

func SetupMockProductService(t *testing.T) *gorm.DB {
	db := MockMigration(t, func(db *gorm.DB) {
		db.Exec("delete from categories, products")
	})
	if db.Error != nil {
		t.FailNow()
	}
	return db
}

func SetupMockEventService(t *testing.T) *gorm.DB {
	db := MockMigration(t, func(db *gorm.DB) {
		db.Exec("delete from users, events")
	})
	if db.Error != nil {
		t.FailNow()
	}
	return db
}