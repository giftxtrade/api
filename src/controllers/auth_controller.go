package controllers

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
)

func (ctx Controller) GetProfile(c *fiber.Ctx) error {
	auth := utils.ParseAuthContext(c.UserContext())
	return c.JSON(types.Result{
		Data: auth,
	})
}

// [GET] /auth/:provider
func (ctx Controller) SignIn(c *fiber.Ctx) error {
	return goth_fiber.BeginAuthHandler(c)
}

// [GET] /auth/{provider}/Callback
func (ctx Controller) Callback(c *fiber.Ctx) error {
	provider_user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		return c.JSON(types.Errors{
			Errors: []string{"could not complete oauth transaction"},
		})
	}

	check_user := types.CreateUser{
		Email: provider_user.Email,
		Name: provider_user.Name,
		ImageUrl: provider_user.AvatarURL,
	}
	var user types.User
	_, err = ctx.Service.UserService.FindOrCreate(&check_user, &user)
	if err != nil {
		return c.JSON(types.Errors{
			Errors: []string{"authentication could not succeed"},
		})
	}
	token, err := utils.GenerateJWT(ctx.Tokens.JwtKey, &user)
	if err != nil {
		return c.JSON(types.Errors{
			Errors: []string{"could not generate token"},
		})
	}
	auth := types.Auth{
		Token: token,
		User: user,
	}
	return c.JSON(types.Result{
		Data: auth,
	})
}