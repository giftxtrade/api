package app

import (
	"context"
	"net/http"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/google/uuid"
)

func UseJwtAuth(app *AppBase, next http.Handler) http.Handler {
	const AUTH_REQ string = "Authorization required"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		// Parse bearer token
		raw_token, err := utils.GetBearerToken(authorization)
		if err != nil {
			w.WriteHeader(401)
			utils.JsonResponse(w, types.Response{Message: AUTH_REQ})
			return
		}

		// Parse JWT
		claims, err := utils.GetJwtClaims(raw_token, app.Tokens.JwtKey)
		if err != nil {
			w.WriteHeader(401)
			utils.JsonResponse(w, types.Response{Message: AUTH_REQ})
			return
		}

		// Get user from id, username, email
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

// Admin only access middleware (uses UseJwtAuth)
func UseAdminOnly(app *AppBase, next http.Handler) http.Handler {
	return UseJwtAuth(
		app, 
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				user := r.Context().Value(types.AuthKey).(types.User)
				if !user.IsAdmin {
					w.WriteHeader(401)
					utils.JsonResponse(w, types.Response{Message: "Access for admin users only"})
					return
				}
				next.ServeHTTP(w, r)
			},
		),
	)
}