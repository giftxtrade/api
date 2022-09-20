package services

import (
	"time"

	"github.com/giftxtrade/api/src/types"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type EventService struct {
	ServiceBase
	UserService UserService
}

func (service *EventService) Create(input *types.CreateEvent, user *types.User, output *types.Event) error {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return err
	}

	output.Name = input.Name
	output.Description = input.Description
	output.Budget = input.Budget
	output.InviteMessage = input.InviteMessage
	output.DrawAt = input.DrawAt
	output.CloseAt = input.CloseAt
	output.Slug = slug.Make(input.Name)
	output.CreatedById = user.ID
	output.CreatedBy = *user
	output.ModifiedById = user.ID
	output.ModifiedBy = *user
	
	return service.DB.
		Table(service.TABLE).
		Create(output).
		Error
}

func (service *EventService) FindById(id string, output *types.Event) error {
	return service.DB.
		Table(service.TABLE).
		Preload("CreatedBy").
		Preload("ModifiedBy").
		Where("id = ?", id).
		First(output).
		Error
}

// update event given a user that modified it.
// event must be an already existing row.
// Boolean value is true if event was updated, otherwise false (with error).
func (service *EventService) Patch(user *types.User, input *types.CreateEvent, event *types.Event) (bool, error) {
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
		Save(event).
		Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (service *EventService) Delete(id string) error {
	parsed_uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return service.DB.
		Table(service.TABLE).
		Delete(types.Event{
			Base: types.Base{
				ID: parsed_uuid,
			},
		}).
		Error
}