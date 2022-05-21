package services

import (
	"github.com/giftxtrade/api/src/types"
)

type UserService struct {
	Service
}

func (service *UserService) FindByEmail(email string) (*types.User, error) {
	var user types.User
	err := service.DB.
		Table(service.TABLE).
		Where("email = ?", email).
		First(&user).
		Error
	return &user, err
}

func (service *UserService) FindById(id string) (*types.User, error) {
	var user types.User
	err := service.DB.
		Table(service.TABLE).
		Where("id = ?", id).
		First(&user).
		Error
	return &user, err
}

func (service *UserService) FindByIdAndEmail(id string, email string) (*types.User, error) {
	var user types.User
	err := service.DB.
		Table(service.TABLE).
		Where("id = ? AND email = ?", id, email).
		First(&user).
		Error
	return &user, err
}

func (service *UserService) FindByIdOrEmail(id string, email string) (*types.User, error) {
	var user types.User
	err := service.DB.
		Table(service.TABLE).
		Where("id = ? OR email = ?", id, email).
		First(user).
		Error
	return &user, err
}

// finds a user by email or creates one if not found. 
// boolean value is true if a new user is created, otherwise false
func (service *UserService) FindOrCreate(create_user *types.CreateUser) (*types.User, bool, error) {
	found := false
	user, err := service.FindByEmail(create_user.Email)
	if *user == (types.User{}) {
		user, err = service.Create(create_user)
		if err == nil {
			found = true
		}
	}
	return user, found, err
}

func (service *UserService) Create(create_user *types.CreateUser) (*types.User, error) {
	user := types.User{
		Name: create_user.Name,
		Email: create_user.Email,
		ImageUrl: create_user.ImageUrl,
	}
	err := service.DB.
		Table(service.TABLE).
		Create(&user).
		Error
	return &user, err
}