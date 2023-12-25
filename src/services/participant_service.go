package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
)

type ParticipantService struct {
	ServiceBase
}

func (s *ParticipantService) BulkCreateParticipant(
	tx *sql.Tx, ctx context.Context,
	user *types.User,
	event *database.Event,
	input []types.CreateParticipant,
) ([]types.Participant, error) {
	q := s.Querier.WithTx(tx)

	found_creator_participant := false
	var creator_participant types.CreateParticipant
	participants := make([]types.Participant, len(input))
	for i, p := range input {
		data := CreateParticipantToDbCreateParticipantParams(p, event)
		if p.Organizer && p.Email == user.Email {
			found_creator_participant = true
			creator_participant = p
			data.UserID = sql.NullInt64{
				Valid: true,
				Int64: user.ID,
			}
			data.Accepted = true
		}

		new_participant, err := q.CreateParticipant(ctx, data)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("could not create participant %s (%s)", p.Name, p.Email)
		}
		participants[i] = DbParticipantToParticipant(new_participant, event, nil)
	}

	if !found_creator_participant {
		tx.Rollback()
		return nil, fmt.Errorf(
			"%s (%s) must have the organizer field set to 'true'",
			creator_participant.Name, 
			creator_participant.Email,
		)
	}
	return participants, nil
}

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
	}
	if event != nil {
		result.Event = DbEventToEvent(*event, nil)
		result.EventID = event.ID
	}
	if user != nil {
		result.User = DbUserToUser(*user)
		result.UserID = user.ID
	}
	return result
}
