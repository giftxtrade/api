package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func UseJwtAuth(app *AppBase, next http.Handler) http.Handler {
	const AUTH_REQ string = "Authorization required"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		parsed_authorization := strings.Split(authorization, " ")
		if parsed_authorization[0] != "Bearer" || len(parsed_authorization) < 2 {
			w.WriteHeader(401)
			utils.JsonResponse(w, types.Response{Message: AUTH_REQ})
			return
		}
		raw_token := parsed_authorization[1]

		// Parse JWT
		token, token_err := jwt.Parse(raw_token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(app.Tokens.JwtKey), nil
		})
		if token_err != nil {
			w.WriteHeader(401)
			utils.JsonResponse(w, types.Response{Message: AUTH_REQ})
			return
		}
		
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(401)
			utils.JsonResponse(w, types.Response{Message: "Could not fetch JWT claims"})
			return
		}

		// Get user from id, username, email
		user := types.User{}
		app.DB.Table("users").Find(
			&user, 
			"id = ? AND username = ? AND email = ?", 
			claims["id"],
			claims["username"],
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