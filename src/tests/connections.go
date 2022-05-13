package tests

import (
	"fmt"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/DATA-DOG/go-sqlmock"
)

func MockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db_conn, mock, err := sqlmock.New()
    if err != nil {
		return nil, mock, fmt.Errorf("failed to open mock sql db, got error: %v", err)
    }
    if db_conn == nil {
        return nil, mock, fmt.Errorf("mock db is null")
    }

    dialector := postgres.New(postgres.Config{
        DSN:                  "sqlmock_db_0",
        DriverName:           "postgres",
        Conn:                 db_conn,
        PreferSimpleProtocol: true,
    })
    db, err_gorm := gorm.Open(dialector, &gorm.Config{})
    if err_gorm != nil {
        return nil, mock, fmt.Errorf("failed to open gorm v2 db, got error: %v", err_gorm)
    }

    if db == nil {
        return nil, mock, fmt.Errorf("gorm db is null")
    }
	return db, mock, nil
}

func NewMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := MockDB()
	if (err != nil) {
		fmt.Print(err)
		t.FailNow()
	}
	return db, mock, nil
}