package utils

import (
	"database/sql"
	"fmt"

	"github.com/giftxtrade/api/src/types"
)

func DbConnectionString(options types.DbConnectionOptions) string {
	sslmode_val := "enable"
	if !options.SslMode {
		sslmode_val = "disable"
	}
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Chicago", 
		options.Host,
		options.User,
		options.Password,
		options.DbName,
		options.Port,
		sslmode_val,
	)
	return dns
}

func DbConfig() (types.DbConnection, error) {
	var db_config types.DbConnection
	err := FileMapper("db_config.json", &db_config)
	return db_config, err
}

func CreateDbConnection(options types.DbConnectionOptions) (*sql.DB, error) {
	return sql.Open("postgres", DbConnectionString(options))
}

func NewDbConnection() (*sql.DB, error) {
	config, err := DbConfig()
	if err != nil {
		return nil, err
	}
	return CreateDbConnection(types.DbConnectionOptions{
		Host: config.Host, 
		User: config.Username, 
		Password: config.Password, 
		DbName: config.DbName, 
		Port: config.Port, 
		SslMode: false, // TODO: mark sslmode as true in production
		DisableLogger: false,
	})
}
