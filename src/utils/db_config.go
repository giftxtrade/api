package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDbConnection(host string, user string, password string, db_name string, port string, sslmode bool) (*gorm.DB, error) {
	sslmode_val := "enable"
	if !sslmode {
		sslmode_val = "disable"
	}
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Chicago", 
		host,
		user,
		password,
		db_name,
		port,
		sslmode_val,
	)
	return gorm.Open(postgres.Open(dns), &gorm.Config{})
}

func NewDbConnection() (*gorm.DB, error) {
	config, err := DbConfig()
	if err != nil {
		return nil, err
	}
	// TODO: mark sslmode as true in production
	return CreateDbConnection(config.Host, config.Username, config.Password, config.DbName, config.Port, false)
}
