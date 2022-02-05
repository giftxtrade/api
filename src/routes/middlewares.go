package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func UseJwtAuth(app *AppBase, next http.Handler) http.Handler {
	const AUTH_REQ string = "Authorization required"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		// Parse bearer token
		raw_token, err := get_bearer_token(authorization)
		if err != nil {
			w.WriteHeader(401)
			utils.JsonResponse(w, types.Response{Message: AUTH_REQ})
			return
		}

		// Parse JWT
		claims, err := get_jwt_claims(raw_token, app.Tokens.JwtKey)
		if err != nil {
			w.WriteHeader(401)
			utils.JsonResponse(w, types.Response{Message: AUTH_REQ})
			return
		}

		// Get user from claims
		user := types.User{}
		app.DB.Table("users").Find(
			&user, 
			"id = ? AND email = ?", 
			claims["id"],
			claims["email"],
		)
		if user.ID == uuid.Nil {
			w.WriteHeader(401)
			utils.JsonResponse(w, types.Response{Message: AUTH_REQ})
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), types.AuthKey, types.Auth{
			Token: raw_token,
			User: user,
		}))
		// Serve handler with updated request
		next.ServeHTTP(w, r)
	})
}

// Given a bearer token ("Bearer <TOKEN>") returns the token or an error if parsing was unsuccessful
func get_bearer_token(authorization string) (string, error) {
	parsed_authorization := strings.Split(authorization, " ")
	if parsed_authorization[0] != "Bearer" || len(parsed_authorization) < 2 {
		return "", fmt.Errorf("could not parse bearer token")
	}
	token := parsed_authorization[1]
	return token, nil
}

// Given a raw jwt token and an encryption key return the mapped jwt claims or an error
func get_jwt_claims(jwt_token string, key string) (jwt.MapClaims, error) {
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

func AdminOnly(app *AppBase, next http.Handler) http.Handler {
	return UseJwtAuth(app, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user := r.Context().Value(types.AuthKey)
		// fmt.Println(user)
		next.ServeHTTP(w, r)
	}))
}