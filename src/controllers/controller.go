package controllers

import (
	"context"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
)

type Controller struct {
	types.AppContext
	Service services.Service
}

type IController interface {
	CreateController(router *mux.Router, path string)
}

// Authentication middleware. Saves user data in request context within types.AuthKey key
func (ctx *Controller) UseJwtAuth(c *fiber.Ctx) error {
	if err := ctx.authenticate_user(c); err != nil {
		return err
	}
	return c.Next()
}

// Admin only access middleware (uses UseJwtAuth)
func (ctx *Controller) UseAdminOnly(c *fiber.Ctx) error {
	if err := ctx.authenticate_user(c); err != nil {
		return err
	}
	
	auth := utils.ParseAuthContext(c.UserContext())
	if !auth.User.IsAdmin {
		return c.JSON(types.Errors{
			Errors: []string{"access for admin users only"},
		})
	}
	return c.Next()
}

func (ctx Controller) authenticate_user(c *fiber.Ctx) error {
	const AUTH_REQ string = "authorization required"

	authorization := c.Get(types.AuthHeader)
	// Parse bearer token
	raw_token, err := utils.GetBearerToken(authorization)
	if err != nil {
		return c.JSON(types.Errors{
			Errors: []string{AUTH_REQ},
		})
	}

	// Parse JWT
	claims, err := utils.GetJwtClaims(raw_token, ctx.Tokens.JwtKey)
	if err != nil {
		return c.JSON(types.Errors{
			Errors: []string{AUTH_REQ},
		})
	}

	// Get user from id, username, email
	var user types.User
	id, email := claims["id"].(string), claims["email"].(string)
	err = ctx.Service.UserService.FindByIdAndEmail(id, email, &user)
	if err != nil {
		return c.JSON(types.Errors{
			Errors: []string{AUTH_REQ},
		})
	}
	c.SetUserContext(context.WithValue(c.UserContext(), types.AuthKey, types.Auth{
		Token: raw_token,
		User: user,
	}))
	return nil
}

func New(app_ctx types.AppContext, service services.Service) Controller {
	controller := Controller{
		AppContext: app_ctx,
		Service: service,
	}
	server := app_ctx.Server

	// create routes
	server.Get("/", controller.Home)
	{
		auth := server.Group("/auth")
		profile := auth.Group("/profile")
		{
			profile.Use(controller.UseJwtAuth)
			profile.Get("", controller.GetProfile)
		}
		auth.Get("/:provider", controller.SignIn)
		auth.Get("/:provider/callback", controller.Callback)
	}
	return controller
}