package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/giftxtrade/api/src/types"
	"github.com/golang-jwt/jwt"
)

func DbConfig() (types.DbConnection, error) {
	db_config_file_data, err := ioutil.ReadFile("db_config.json")
	if err != nil {
		return types.DbConnection{}, errors.New("db_config.json not found")
	}

	var db_config types.DbConnection
	err = json.Unmarshal([]byte(db_config_file_data), &db_config)
	if err != nil {
		return types.DbConnection{}, err
	}
	return db_config, nil
}

func ParseTokens() (types.Tokens, error) {
	tokens_file_data, err := ioutil.ReadFile("tokens.json")
	if err != nil {
		return types.Tokens{}, errors.New("tokens.json not found")
	}
	tokens := types.Tokens{}
	err = json.Unmarshal([]byte(tokens_file_data), &tokens)
	if err != nil {
		return types.Tokens{}, err
	}
	return tokens, nil
}

// Given a bearer token ("Bearer <TOKEN>") returns the token or an error if parsing was unsuccessful
func GetBearerToken(authorization string) (string, error) {
	parsed_authorization := strings.Split(authorization, " ")
	if parsed_authorization[0] != "Bearer" || len(parsed_authorization) < 2 {
		return "", fmt.Errorf("could not parse bearer token")
	}
	token := parsed_authorization[1]
	return token, nil
}

// Given a raw jwt token and an encryption key return the mapped jwt claims or an error
func GetJwtClaims(jwt_token string, key string) (jwt.MapClaims, error) {
	token, token_err := jwt.Parse(jwt_token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(key), nil
	})
	if token_err != nil {
		return nil, fmt.Errorf("could not parse jwt token")
	}
	
	// Get claims stored in parsed JWT token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not fetch jwt claims")
	}
	return claims, nil
}

// Generates a JWT with claims, signed with key
func GenerateJWT(key string, user *types.User) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"username": user.Username,
		"email": user.Email,
		"is_active": user.IsActive,
	})
	token, err := jwt.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}