package services

import (
	"context"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
)

type UserService struct {
	ServiceBase
}

// finds a user by email or creates one if not found. 
// boolean value is true if a new user is created, otherwise false
func (service *UserService) FindOrCreate(ctx context.Context, input types.CreateUser) (database.User, bool, error) {
	user, err := service.Querier.FindUserByEmail(ctx, input.Email)
	if err != nil {
		user, err = service.Querier.CreateUser(ctx, database.CreateUserParams{
			Name: input.Name,
			Email: input.Email,
			ImageUrl: input.ImageUrl,
			Active: true,
			Admin: false,
		})
		
		if user.ID != 0 && err == nil {
			return user, true, nil
		}
		return database.User{}, false, err
	}
	return user, false, nil
}
