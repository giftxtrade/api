package app

import (
	"context"
	"net/http"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
)

// Authentication middleware. Saves user data in request context within types.AuthKey key
func (app *AppBase) UseJwtAuth(next http.Handler) http.Handler {
	const AUTH_REQ string = "authorization required"

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("Authorization")
			// Parse bearer token
			raw_token, err := utils.GetBearerToken(authorization)
			if err != nil {
				utils.FailResponseUnauthorized(w, AUTH_REQ)
				return
			}

			// Parse JWT
			claims, err := utils.GetJwtClaims(raw_token, app.Tokens.JwtKey)
			if err != nil {
				utils.FailResponseUnauthorized(w, AUTH_REQ)
				return
			}

			// Get user from id, username, email
			user := app.UserServices.FindByIdAndEmail(claims["id"].(string), claims["email"].(string))
			if user == (types.User{}) {
				utils.FailResponseUnauthorized(w, AUTH_REQ)
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
func (app *AppBase) UseAdminOnly(next http.Handler) http.Handler {
	return app.UseJwtAuth(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				auth := utils.ParseAuthContext(r.Context())
				if !auth.User.IsAdmin {
					utils.FailResponseUnauthorized(w, "access for admin users only")
					return
				}
				next.ServeHTTP(w, r)
			},
		),
	)
}