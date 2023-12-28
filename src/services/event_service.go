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
	new_event, err := q.CreateEvent(ctx, CreateEventToDbCreateEventParams(input))
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
	mapped_event := DbEventToEvent(new_event, participants)
	return mapped_event, nil
}

func CreateEventToDbCreateEventParams(input types.CreateEvent) database.CreateEventParams {
	return database.CreateEventParams{
		Name: input.Name,
		Description: sql.NullString{
			String: input.Description,
			Valid: input.Description != "",
		},
		Budget: fmt.Sprintf("%f", input.Budget),
		InvitationMessage: input.InviteMessage,
		DrawAt: input.DrawAt,
		CloseAt: input.CloseAt,
	}
}

func DbEventToEvent(event database.Event, participants []types.Participant) types.Event {
	return types.Event{
		ID: event.ID,
		Name: event.Name,
		Description: event.Description.String,
		Budget: event.Budget,
		InvitationMessage: event.InvitationMessage,
		DrawAt: event.DrawAt,
		CloseAt: event.CloseAt,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
		Participants: participants,
	}
}

func DbFindAllEventsWithUserRowToEvent(rows []database.FindAllEventsWithUserRow) []types.Event {
	events := []types.Event{}
	var prev_event_id int64 = 0
	for _, row := range rows {
		if row.Event.ID != prev_event_id {
			participant := DbParticipantUserToParticipant(row.ParticipantUser, nil)
			mapped_event := DbEventToEvent(row.Event, append([]types.Participant{}, participant)) 
			events = append(events, mapped_event)
			
			prev_event_id = row.Event.ID
			continue
		}
		last_index := len(events) - 1
		events[last_index].Participants = append(
			events[last_index].Participants,
			DbParticipantUserToParticipant(row.ParticipantUser, nil),
		)
	}
	return events
}

func DbFindEventByIdToEvent(rows []database.FindEventByIdRow) types.Event {
	mapped_rows := make([]database.FindAllEventsWithUserRow, len(rows))
	for i, row := range rows {
		mapped_rows[i] = database.FindAllEventsWithUserRow(row)
	}
	return DbFindAllEventsWithUserRowToEvent(mapped_rows)[0]
}
