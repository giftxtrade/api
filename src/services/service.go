package services

import "gorm.io/gorm"

type ServiceBase struct {
	DB *gorm.DB
	TABLE string
}

type Service struct {
	DB *gorm.DB
	UserService UserService
	CategoryService CategoryService
	ProductService ProductService
}

type IService interface {
	CreateService(db *gorm.DB, table string) ServiceBase
	New(db *gorm.DB) Service
}

func CreateService(db *gorm.DB, table string) ServiceBase {
	return ServiceBase{
		DB: db,
		TABLE: table,
	}
}

func New(db *gorm.DB) Service {
	service := Service{
		DB: db,
	}

	service.UserService = UserService{
		ServiceBase: CreateService(db, "users"),
	}
	service.CategoryService = CategoryService{
		ServiceBase: CreateService(db, "categories"),
	}
	service.ProductService = ProductService{
		ServiceBase: CreateService(db, "products"),
		CategoryService: service.CategoryService,
	}
	return service
}