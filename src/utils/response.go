package utils

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/gofiber/fiber/v2"
)

func ResponseWithStatusCode(c *fiber.Ctx, statusCode int, data interface{}) error {
	c.Response().SetStatusCode(statusCode)
	return c.JSON(data)
}

func JsonResponse(c *fiber.Ctx, data interface{}) error {
	return ResponseWithStatusCode(c, 200, data)
}

// Writes a types.Errors json response to the http.ResponseWriter,
// with a default Http 400 status
func FailResponse(c *fiber.Ctx, errors ...string) error {
	return ResponseWithStatusCode(c, 400, types.Errors{
		Errors: errors,
	})
}

func FailResponseUnauthorized(c *fiber.Ctx, errors interface{}) error {
	return ResponseWithStatusCode(c, 401, types.Errors{
		Errors: errors,
	})
}

// Writes a types.Data json response to the http.ResponseWriter,
// with a default Http 200 status
func DataResponse(c *fiber.Ctx, data interface{}) error {
	return ResponseWithStatusCode(c, 200, types.Result{
		Data: data,
	})
}