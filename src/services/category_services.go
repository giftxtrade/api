package services

import (
	"fmt"

	"github.com/giftxtrade/api/src/types"
)

type CategoryServices struct {
	Service
}

func (service *CategoryServices) Create(create_category *types.CreateCategory) (*types.Category, error) {
	if create_category.Name == "" {
		return nil, fmt.Errorf("at least CreateCategory.Name must be provided")
	}

	category := types.Category{
		Name: create_category.Name,
		Description: create_category.Description,
		Url: create_category.Url,
	}
	err := service.DB.
		Table(service.TABLE).
		Create(&category).
		Error
	return &category, err
}

func (service *CategoryServices) Find(name string) (*types.Category, error) {
	var category types.Category
	err := service.DB.
		Table(service.TABLE).
		Where("name = ?", name).
		First(&category).
		Error
	return &category, err
}

func (service *CategoryServices) FindAll() (*[]types.Category, error) {
	categories := new([]types.Category)
	err := service.DB.
		Table(service.TABLE).
		Find(categories).
		Error
	return categories, err
}

func (service *CategoryServices) FindOrCreate(name string) (*types.Category, error) {
	category, err := service.Find(name)
	if err != nil {
		category, err = service.Create(&types.CreateCategory{
			Name: name,
		})
		if err != nil {
			return nil, err
		}
	}
	return category, nil
}