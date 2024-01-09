package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/giftxtrade/api/src/types"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

// Given a JSON file, map the contents into any struct dest
func FileMapper(filename string, dest interface{}) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("%s not found", filename)
	}
	if err = json.Unmarshal(file, dest); err != nil {
		return err
	}
	return nil
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

// Parses json serializable []byte into result data and validates the decoded result
func ParseAndValidateBody[T comparable](validator *validator.Validate, data []byte) (result T, error error) {
	if json.Unmarshal(data, &result) != nil {
		return result, fmt.Errorf("could not parse data")
	}
	if err := validator.Struct(result); err != nil {
		return result, err
	}
	return result, nil
}
