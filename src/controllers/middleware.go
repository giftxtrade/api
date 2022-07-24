package controllers

import (
	"context"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

const AUTH_REQ string = "authorization required"

// Authentication middleware. Saves user data in request context within types.AuthKey key
func (ctx *Controller) UseJwtAuth(c *fiber.Ctx) error {
	if err := ctx.authenticate_user(c); err != nil {
		return utils.FailResponseUnauthorized(c, AUTH_REQ)
	}
	return c.Next()
}

// Admin only access middleware (uses UseJwtAuth)
func (ctx *Controller) UseAdminOnly(c *fiber.Ctx) error {
	if err := ctx.authenticate_user(c); err != nil {
		return utils.FailResponseUnauthorized(c, AUTH_REQ)
	}
	
	auth := utils.ParseAuthContext(c.UserContext())
	if !auth.User.IsAdmin {
		return utils.FailResponseUnauthorized(c, "access for admin users only")
	}
	return c.Next()
}

func (ctx Controller) authenticate_user(c *fiber.Ctx) error {
	authorization := c.Get(types.AuthHeader)
	// Parse bearer token
	raw_token, err := utils.GetBearerToken(authorization)
	if err != nil {
		return err
	}

	// Parse JWT
	claims, err := utils.GetJwtClaims(raw_token, ctx.Tokens.JwtKey)
	if err != nil {
		return err
	}

	// Get user from id, username, email
	var user types.User
	id, email := claims["id"].(string), claims["email"].(string)
	err = ctx.Service.UserService.FindByIdAndEmail(id, email, &user)
	if err != nil {
		return err
	}
	c.SetUserContext(context.WithValue(c.UserContext(), types.AuthKey, types.Auth{
		Token: raw_token,
		User: user,
	}))
	return nil
}