package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
)

func DbConnectionString(options types.DbConnection) string {
	sslmode_val := "enable"
	if !options.SslMode {
		sslmode_val = "disable"
	}
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Chicago", 
		options.Host,
		options.Username,
		options.Password,
		options.DbName,
		strconv.Itoa(int(options.Port)),
		sslmode_val,
	)
	return dns
}

func DbConfig() (types.DbConnection, error) {
	var db_config types.DbConnection
	err := utils.FileMapper("db_config.json", &db_config)
	return db_config, err
}

func CreateDbConnection(options types.DbConnection) (*sql.DB, error) {
	return sql.Open("postgres", DbConnectionString(options))
}

func NewDbConnection() (*sql.DB, error) {
	config, err := DbConfig()
	if err != nil {
		return nil, err
	}
	return CreateDbConnection(config)
}
