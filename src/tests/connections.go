package tests

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/gorm"

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