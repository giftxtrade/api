package mappers

import (
	"database/sql"
	"fmt"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
	"github.com/gosimple/slug"
)

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
		Slug: slug.Make(event.Name),
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

func DbEventsToEventsSimple(event []database.Event) []types.Event {
	events := make([]types.Event, len(event))
	for i, row := range event {
		events[i] = DbEventToEvent(row, nil)
	}
	return events
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
