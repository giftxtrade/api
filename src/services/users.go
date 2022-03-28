package services

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserByEmail(db *gorm.DB, email string) types.User {
	var user types.User
	db.Table("users").Where("email = ?", email).First(&user)
	return user
}

func GetUserById(db *gorm.DB, id string) types.User {
	var user types.User
	db.Table("users").Where("id = ?", id).First(&user)
	return user
}

func GetUserByIdAndEmail(db *gorm.DB, id string, email string) types.User {
	var user types.User
	db.Table("users").Where("id = ? AND email = ?", id, email).First(&user)
	return user
}

func GetUserByEmailOrCreate(db *gorm.DB, email string, user *types.User) types.User {
	search_user := GetUserByEmail(db, email)
	if search_user.ID == uuid.Nil {
		// create new user
		db.Table("users").Create(user)
		return *user
	}
	return search_user
}