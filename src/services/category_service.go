package services

import (
	"github.com/giftxtrade/api/src/types"
)

type CategoryService struct {
	ServiceBase
}

func (service *CategoryService) Create(input *types.CreateCategory, output *types.Category) error {
	if err := service.Validator.Struct(input); err != nil {
		return err
	}
	
	output.Name = input.Name
	output.Description = input.Description
	output.Url = input.Url
	return service.DB.
		Table(service.TABLE).
		Create(output).
		Error
}

func (service *CategoryService) Find(name string, output *types.Category) error {
	return service.DB.
		Table(service.TABLE).
		Where("categories.name = ?", name).
		First(output).
		Error
}

func (service *CategoryService) FindAll(output []types.Category) error {
	return service.DB.
		Table(service.TABLE).
		Find(output).
		Error
}

// find or create a new category
// boolean value is true if a new category is created, otherwise false
func (service *CategoryService) FindOrCreate(name string, output *types.Category) (bool, error) {
	if err := service.Find(name, output); err != nil {
		err = service.Create(&types.CreateCategory{
			Name: name,
		}, output)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}