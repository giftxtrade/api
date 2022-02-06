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
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Chicago", 
		config.Host,
		config.Username,
		config.Password,
		config.DbName,
		config.Port,
	)
	return gorm.Open(postgres.Open(dns), &gorm.Config{})
}
