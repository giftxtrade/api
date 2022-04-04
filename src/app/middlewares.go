package app

import (
	"context"
	"net/http"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/google/uuid"
)

// Authentication middleware. Saves user data in request context within types.AuthKey key
func UseJwtAuth(app *AppBase, next http.Handler) http.Handler {
	const AUTH_REQ string = "authorization required"

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("Authorization")
			// Parse bearer token
			raw_token, err := utils.GetBearerToken(authorization)
			if err != nil {
				utils.FailResponseUnauthorized(w, "authorization required")
				return
			}

			// Parse JWT
			claims, err := utils.GetJwtClaims(raw_token, app.Tokens.JwtKey)
			if err != nil {
				utils.FailResponseUnauthorized(w, "authorization required")
				return
			}

			// Get user from id, username, email
			user := services.GetUserByIdAndEmail(app.DB, claims["id"].(string), claims["email"].(string))
			if user.ID == uuid.Nil {
				utils.FailResponseUnauthorized(w, "authorization required")
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
func UseAdminOnly(app *AppBase, next http.Handler) http.Handler {
	return UseJwtAuth(
		app, 
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				user := r.Context().Value(types.AuthKey).(types.User)
				if !user.IsAdmin {
					utils.FailResponseUnauthorized(w, "access for admin users only")
					return
				}
				next.ServeHTTP(w, r)
			},
		),
	)
}