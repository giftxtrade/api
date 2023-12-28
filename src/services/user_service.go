package services

import (
	"context"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/mappers"
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
		user, err = s.Querier.CreateUser(ctx, mappers.CreateUserToCreateUserParams(input))
		
		if user.ID != 0 && err == nil {
			return user, true, nil
		}
		return database.User{}, false, err
	}
	return user, false, nil
}
