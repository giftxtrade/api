package services

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/go-playground/validator/v10"
)

type CategoryService struct {
	ServiceBase
}

func (service *CategoryService) Create(create_category *types.CreateCategory, category *types.Category) error {
	validate := validator.New()
	if err := validate.Struct(create_category); err != nil {
		return err
	}
	
	category.Name = create_category.Name
	category.Description = create_category.Description
	category.Url = create_category.Url
	return service.DB.
		Table(service.TABLE).
		Create(category).
		Error
}

func (service *CategoryService) Find(name string, category *types.Category) error {
	return service.DB.
		Table(service.TABLE).
		Where("name = ?", name).
		First(category).
		Error
}

func (service *CategoryService) FindAll(categories []types.Category) error {
	return service.DB.
		Table(service.TABLE).
		Find(categories).
		Error
}

// find or create a new category
// boolean value is true if a new category is created, otherwise false
func (service *CategoryService) FindOrCreate(name string, category *types.Category) (bool, error) {
	if err := service.Find(name, category); err != nil {
		err = service.Create(&types.CreateCategory{
			Name: name,
		}, category)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}