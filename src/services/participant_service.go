package services

import "github.com/giftxtrade/api/src/types"

type ParticipantService struct {
	ServiceBase
	UserService UserService
}

// Creates a new participant for a given event.
// Note that participant_user is optional
func (service ParticipantService) Create(user *types.User, participant_user *types.User, event *types.Event, input *types.CreateParticipant, output *types.Participant) error {
	if err := service.Validator.Struct(input); err != nil {
		return err
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
		output.Accepted = true
		output.User = *participant_user
		output.UserId = participant_user.ID
	}

	return service.DB.
		Table(service.TABLE).
		Create(output).
		Error
}