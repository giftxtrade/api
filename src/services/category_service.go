package services

import (
	"fmt"

	"github.com/giftxtrade/api/src/types"
)

type CategoryService struct {
	Service
}

func (service *CategoryService) Create(create_category *types.CreateCategory) (*types.Category, error) {
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

func (service *CategoryService) Find(name string) (*types.Category, error) {
	var category types.Category
	err := service.DB.
		Table(service.TABLE).
		Where("name = ?", name).
		First(&category).
		Error
	return &category, err
}

func (service *CategoryService) FindAll() (*[]types.Category, error) {
	categories := new([]types.Category)
	err := service.DB.
		Table(service.TABLE).
		Find(categories).
		Error
	return categories, err
}

// find or create a new category
// boolean value is true if a new user is created, otherwise false
func (service *CategoryService) FindOrCreate(name string) (*types.Category, bool, error) {
	category, err := service.Find(name)
	if err != nil {
		category, err = service.Create(&types.CreateCategory{
			Name: name,
		})
		if err != nil {
			return nil, false, err
		}
		return category, true, nil
	}
	return category, false, nil
}