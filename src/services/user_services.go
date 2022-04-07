package services

import (
	"github.com/giftxtrade/api/src/types"
)

type UserService struct {
	Service
}

// Find user by either the id or email
func (service *UserService) Find(key string) types.User {
	var user types.User
	service.DB.
		Table(service.TABLE).
		Where("id = ? OR email = ?", key, key).
		First(&user)
	return user
}

func (service *UserService) FindByEmail(email string) types.User {
	var user types.User
	service.DB.
		Table(service.TABLE).
		Where("email = ?", email).
		First(&user)
	return user
}

func (service *UserService) FindById(id string) types.User {
	var user types.User
	service.DB.
		Table(service.TABLE).
		Where("id = ?", id).
		First(&user)
	return user
}

func (service *UserService) FindByIdAndEmail(id string, email string) types.User {
	var user types.User
	service.DB.
		Table(service.TABLE).
		Where("id = ? AND email = ?", id, email).
		First(&user)
	return user
}

func (service *UserService) FindByIdOrEmail(id string, email string) types.User {
	var user types.User
	service.DB.
		Table(service.TABLE).
		Where("id = ? OR email = ?", id, email).
		First(user)
	return user
}

func (service *UserService) FindOrCreate(create_user *types.CreateUser) types.User {
	user := service.FindByEmail(create_user.Email)
	if user == (types.User{}) {
		user = service.Create(create_user)
	}
	return user
}

func (service *UserService) Create(create_user *types.CreateUser) types.User {
	user := types.User{
		Name: create_user.Name,
		Email: create_user.Email,
		ImageUrl: create_user.ImageUrl,
	}
	service.DB.
		Table(service.TABLE).
		Create(&user)
	return user
}