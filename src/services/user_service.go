package services

import (
	"github.com/giftxtrade/api/src/types"
)

type UserService struct {
	ServiceBase
}

func (service *UserService) FindByEmail(email string, output *types.User) error {
	return service.DB.
		Table(service.TABLE).
		Where("email = ?", email).
		First(output).
		Error
}

func (service *UserService) FindById(id string, output *types.User) error {
	return service.DB.
		Table(service.TABLE).
		Where("id = ?", id).
		First(output).
		Error
}

func (service *UserService) FindByIdAndEmail(id string, email string, output *types.User) error {
	return service.DB.
		Table(service.TABLE).
		Where("id = ? AND email = ?", id, email).
		First(output).
		Error
}

func (service *UserService) FindByIdOrEmail(id string, email string, output *types.User) error {
	return service.DB.
		Table(service.TABLE).
		Where("id = ? OR email = ?", id, email).
		First(output).
		Error
}

// finds a user by email or creates one if not found. 
// boolean value is true if a new user is created, otherwise false
func (service *UserService) FindOrCreate(input *types.CreateUser, output *types.User) (bool, error) {
	if err := service.FindByEmail(input.Email, output); err != nil {
		if err = service.Create(input, output); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (service *UserService) Create(input *types.CreateUser, output *types.User) error {
	if err := service.Validator.Struct(input); err != nil {
		return err
	}

	output.Name = input.Name
	output.Email = input.Email
	output.ImageUrl = input.ImageUrl
	return service.DB.
		Table(service.TABLE).
		Create(output).
		Error
}

func (service *UserService) DeleteById(key string) error {
	return service.DB.
		Table(service.TABLE).
		Where("id = ?", key).
		Delete(&types.User{}).
		Error
}