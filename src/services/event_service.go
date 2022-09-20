package services

import (
	"time"

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

func (service *EventService) FindById(id string, event *types.Event) error {
	return service.DB.
		Table(service.TABLE).
		Preload("CreatedBy").
		Preload("ModifiedBy").
		Where("id = ?", id).
		First(event).
		Error
}

// update event given an event_id, user that modified it, and the destination.
// Boolean value is true if event was updated, otherwise false (with error).
func (service *EventService) Patch(id string, user *types.User, input *types.CreateEvent, event *types.Event) (bool, error) {
	updated := false

	if input.Name != "" && input.Name != event.Name {
		event.Name = input.Name
		updated = true
	}
	if input.Budget != 0 && input.Budget != event.Budget {
		event.Budget = input.Budget
		updated = true
	}
	if input.Description != "" && input.Description != event.Description {
		event.Description = input.Description
		updated = true
	}
	nil_time := time.Time{}
	if input.CloseAt != nil_time && input.CloseAt != event.CloseAt {
		event.CloseAt = input.CloseAt
		updated = true
	}
	if input.DrawAt != nil_time && input.DrawAt != event.DrawAt {
		event.DrawAt = input.DrawAt
		updated = true
	}
	if input.InviteMessage != "" && input.InviteMessage != event.InviteMessage {
		event.InviteMessage = input.InviteMessage
		updated = true
	}

	if !updated {
		return false, nil
	}
	event.ModifiedById = user.ID
	event.ModifiedBy = *user
	err := service.DB.
		Table(service.TABLE).
		Where("id = ?", id).
		Save(event).
		Error
	if err != nil {
		return false, err
	}
	return true, nil
}