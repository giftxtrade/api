package controllers

import (
	"net/http"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (controller *Controller) Home(c *fiber.Ctx) error {
	return c.JSON(types.Response{
		Message: "GiftTrade REST API ⚡",
	})
}

func (controller *Controller) NotFound(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithStatusCode(w, 404, types.Errors{
		Errors: []string{"resource not found"},
	})
}