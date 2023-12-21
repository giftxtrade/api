package controllers

import (
	"database/sql"
	"fmt"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctr *Controller) CreateEvent(c *fiber.Ctx) error {
	var create_product types.CreateEvent
	if c.BodyParser(&create_product) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}
	if err := ctr.Validator.Struct(create_product); err != nil {
		return utils.FailResponse(c, err.Error())
	}

	new_event, err := ctr.Querier.CreateEvent(c.Context(), database.CreateEventParams{
		Name: create_product.Name,
		Description: sql.NullString{
			String: create_product.Description,
			Valid: create_product.Description != "",
		},
		Budget: fmt.Sprintf("%f", create_product.Budget),
		InvitationMessage: create_product.InviteMessage,
		DrawAt: create_product.DrawAt,
		CloseAt: create_product.CloseAt,
	})
	if err != nil {
		return utils.FailResponse(c, "could not create event")
	}
	return utils.DataResponseCreated(c, new_event)
}
