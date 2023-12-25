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
func (s *UserService) FindOrCreate(ctx context.Context, input types.CreateUser) (database.User, bool, error) {
	user, err := s.Querier.FindUserByEmail(ctx, input.Email)
	if err != nil {
		user, err = s.Querier.CreateUser(ctx, s.CreateUserToCreateUserParams(input))
		
		if user.ID != 0 && err == nil {
			return user, true, nil
		}
		return database.User{}, false, err
	}
	return user, false, nil
}

func (s *UserService) CreateUserToCreateUserParams(input types.CreateUser) database.CreateUserParams {
	return database.CreateUserParams{
		Name: input.Name,
		Email: input.Email,
		ImageUrl: input.ImageUrl,
		Active: true,
		Admin: false,
	}
}

func (s *UserService) DbUserToUser(user database.User) types.User {
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
