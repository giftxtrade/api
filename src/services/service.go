package services

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ServiceBase struct {
	DB *gorm.DB
	TABLE string
	Validator *validator.Validate
}

type Service struct {
	DB *gorm.DB
	UserService UserService
	CategoryService CategoryService
	ProductService ProductService
	EventService EventService
	ParticipantService ParticipantService
}

type IService interface {
	CreateService(db *gorm.DB, table string) ServiceBase
	New(db *gorm.DB) Service
}

func CreateService(db *gorm.DB, table string, validator *validator.Validate) ServiceBase {
	return ServiceBase{
		DB: db,
		TABLE: table,
		Validator: validator,
	}
}

func New(db *gorm.DB, validator *validator.Validate) Service {
	service := Service{
		DB: db,
	}

	service.UserService = UserService{
		ServiceBase: CreateService(db, "users", validator),
	}
	service.CategoryService = CategoryService{
		ServiceBase: CreateService(db, "categories", validator),
	}
	service.ProductService = ProductService{
		ServiceBase: CreateService(db, "products", validator),
		CategoryService: service.CategoryService,
	}
	service.EventService = EventService{
		ServiceBase: CreateService(db, "events", validator),
		UserService: service.UserService,
	}
	service.ParticipantService = ParticipantService{
		ServiceBase: CreateService(db, "participants", validator),
		UserService: service.UserService,
	}
	return service
}