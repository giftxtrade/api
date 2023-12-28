package controllers

import (
	"github.com/giftxtrade/api/src/mappers"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
)

// [GET] /auth/profile (authentication required)
func (ctx Controller) GetProfile(c *fiber.Ctx) error {
	auth := ParseAuthContext(c.UserContext())
	return utils.DataResponse(c, auth)
}

// [GET] /auth/:provider
func (ctx Controller) SignIn(c *fiber.Ctx) error {
	return goth_fiber.BeginAuthHandler(c)
}

// [GET] /auth/:provider/callback
func (ctx Controller) Callback(c *fiber.Ctx) error {
	provider_user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		return utils.FailResponse(c, "could not complete oauth transaction")
	}

	check_user := types.CreateUser{
		Email: provider_user.Email,
		Name: provider_user.Name,
		ImageUrl: provider_user.AvatarURL,
	}
	user, created, err := ctx.Service.UserService.FindOrCreate(c.Context(), check_user)
	if err != nil {
		return utils.FailResponse(c, "authentication could not succeed")
	}
	token, err := GenerateJWT(ctx.Tokens.JwtKey, &user)
	if err != nil {
		return utils.FailResponse(c, "could not generate token")
	}
	auth := types.Auth{
		Token: token,
		User: mappers.DbUserToUser(user),
	}

	if created {
		return utils.DataResponseCreated(c, auth)
	}
	return utils.DataResponse(c, auth)
}