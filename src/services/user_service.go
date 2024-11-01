package services

import (
	"context"
	"fmt"

	"github.com/giftxtrade/api/src/database/jet/postgres/public/model"
	"github.com/giftxtrade/api/src/database/jet/postgres/public/table"
	"github.com/giftxtrade/api/src/types"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/golang-jwt/jwt"
)

type UserService struct {
	ServiceBase
}

func (s *UserService) FindUserByEmail(ctx context.Context, email string) (user types.User, err error) {
	qb := table.User.
		SELECT(table.User.AllColumns).
		WHERE(table.User.Email.EQ(postgres.String(email))).
		LIMIT(1)
	err = qb.QueryContext(ctx, s.DB, &user)
	return user, err
}

func (s *UserService) CreateUser(ctx context.Context, input types.CreateUser) (user types.User, err error) {
	qb := table.User.
		INSERT(
			table.User.Name,
			table.User.Email,
			table.User.ImageURL,
			table.User.Phone,
			table.User.Active,
			table.User.Admin,
		).MODEL(model.User{
			Name: input.Name,
			Email: input.Email,
			ImageURL: input.ImageUrl,
			Phone: &input.Phone,
			Active: false,
			Admin: false,
		}).
		RETURNING(table.User.AllColumns)
	qb.QueryContext(ctx, s.DB, &user)
	return user, err
}

// finds a user by email or creates one if not found. 
// boolean value is true if a new user is created, otherwise false
func (s *UserService) FindOrCreate(ctx context.Context, input types.CreateUser) (types.User, bool, error) {
	user, err := s.FindUserByEmail(ctx, input.Email)
	if err != nil {
		user, err = s.CreateUser(ctx, input)
		if err != nil {
			return types.User{}, false, err
		}
		return user, true, nil
	}
	return user, false, nil
}

// Generates a JWT with claims, signed with key
func (s *UserService) GenerateJWT(key string, user *types.User) (string, error) {
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
		User: user,
	}
	return auth, created, nil
}
