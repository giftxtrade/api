package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect() (*gorm.DB, error) {
	config, err := DbConfig()
	if err != nil {
		return nil, err
	}
	dns := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Chicago", 
		config.Username, 
		config.Password, 
		config.DbName,
	)
	return gorm.Open(postgres.Open(dns), &gorm.Config{})
}