package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
	"google.golang.org/api/oauth2/v2"
)

// [GET] /auth/profile (authentication required)
func (ctr Controller) GetProfile(c *fiber.Ctx) error {
	auth := GetAuthContext(c.UserContext())
	return utils.DataResponse(c, auth)
}

// [GET] /auth/:provider
func (ctr Controller) SignIn(c *fiber.Ctx) error {
	return goth_fiber.BeginAuthHandler(c)
}

// [GET] /auth/:provider/callback
func (ctr Controller) Callback(c *fiber.Ctx) error {
	provider_user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		return utils.FailResponse(c, "could not complete oauth transaction")
	}

	check_user := types.CreateUser{
		Email: provider_user.Email,
		Name: provider_user.Name,
		ImageUrl: provider_user.AvatarURL,
	}
	auth, created, err := ctr.Service.UserService.GenerateAuthUser(c.Context(), check_user)
	if err != nil {
		return utils.FailResponse(c, err.Error())
	}
	if created {
		return utils.DataResponseCreated(c, auth)
	}
	return utils.DataResponse(c, auth)
}

func (ctr *Controller) GoogleVerify(c *fiber.Ctx) error {
	access_token := c.Query("access_token", "")
	if access_token == "" {
		return utils.FailResponse(c, "no access token provided")
	}

	res, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", access_token))
	if err != nil {
		return utils.FailResponse(c, "invalid access token")
	}
	userDataRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return utils.FailResponse(c, "could not read body")
	}
	var userData oauth2.Userinfo
	if err := json.Unmarshal(userDataRaw, &userData); err != nil {
		return utils.FailResponse(c, "could not parse response")
	}

	check_user := types.CreateUser{
		Email: userData.Email,
		Name: userData.Name,
		ImageUrl: userData.Picture,
	}
	auth, created, err := ctr.Service.UserService.GenerateAuthUser(c.Context(), check_user)
	if err != nil {
		return utils.FailResponse(c, err.Error())
	}
	if created {
		return utils.DataResponseCreated(c, auth)
	}
	return utils.DataResponse(c, auth)
}
