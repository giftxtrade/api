package services

import (
	"github.com/giftxtrade/api/src/types"
	"gorm.io/gorm"
)

const TABLE string = "users"

// Find user by either the id or email
func GetUser(db *gorm.DB, key string) types.User {
	var user types.User
	db.Table(TABLE).Where("id = ? OR email = ?", key, key).First(&user)
	return user
}

func GetUserByEmail(db *gorm.DB, email string) types.User {
	var user types.User
	db.Table(TABLE).Where("email = ?", email).First(&user)
	return user
}

func GetUserById(db *gorm.DB, id string) types.User {
	var user types.User
	db.Table(TABLE).Where("id = ?", id).First(&user)
	return user
}

func GetUserByIdAndEmail(db *gorm.DB, id string, email string) types.User {
	var user types.User
	db.Table(TABLE).Where("id = ? AND email = ?", id, email).First(&user)
	return user
}

func GetUserByIdOrEmail(db *gorm.DB, id string, email string) types.User {
	var user types.User
	db.Table(TABLE).Where("id = ? OR email = ?", id, email).First(user)
	return user
}

func GetUserOrCreate(db *gorm.DB, create_user *types.CreateUser) types.User {
	user := GetUserByEmail(db, create_user.Email)
	if user == (types.User{}) {
		user = CreateUser(db, create_user)
	}
	return user
}

func CreateUser(db *gorm.DB, create_user *types.CreateUser) types.User {
	user := types.User{
		Name: create_user.Name,
		Email: create_user.Email,
		ImageUrl: create_user.ImageUrl,
	}
	db.Table(TABLE).Create(&user)
	return user
}