package mappers

import (
	"database/sql"
	"fmt"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
	"github.com/gosimple/slug"
	"golang.org/x/exp/maps"
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

func DbEventToEvent(event database.Event, participants []types.Participant, links []types.Link) types.Event {
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
		Links: links,
	}
}

func DbEventsToEventsSimple(event []database.Event) []types.Event {
	events := make([]types.Event, len(event))
	for i, row := range event {
		events[i] = DbEventToEvent(row, nil, nil)
	}
	return events
}

func DbEventLinkToEvent(event_link database.EventLink) types.Event {
	db_event := database.Event{
		ID: event_link.ID,
		Name: event_link.Name,
		Description: event_link.Description,
		Budget: event_link.Budget,
		InvitationMessage: event_link.InvitationMessage,
		DrawAt: event_link.DrawAt,
		CloseAt: event_link.CloseAt,
		CreatedAt: event_link.CreatedAt,
		UpdatedAt: event_link.UpdatedAt,
	}
	return DbEventToEvent(db_event, nil, nil)
}

func DbFindAllEventsWithUserRowToEvent(rows []database.FindAllEventsWithUserRow) []types.Event {
	events := []types.Event{}
	var prev_event_id int64 = 0
	for _, row := range rows {
		if row.Event.ID != prev_event_id {
			participant := DbParticipantUserToParticipant(row.ParticipantUser, nil)
			mapped_event := DbEventToEvent(row.Event, []types.Participant{participant}, nil) 
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
	event := DbEventLinkToEvent(rows[0].EventLink)
	link_map := map[int64]types.Link{}
	participant_map := map[int64]types.Participant{}
	participants := []types.Participant{} // TODO: optimize by preallocating memory
	for _, row := range rows {
		el := row.EventLink
		if el.LinkID.Valid && link_map[el.LinkID.Int64] == (types.Link{}) {
			link_map[el.LinkID.Int64] = types.Link{
				ID: el.LinkID.Int64,
				Code: el.LinkCode.String,
				ExpirationDate: el.LinkExpirationDate.Time,
				EventID: el.ID,
			}
		}

		pu := row.ParticipantUser
		if pu != (database.ParticipantUser{}) && participant_map[pu.ID] == (types.Participant{}) {
			mapped_participant := DbParticipantUserToParticipant(pu, nil)
			participant_map[pu.ID] = mapped_participant
			participants = append(participants, mapped_participant)
		}
	}
	event.Links = maps.Values(link_map)
	event.Participants = participants
	return event
}
