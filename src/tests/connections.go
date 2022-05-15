package tests

import (
	"fmt"
	"testing"

	"gorm.io/gorm"

	"github.com/giftxtrade/api/src/utils"
)

func NewMockDB(t *testing.T) (*gorm.DB, error) {
	db, err := utils.CreateDbConnection("localhost", "postgres", "password", "giftxtrade_test_db", "5432", false)
	if (err != nil) {
		fmt.Print(err)
		t.FailNow()
	}
	return db, nil
}