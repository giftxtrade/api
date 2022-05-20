package tests

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/gorm"

	"github.com/giftxtrade/api/src/utils"
)

func NewMockDB(t *testing.T) (*gorm.DB, error) {
	test_db := os.Getenv("TEST_DB")
	if (test_db == "") {
		test_db = "giftxtrade_test_db"
	}
	db, err := utils.CreateDbConnection("localhost", "postgres", "password", test_db, "5432", false)
	if (err != nil) {
		fmt.Print(err)
		t.FailNow()
	}
	return db, nil
}