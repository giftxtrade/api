package utils

import (
	"context"
	"net/http"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
)

// Authentication middleware. Saves user data in request context within types.AuthKey key
func UseJwtAuth(jwt_key string, user_services *services.UserService, next http.Handler) http.Handler {
	const AUTH_REQ string = "authorization required"

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("Authorization")
			// Parse bearer token
			raw_token, err := GetBearerToken(authorization)
			if err != nil {
				FailResponseUnauthorized(w, AUTH_REQ)
				return
			}

			// Parse JWT
			claims, err := GetJwtClaims(raw_token, jwt_key)
			if err != nil {
				FailResponseUnauthorized(w, AUTH_REQ)
				return
			}

			// Get user from id, username, email
			user := user_services.FindByIdAndEmail(claims["id"].(string), claims["email"].(string))
			if user == (types.User{}) {
				FailResponseUnauthorized(w, AUTH_REQ)
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), types.AuthKey, types.Auth{
				Token: raw_token,
				User: user,
			}))
			// Serve handler with updated request
			next.ServeHTTP(w, r)
		},
	)
}

// Admin only access middleware (uses UseJwtAuth)
func UseAdminOnly(jwt_key string, user_services *services.UserService, next http.Handler) http.Handler {
	return UseJwtAuth(
		jwt_key, 
		user_services,
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				auth := ParseAuthContext(r.Context())
				if !auth.User.IsAdmin {
					FailResponseUnauthorized(w, "access for admin users only")
					return
				}
				next.ServeHTTP(w, r)
			},
		),
	)
}