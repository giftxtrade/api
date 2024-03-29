package utils

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/gofiber/fiber/v2"
)

func ResponseWithStatusCode[T any](c *fiber.Ctx, statusCode int, data T) error {
	return c.Status(statusCode).JSON(data)
}

// Generic json response with status code 200
func JsonResponse[T any](c *fiber.Ctx, data T) error {
	return ResponseWithStatusCode(c, fiber.StatusOK, data)
}

// types.Error json response with status code 400
func FailResponse(c *fiber.Ctx, errors ...string) error {
	return ResponseWithStatusCode(c, fiber.StatusBadRequest, types.Errors{
		Errors: errors,
	})
}

// types.Error json response with status code 401
func FailResponseUnauthorized(c *fiber.Ctx, errors ...string) error {
	return ResponseWithStatusCode(c, fiber.StatusUnauthorized, types.Errors{
		Errors: errors,
	})
}

// types.Error json response with status code 404
func FailResponseNotFound(c *fiber.Ctx, errors ...string) error {
	return ResponseWithStatusCode(c, fiber.StatusNotFound, types.Errors{
		Errors: errors,
	})
}

// types.Data json response with status code 200
func DataResponse[T any](c *fiber.Ctx, data T) error {
	return ResponseWithStatusCode(c, fiber.StatusOK, data)
}

// types.Data json response with status code 201
func DataResponseCreated[T any](c *fiber.Ctx, data T) error {
	return ResponseWithStatusCode(c, fiber.StatusCreated, data)
}