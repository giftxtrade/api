package services

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Find user by either the id or email
func GetUser(db *gorm.DB, key string) types.User {
	var user types.User
	db.Table("users").Where("id = ? OR email = ?", key, key).First(&user)
	return user
}

func GetUserByEmail(db *gorm.DB, email string) types.User {
	var user types.User
	db.Table("users").Where("email = ?", email).First(&user)
	return user
}

func GetUserById(db *gorm.DB, id uuid.UUID) types.User {
	var user types.User
	db.Table("users").Where("id = ?", id).First(&user)
	return user
}

func GetUserByIdAndEmail(db *gorm.DB, id uuid.UUID, email string) types.User {
	var user types.User
	db.Table("users").Where("id = ? AND email = ?", id.String(), email).First(&user)
	return user
}

func GetUserByIdOrEmail(db *gorm.DB, id uuid.UUID, email string) types.User {
	var user types.User
	db.Table("users").Where("id = ? OR email = ?", id.String, email).First(user)
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