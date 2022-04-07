package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/giftxtrade/api/src/types"
	"github.com/golang-jwt/jwt"
	"golang.org/x/net/context"
)

// Given a JSON file, map the contents into any struct dest
func FileMapper(filename string, dest interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("%s not found", filename)
	}
	if err = json.Unmarshal(file, dest); err != nil {
		return err
	}
	return nil
}

func DbConfig() (types.DbConnection, error) {
	var db_config types.DbConnection
	err := FileMapper("db_config.json", &db_config)
	return db_config, err
}

func ParseTokens() (types.Tokens, error) {
	var tokens types.Tokens
	err := FileMapper("tokens.json", &tokens)
	return tokens, err
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
		"name": user.Name,
		"email": user.Email,
		"image_url": user.ImageUrl,
	})
	token, err := jwt.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Given a context, find and return the auth struct using the types.AuthKey key
func ParseAuthContext(context context.Context) types.Auth {
	auth := context.Value(types.AuthKey).(types.Auth)
	return auth
}