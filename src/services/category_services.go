package services

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/google/uuid"
)

type CategoryServices struct {
	Service
}

func (service *CategoryServices) Create(create_category *types.CreateCategory) types.Category {
	category := types.Category{
		Name: create_category.Name,
		Description: create_category.Description,
		Url: create_category.Url,
	}
	service.DB.
		Table(service.TABLE).
		Create(&category)
	return category
}

func (service *CategoryServices) Find(name string) types.Category {
	var category types.Category
	service.DB.
		Table(service.TABLE).
		Where("name = ?", name).
		First(&category)
	return category
}

func (service *CategoryServices) FindAll() *[]types.Category {
	var categories []types.Category
	service.DB.
		Table(service.TABLE).
		Find(&categories)
	return &categories
}

func (service *CategoryServices) FindOrCreate(name string) types.Category {
	category := service.Find(name)
	if category.ID == uuid.Nil {
		category = service.Create(&types.CreateCategory{
			Name: name,
		})
	}
	return category
}