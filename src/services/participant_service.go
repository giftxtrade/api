package services

import (
	"fmt"

	"github.com/giftxtrade/api/src/types"
	"github.com/google/uuid"
)

type ParticipantService struct {
	ServiceBase
	UserService UserService
}

// Creates a new participant for a given event.
// Note that participant_user is optional
func (service ParticipantService) Create(
	user *types.User, 
	participant_user *types.User, 
	event *types.Event, 
	input *types.CreateParticipant, 
	output *types.Participant,
) error {
	err := service.input_to_participant(user, participant_user, event, input, output)
	if err != nil {
		return err
	}
	return service.DB.
		Table(service.TABLE).
		Create(output).
		Error
}

func (service ParticipantService) BulkCreate(
	user *types.User, 
	event *types.Event, 
	input []types.CreateParticipant,
) ([]types.Participant, error) {
	size := len(input)
	participants := make([]types.Participant, size)
	for i, participant_input := range input {
		participant := types.Participant{}
		var participant_user *types.User = nil
		if participant_input.Email == user.Email {
			participant_user = user
		}
		err := service.input_to_participant(user, participant_user, event, &participant_input, &participant)
		if err != nil {
			return nil, err
		}
		participants[i] = participant
	}

	create_err := service.DB.
		Table(service.TABLE).
		CreateInBatches(participants, size).
		Error

	if create_err != nil {
		return nil, create_err
	}
	return participants, nil
}

func (service ParticipantService) FindById(id string, output *types.Participant) error {
	return service.DB.
		Table(service.TABLE).
		Joins("CreatedBy").
		Joins("ModifiedBy").
		Joins("Event").
		Joins("User").
		Where("participants.id = ?", id).
		First(output).
		Error
}

func (service ParticipantService) Find(
	email string, 
	event_id string, 
	output *types.Participant,
) error {
	return service.DB.
		Table(service.TABLE).
		Joins("CreatedBy").
		Joins("ModifiedBy").
		Joins("Event").
		Joins("User").
		Where(
			"participants.event_id = ? AND participants.email = ?", 
			event_id, 
			email,
		).
		First(output).
		Error
}

func (service ParticipantService) Update(
	id string, 
	user *types.User, 
	participant_user *types.User,
	input *types.CreateParticipant, 
	output *types.Participant,
) (bool, error) {
	find_err := service.FindById(id, output)
	if find_err != nil {
		return false, find_err
	}

	updated := false
	if input.Address != "" && input.Address != output.Address {
		output.Address = input.Address
		updated = true
	}
	if input.Nickname != "" && input.Nickname != output.Nickname {
		output.Nickname = input.Nickname
		updated = true
	}
	if input.Participates != output.Participates {
		output.Participates = input.Participates
		updated = true
	}
	if !output.UserId.Valid && participant_user != nil {
		if participant_user.Email != output.Email {
			return false, fmt.Errorf("emails don't match")
		}

		output.UserId = uuid.NullUUID{
			Valid: true,
			UUID: participant_user.ID,
		}
		output.User = *participant_user
		output.Accepted = true
	}
	
	if updated {
		output.ModifiedBy = *user
		output.ModifiedById = user.ID
		err := service.DB.
			Table(service.TABLE).
			Save(output).
			Error
		return true, err
	}
	return false, nil	
}

func (service ParticipantService) Delete(id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return service.DB.
		Table(service.TABLE).
		Delete(&types.Participant{
			Base: types.Base{
				ID: uuid,
			},
		}).
		Error
}

// Identical to Find but with no joins
func (service ParticipantService) find_no_joins(
	email string, 
	event_id string, 
	output *types.Participant,
) error {
	return service.DB.
		Table(service.TABLE).
		Where(
			"participants.event_id = ? AND participants.email = ?", 
			event_id, 
			email,
		).
		First(output).
		Error
}

func (service ParticipantService) input_to_participant(
	user *types.User, 
	participant_user *types.User, 
	event *types.Event, 
	input *types.CreateParticipant, 
	output *types.Participant,
) error {
	if err := service.Validator.Struct(input); err != nil {
		return err
	}

	// check if participant with the email already exists for the event
	found_participant := types.Participant{}
	found_err := service.find_no_joins(input.Email, event.ID.String(), &found_participant)
	if found_err == nil {
		return fmt.Errorf("participant already exists")
	}

	output.CreatedBy = *user
	output.CreatedById = user.ID
	output.ModifiedBy = *user
	output.ModifiedById = user.ID

	output.Email = input.Email
	output.Nickname = input.Nickname
	output.Address = input.Address
	output.Organizer = input.Organizer
	output.Participates = input.Participates
	
	output.EventId = event.ID
	output.Event = *event

	if participant_user != nil {
		// check if participant_user.Email matches Email
		if participant_user.Email != input.Email {
			return fmt.Errorf("emails don't match")
		}

		output.Accepted = true
		output.User = *participant_user
		output.UserId = uuid.NullUUID{
			Valid: true,
			UUID: participant_user.ID,
		}
	} else {
		output.Accepted = false
	}
	return nil
}