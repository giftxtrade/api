package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/mappers"
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
	defer q.Close()

	found_creator_participant := false
	participants := make([]types.Participant, len(input))
	for i, p := range input {
		data := mappers.CreateParticipantToDbCreateParticipantParams(p, event)
		if p.Organizer && p.Email == user.Email {
			found_creator_participant = true
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
		participants[i] = mappers.DbParticipantToParticipant(new_participant, event, nil)
	}

	if !found_creator_participant {
		tx.Rollback()
		return nil, fmt.Errorf(
			"%s (%s) must be in the participant list and have the organizer field set to 'true'",
			user.Name, 
			user.Email,
		)
	}
	return participants, nil
}
