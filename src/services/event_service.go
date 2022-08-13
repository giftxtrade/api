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

func (service *EventService) Create(create_event *types.CreateEvent, user *types.User, event *types.Event) error {
	validate := validator.New()
	if err := validate.Struct(create_event); err != nil {
		return err
	}

	event.Name = create_event.Name
	event.Description = create_event.Description
	event.Budget = create_event.Budget
	event.InviteMessage = create_event.InviteMessage
	event.DrawAt = create_event.DrawAt
	event.CloseAt = create_event.CloseAt
	event.Slug = slug.Make(create_event.Name)
	event.CreatedById = user.ID
	event.CreatedBy = *user
	event.ModifiedById = user.ID
	event.ModifiedBy = *user
	
	return service.DB.
		Table(service.TABLE).
		Create(event).
		Error
}
	return service.DB.
		Table(service.TABLE).
		Create(&new_event).
		Error
}