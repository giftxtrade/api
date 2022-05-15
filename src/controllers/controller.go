package controllers

import (
	"context"
	"net/http"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gorilla/mux"
)

type Controller struct {
	types.AppContext
}

type IController interface {
	CreateController(router *mux.Router, path string)
}

// Authentication middleware. Saves user data in request context within types.AuthKey key
func (ctx *Controller) UseJwtAuth(next http.Handler) http.Handler {
	const AUTH_REQ string = "authorization required"
	user_services := services.UserService{
		Service: services.Service{
			DB: ctx.DB,
			TABLE: "users",
		},
	}

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
			claims, err := utils.GetJwtClaims(raw_token, ctx.Tokens.JwtKey)
			if err != nil {
				utils.FailResponseUnauthorized(w, AUTH_REQ)
				return
			}

			// Get user from id, username, email
			user, err := user_services.FindByIdAndEmail(claims["id"].(string), claims["email"].(string))
			if err != nil {
				utils.FailResponseUnauthorized(w, AUTH_REQ)
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), types.AuthKey, types.Auth{
				Token: raw_token,
				User: *user,
			}))
			// Serve handler with updated request
			next.ServeHTTP(w, r)
		},
	)
}

// Admin only access middleware (uses UseJwtAuth)
func (ctx *Controller) UseAdminOnly(next http.Handler) http.Handler {
	return ctx.UseJwtAuth(
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