package mappers

import (
	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
)

func CreateUserToCreateUserParams(input types.CreateUser) database.CreateUserParams {
	return database.CreateUserParams{
		Name: input.Name,
		Email: input.Email,
		ImageUrl: input.ImageUrl,
		Active: true,
		Admin: false,
	}
}

func DbUserToUser(user database.User) types.User {
	return types.User{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		ImageUrl: user.ImageUrl,
		Active: user.Active,
		Phone: user.Phone.String,
		Admin: user.Admin,
	}
}
