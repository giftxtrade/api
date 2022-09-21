package controllers

import (
	"strings"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctrl Controller) CreateEvent(ctx *fiber.Ctx) error {
	event_service := ctrl.Service.EventService
	cur_auth := utils.ParseAuthContext(ctx.UserContext())

	var input types.CreateEvent
	if err := ctx.BodyParser(&input); err != nil {
		return utils.FailResponse(ctx, "could not parse body")
	}
	validate_input := ctrl.Validator.Struct(&input)
	if validate_input != nil {
		errors := strings.Split(validate_input.Error(), "\n")
		return utils.FailResponse(ctx, errors...)
	}

	var new_event types.Event
	create_err := event_service.Create(&input, &cur_auth.User, &new_event)
	if create_err != nil {
		return utils.FailResponse(ctx, "could not create event")
	}

	return utils.DataResponseCreated(ctx, &new_event)
}