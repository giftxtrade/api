package utils

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func SetUserContext(c *fiber.Ctx, key interface{}, value interface{}) {
	c.SetUserContext(context.WithValue(c.UserContext(), key, value))
}
