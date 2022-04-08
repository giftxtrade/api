package services

import "github.com/giftxtrade/api/src/types"

type CategoryService struct {
	Service
}

func (service *CategoryService) Create(create_category *types.CreateCategory) types.Category {
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

func (service *CategoryService) Find(name string) types.Category {
	var category types.Category
	service.DB.
		Table(service.TABLE).
		Where("name = ?", name).
		First(&category)
	return category
}

func (service *CategoryService) FindAll() *[]types.Category {
	var categories []types.Category
	service.DB.
		Table(service.TABLE).
		Find(&categories)
	return &categories
}