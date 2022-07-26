package services

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
)

type EventService struct {
	ServiceBase
	UserService UserService
}

func (service *EventService) Create(create_event *types.CreateEvent, event *types.Event) error {
	validate := validator.New()
	if err := validate.Struct(create_event); err != nil {
		return err
	}

	new_event := types.Event{
		Name: create_event.Name,
		Description: event.Description,
		Budget: create_event.Budget,
		InviteMessage: create_event.InviteMessage,
		DrawAt: create_event.DrawAt,
		CloseAt: create_event.CloseAt,
		Slug: slug.Make(create_event.Name),
		UserActionBase: types.UserActionBase{
			CreatedById: create_event.CreatedBy.ID,
			CreatedBy: create_event.CreatedBy,
			ModifiedById: create_event.CreatedBy.ID,
			ModifiedBy: create_event.CreatedBy,
		},
	}
	return service.DB.
		Table(service.TABLE).
		Create(new_event).
		Error
}