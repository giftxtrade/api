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

func GetUserOrCreate(db *gorm.DB, user *types.User) types.User {
	search_user := GetUserByIdOrEmail(db, user.ID, user.Email)
	if search_user == (types.User{}) {
		// create new user
		db.Table("users").Create(&user)
		return *user
	}
	return search_user
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