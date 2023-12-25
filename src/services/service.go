package services

import (
	"database/sql"

	"github.com/giftxtrade/api/src/database"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ServiceBase struct {
	DB *sql.DB
	Querier *database.Queries
	Validator *validator.Validate
}

type Service struct {
	DB *sql.DB
	UserService UserService
	ProductService ProductService
	ParticipantService ParticipantService
	EventService EventService
}

type IService interface {
	CreateService(db *gorm.DB, table string) ServiceBase
	New(db *gorm.DB) Service
}

func New(db *sql.DB, querier *database.Queries, validator *validator.Validate) Service {
	service_base := ServiceBase {
		DB: db,
		Querier: querier,
		Validator: validator,
	}
	
	service := Service{
		DB: db,
	}
	service.UserService = UserService{
		ServiceBase: service_base,
	}
	service.ProductService = ProductService{
		ServiceBase: service_base,
	}
	service.ParticipantService = ParticipantService{
		ServiceBase: service_base,
	}
	service.EventService = EventService{
		ServiceBase: service_base,
		ParticipantService: service.ParticipantService,
	}
	return service
}