package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

const AUTH_REQ string = "authorization required"
const AUTH_KEY types.AuthKeyType = "auth"
const AUTH_HEADER string = "Authorization"

type Auth struct {
	User database.User `json:"user"`
	Token string `json:"token"`
}

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
	
	auth := ParseAuthContext(c.UserContext())
	if !auth.User.Admin {
		return utils.FailResponseUnauthorized(c, "access for admin users only")
	}
	return c.Next()
}

func (ctx Controller) authenticate_user(c *fiber.Ctx) error {
	authorization := c.Get(AUTH_HEADER)
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
	id_raw, email := claims["id"].(string), claims["email"].(string)
	id, err := strconv.ParseInt(id_raw, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid claim id")
	}
	user, err := ctx.Querier.FindUserByIdAndEmail(c.Context(), database.FindUserByIdAndEmailParams{
		ID: id,
		Email: email,
	})
	if err != nil {
		return err
	}
	c.SetUserContext(context.WithValue(c.UserContext(), AUTH_KEY, Auth{
		Token: raw_token,
		User: user,
	}))
	return nil
}

// Generates a JWT with claims, signed with key
func GenerateJWT(key string, user *database.User) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": fmt.Sprint(user.ID),
		"name": user.Name,
		"email": user.Email,
		"imageUrl": user.ImageUrl,
	})
	token, err := jwt.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Given a context, find and return the auth struct using the types.AuthKey key
func ParseAuthContext(context context.Context) Auth {
	auth := context.Value(AUTH_KEY).(Auth)
	return auth
}
