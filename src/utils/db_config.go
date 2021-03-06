package utils

import (
	"fmt"

	"github.com/giftxtrade/api/src/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateDbConnection(options types.DbConnectionOptions) (*gorm.DB, error) {
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
	config := &gorm.Config{}
	if options.DisableLogger {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	return gorm.Open(postgres.Open(dns), config)
}

func NewDbConnection() (*gorm.DB, error) {
	config, err := DbConfig()
	if err != nil {
		return nil, err
	}
	// TODO: mark sslmode as true in production
	return CreateDbConnection(types.DbConnectionOptions{
		Host: config.Host, 
		User: config.Username, 
		Password: config.Password, 
		DbName: config.DbName, 
		Port: config.Port, 
		SslMode: false, 
		DisableLogger: true,
	})
}
