package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

const AUTH_REQ string = "authorization required"
const AUTH_KEY types.AuthKeyType = "auth"
const AUTH_HEADER string = "Authorization"
const EVENT_ID_PARAM_KEY types.EventIdParamKeyType = "EVENT_ID_PARAM"

// Authentication middleware. Saves user data in request user context with the `AUTH_KEY` key
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
	utils.SetUserContext(c, AUTH_KEY, types.Auth{
		Token: raw_token,
		User: types.User{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			ImageUrl: user.ImageUrl,
			Active: user.Active,
			Phone: user.Phone.String,
			Admin: user.Admin,
		},
	})
	return nil
}

// Given a context, find and return the auth struct using the types.AuthKey key
func ParseAuthContext(context context.Context) types.Auth {
	auth := context.Value(AUTH_KEY).(types.Auth)
	return auth
}

// Verifies if auth user is a valid participant of an event
// based on the URL param `:event_id`.
// 
// Saves the event_id (int64) in the request user context with the `EVENT_ID_PARAM_KEY` key
func (ctr *Controller) UseEventAuthWithParam(c *fiber.Ctx) error {
	auth_user := ParseAuthContext(c.UserContext())
	id_raw := c.Params("event_id")
	id, err := strconv.ParseInt(id_raw, 10, 64)
	if err != nil {
		return utils.FailResponseNotFound(c, "event not found")
	}

	event_id, err := ctr.Querier.FindEventForUser(c.Context(), database.FindEventForUserParams{
		EventID: id,
		UserID: sql.NullInt64{
			Valid: true,
			Int64: auth_user.User.ID,
		},
		Email: sql.NullString{
			Valid: true,
			String: auth_user.User.Email,
		},
	})
	if err != nil || event_id != id {
		return utils.FailResponseNotFound(c, "event not found")
	}
	utils.SetUserContext(c, EVENT_ID_PARAM_KEY, event_id)
	return c.Next()
}
