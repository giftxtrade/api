package services

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
	TABLE string
}

type IService interface {
	New(db *gorm.DB, table_name string) *Service
}

func New(db *gorm.DB, table_name string) *Service {
	return &Service{
		DB: db,
		TABLE: table_name,
	}
}