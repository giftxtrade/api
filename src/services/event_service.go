package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
)

type EventService struct {
	ServiceBase
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
	new_event, err := q.CreateEvent(ctx, database.CreateEventParams{
		Name: input.Name,
		Description: sql.NullString{
			String: input.Description,
			Valid: input.Description != "",
		},
		Budget: fmt.Sprintf("%f", input.Budget),
		InvitationMessage: input.InviteMessage,
		DrawAt: input.DrawAt,
		CloseAt: input.CloseAt,
	})
	if err != nil {
		tx.Rollback()
		return types.Event{}, fmt.Errorf("could not create event")
	}

	// create participants for event in transaction scope
	participants, err := s.BulkCreateParticipant(tx, ctx, user, &new_event, input.Participants)
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
	mapped_event := types.Event{
		ID: new_event.ID,
		Name: new_event.Name,
		Description: new_event.Description.String,
		Budget: new_event.Budget,
		InvitationMessage: new_event.Budget,
		DrawAt: new_event.DrawAt,
		CloseAt: new_event.CloseAt,
		CreatedAt: new_event.CreatedAt,
		UpdatedAt: new_event.UpdatedAt,
		Participants: participants,
	}
	return mapped_event, nil
}

func (s *EventService) BulkCreateParticipant(
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
		data := database.CreateParticipantParams{
			Name: p.Name,
			Email: p.Email,
			Organizer: p.Organizer,
			Participates: p.Participates,
			Accepted: false,
			EventID: event.ID,
			Address: sql.NullString{
				Valid: p.Address != "",
				String: p.Address,
			},
		}
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
		participants[i] = types.Participant{
			ID: new_participant.ID,
			Name: new_participant.Name,
			Email: new_participant.Email,
			Address: new_participant.Address,
			Organizer: new_participant.Organizer,
			Participates: new_participant.Participates,
			Accepted: new_participant.Accepted,
			EventID: event.ID,
			UserID: new_participant.UserID.Int64,
		}
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
