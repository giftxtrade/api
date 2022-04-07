package services

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
	TABLE string
}