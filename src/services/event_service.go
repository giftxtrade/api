package services

import (
	"context"
	"fmt"

	"github.com/giftxtrade/api/src/mappers"
	"github.com/giftxtrade/api/src/types"
)

type EventService struct {
	ServiceBase
	ParticipantService ParticipantService
}

func (s *EventService) CreateEvent(ctx context.Context, user *types.User, input types.CreateEvent) (types.Event, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return types.Event{}, fmt.Errorf("could not process. error with database transaction")
	}
	q := s.Querier.WithTx(tx)
	defer q.Close()

	// create new event in transaction scope
	new_event, err := q.CreateEvent(ctx, mappers.CreateEventToDbCreateEventParams(input))
	if err != nil {
		tx.Rollback()
		return types.Event{}, fmt.Errorf("could not create event")
	}

	// create participants for event in transaction scope
	participants, err := s.ParticipantService.BulkCreateParticipant(tx, ctx, user, &new_event, input.Participants)
	if err != nil {
		tx.Rollback()
		return types.Event{}, err
	}

	// commit all changes! create event, and all participants
	if tx.Commit() != nil {
		tx.Rollback()
		return types.Event{}, fmt.Errorf("could not commit transaction")
	}
	// build new event dto
	mapped_event := mappers.DbEventToEvent(new_event, participants, nil)
	return mapped_event, nil
}
