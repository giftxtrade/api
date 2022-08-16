package services

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	ServiceBase
}

func (service *UserService) FindByEmail(email string, user *types.User) error {
	return service.DB.
		Table(service.TABLE).
		Where("email = ?", email).
		First(user).
		Error
}

func (service *UserService) FindById(id string, user *types.User) error {
	return service.DB.
		Table(service.TABLE).
		Where("id = ?", id).
		First(user).
		Error
}

func (service *UserService) FindByIdAndEmail(id string, email string, user *types.User) error {
	return service.DB.
		Table(service.TABLE).
		Where("id = ? AND email = ?", id, email).
		First(user).
		Error
}

func (service *UserService) FindByIdOrEmail(id string, email string, user *types.User) error {
	return service.DB.
		Table(service.TABLE).
		Where("id = ? OR email = ?", id, email).
		First(user).
		Error
}

// finds a user by email or creates one if not found. 
// boolean value is true if a new user is created, otherwise false
func (service *UserService) FindOrCreate(create_user *types.CreateUser, user *types.User) (bool, error) {
	if err := service.FindByEmail(create_user.Email, user); err != nil {
		if err = service.Create(create_user, user); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (service *UserService) Create(create_user *types.CreateUser, user *types.User) error {
	validate := validator.New()
	if err := validate.Struct(create_user); err != nil {
		return err
	}

	user.Name = create_user.Name
	user.Email = create_user.Email
	user.ImageUrl = create_user.ImageUrl
	return service.DB.
		Table(service.TABLE).
		Create(user).
		Error
}

func (service *UserService) DeleteById(key string) error {
	return service.DB.
		Table(service.TABLE).
		Where("id = ?", key).
		Delete(&types.User{}).
		Error
}