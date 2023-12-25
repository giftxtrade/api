package controllers

import (
	"database/sql"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctr *Controller) CreateEvent(c *fiber.Ctx) error {
	auth_user := ParseAuthContext(c.UserContext())
	var input types.CreateEvent
	if c.BodyParser(&input) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}
	if err := ctr.Validator.Struct(input); err != nil {
		return utils.FailResponse(c, err.Error())
	}

	event, err := ctr.Service.EventService.CreateEvent(c.Context(), &auth_user.User, input)
	if err != nil {
		return utils.FailResponse(c, "could not create event", err.Error())
	}
	return utils.DataResponseCreated(c, event)
}

func (ctr *Controller) GetEvents(c *fiber.Ctx) error {
	auth_user := ParseAuthContext(c.UserContext())
	events, err := ctr.Querier.FindAllEventsWithUser(c.Context(), sql.NullInt64{
		Valid: true,
		Int64: auth_user.User.ID,
	})
	if err != nil {
		return utils.FailResponse(c, "could not return events", err.Error())
	}

	mapped_events := services.DbFindAllEventsWithUserRowToEvent(events)
	return utils.DataResponse(c, mapped_events)
}
