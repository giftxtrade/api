package services

import (
	"context"
	"fmt"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/mappers"
	"github.com/giftxtrade/api/src/types"
	"github.com/golang-jwt/jwt"
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

// Generates a JWT with claims, signed with key
func (s *UserService) GenerateJWT(key string, user *database.User) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": fmt.Sprint(user.ID),
		"name": user.Name,
		"email": user.Email,
		"imageUrl": user.ImageUrl,
	})
	token, err := jwt.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserService) GenerateAuthUser(ctx context.Context, input types.CreateUser) (types.Auth, bool, error) {
	user, created, err := s.FindOrCreate(ctx, input)
	if err != nil {
		return types.Auth{}, false, fmt.Errorf("authentication could not succeed")
	}
	token, err := s.GenerateJWT(s.Tokens.JwtKey, &user)
	if err != nil {
		return types.Auth{}, false, fmt.Errorf("could not generate token")
	}
	auth := types.Auth{
		Token: token,
		User: mappers.DbUserToUser(user),
	}
	return auth, created, nil
}
