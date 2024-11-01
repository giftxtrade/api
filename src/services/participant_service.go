package services

import (
	"context"
	"fmt"

	"github.com/giftxtrade/api/src/database/jet/postgres/public/table"
	"github.com/giftxtrade/api/src/types"
	"github.com/go-jet/jet/v2/qrm"
)

type ParticipantService struct {
	ServiceBase
}

func (s *ParticipantService) BulkCreateParticipant(
	ctx context.Context,
	user *types.User,
	event *types.Event,
	input []types.CreateParticipant,
) (participants []types.Participant, err error) {
	found_owner_participant := false
	for i, p := range input {
		if p.Organizer && p.Email == user.Email {
			found_owner_participant = true
			input[i].UserID = &user.ID
		}
		input[i].EventID = event.ID
	}
	if !found_owner_participant {
		return nil, fmt.Errorf(
			"%s (%s) must be in the participant list and have the organizer field set to 'true'",
			user.Name, 
			user.Email,
		)
	}

	qb := table.Participant.
		INSERT(
			table.Participant.Email,
			table.Participant.Name,
			table.Participant.Address,
			table.Participant.Organizer,
			table.Participant.Participates,
			table.Participant.UserID,
			table.Participant.EventID,
		).
		MODELS(input).
		RETURNING(table.Participant.AllColumns)
	var db qrm.Queryable = s.DB
	if s.TX != nil {
		db = s.TX
	}
	err = qb.QueryContext(ctx, db, &participants)
	return participants, err
}
