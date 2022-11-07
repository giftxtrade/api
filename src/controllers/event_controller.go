package controllers

import (
	"strings"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctx Controller) CreateEvent(c *fiber.Ctx) error {
	event_service := ctx.Service.EventService
	cur_auth := utils.ParseAuthContext(c.UserContext())

	var input types.CreateEvent
	if err := c.BodyParser(&input); err != nil {
		return utils.FailResponse(c, "could not parse body")
	}
	var new_event types.Event
	create_err := event_service.Create(&input, &cur_auth.User, &new_event)
	if create_err != nil {
		errors := strings.Split(create_err.Error(), "\n")
		return utils.FailResponse(c, errors...)
	}

	return utils.DataResponseCreated(c, &new_event)
}

func (ctx Controller) GetAllEvents(c *fiber.Ctx) error {
	event_service := ctx.Service.EventService
	cur_auth := utils.ParseAuthContext(c.UserContext())

	events := new([]types.Event)
	if err := event_service.FindAllForUser(&cur_auth.User, events); err != nil {
		return utils.FailResponse(c, err.Error())
	}
	return utils.DataResponse(c, events)
}