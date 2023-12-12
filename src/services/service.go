package services

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ServiceBase struct {
	DB *sql.DB
	Validator *validator.Validate
}

type Service struct {
	DB *sql.DB
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

func CreateService(db *sql.DB, validator *validator.Validate) ServiceBase {
	return ServiceBase{
		DB: db,
		Validator: validator,
	}
}

func New(db *sql.DB, validator *validator.Validate) Service {
	service := Service{
		DB: db,
	}

	service.UserService = UserService{
		ServiceBase: CreateService(db, validator),
	}
	service.CategoryService = CategoryService{
		ServiceBase: CreateService(db, validator),
	}
	service.ProductService = ProductService{
		ServiceBase: CreateService(db, validator),
		CategoryService: service.CategoryService,
	}
	service.ParticipantService = ParticipantService{
		ServiceBase: CreateService(db, validator),
		UserService: service.UserService,
	}
	service.EventService = EventService{
		ServiceBase: CreateService(db, validator),
		UserService: service.UserService,
		ParticipantService: service.ParticipantService,
	}
	return service
}