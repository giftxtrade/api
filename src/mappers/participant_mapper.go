package mappers

import (
	"database/sql"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
)

func CreateParticipantToDbCreateParticipantParams(input types.CreateParticipant, event *database.Event) database.CreateParticipantParams {
	return database.CreateParticipantParams{
		Name: input.Name,
		Email: input.Email,
		Organizer: input.Organizer,
		Participates: input.Participates,
		Accepted: false,
		EventID: event.ID,
		Address: sql.NullString{
			Valid: input.Address != "",
			String: input.Address,
		},
	}
}

func DbParticipantToParticipant(participant database.Participant, event *database.Event, user *database.User) types.Participant {
	result := types.Participant{
		ID: participant.ID,
		Name: participant.Name,
		Email: participant.Email,
		Address: participant.Address.String,
		Organizer: participant.Organizer,
		Participates: participant.Participates,
		Accepted: participant.Accepted,
		EventID: participant.EventID,
	}
	if participant.UserID.Valid {
		result.UserID = participant.UserID.Int64
	}
	if event != nil {
		event := DbEventToEvent(*event, nil, nil)
		result.Event = &event
		result.EventID = event.ID
	}
	if user != nil {
		user := DbUserToUser(*user)
		result.User = &user
		result.UserID = user.ID
	}
	return result
}

func DbParticipantUserToParticipant(participant_user database.ParticipantUser, event *database.Event) types.Participant {
	var user *database.User = nil
	if participant_user.UserID.Valid {
		user = &database.User{
			ID: participant_user.UserID.Int64,
			Name: participant_user.UserName.String,
			Email: participant_user.UserEmail.String,
			ImageUrl: participant_user.UserImageUrl.String,
		}
	}
	return DbParticipantToParticipant(
		database.Participant{
			ID: participant_user.ID,
			Name: participant_user.Name,
			Email: participant_user.Email,
			Address: participant_user.Address,
			Organizer: participant_user.Organizer,
			Participates: participant_user.Participates,
			Accepted: participant_user.Accepted,
			EventID: participant_user.EventID,
			UserID: participant_user.UserID,
		},
		event,
		user,
	)
}
