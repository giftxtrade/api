package controllers

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctx Controller) Home(c *fiber.Ctx) error {
	return utils.JsonResponse(c, types.Response{
		Message: "GiftTrade REST API ⚡",
	})
}

func (ctx Controller) NotFound(c *fiber.Ctx) error {
	return utils.ResponseWithStatusCode(c, 404, types.Errors{
		Errors: []string{"resource not found"},
	})
}