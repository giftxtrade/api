package services

import (
	"context"
	"fmt"

	"github.com/giftxtrade/api/src/database/jet/postgres/public/table"
	"github.com/giftxtrade/api/src/types"
	"github.com/go-jet/jet/v2/postgres"
)

type EventService struct {
	ServiceBase
	ParticipantService ParticipantService
}

func (s *EventService) FindEventsForUser(ctx context.Context, user types.User) (events []types.Event, err error) {
	events = []types.Event{}
	participant_user_sub_query := table.Participant.SELECT(
		table.Participant.AllColumns,
		table.User.AllColumns,
	).
	FROM(
		table.Participant.
			LEFT_JOIN(table.User, table.Participant.UserID.EQ(table.User.ID)),
	).
	ORDER_BY(table.Participant.ID.ASC()).
	AsTable(table.Participant.TableName())

	query := table.Event.SELECT(
		table.Event.AllColumns,
		table.Link.AllColumns,
		participant_user_sub_query.AllColumns(),
	).FROM(
		table.Event.
			LEFT_JOIN(table.Link, table.Event.ID.EQ(table.Link.EventID)).
			LEFT_JOIN(table.Participant.AS("p1"), table.Event.ID.EQ(table.Participant.AS("p1").EventID)).
			LEFT_JOIN(
				participant_user_sub_query,
				table.Event.ID.EQ(table.Participant.EventID.From(participant_user_sub_query)),
			),
	).
	WHERE(
		table.Participant.AS("p1").UserID.EQ(postgres.Int(user.ID)),
	).
	ORDER_BY(
		table.Event.DrawAt.ASC(),
		table.Event.CloseAt.ASC(),
		table.Participant.ID.From(participant_user_sub_query).ASC(),
	)
	err = query.QueryContext(ctx, s.DB, &events)
	return events, err
}

func (s *EventService) FindEventById(ctx context.Context, user types.User, event_id int64) (event types.Event, err error) {
	query := table.Event.SELECT(
		table.Event.AllColumns,
		table.Participant.AllColumns,
		table.User.AllColumns,
		table.Link.AllColumns,
		table.Wish.AllColumns,
		table.Product.AllColumns,
	).FROM(
		table.Event.
			LEFT_JOIN(table.Link, table.Event.ID.EQ(table.Link.EventID)).
			INNER_JOIN(table.Participant, table.Event.ID.EQ(table.Participant.EventID)).
			LEFT_JOIN(table.User, table.Participant.UserID.EQ(table.User.ID)).
			LEFT_JOIN(
				table.Wish,
				table.Event.ID.EQ(table.Wish.EventID).
				AND(
					table.Wish.UserID.EQ(postgres.Int(user.ID)),
				),
			).
			LEFT_JOIN(table.Product, table.Wish.ProductID.EQ(table.Product.ID)),
	).WHERE(table.Event.ID.EQ(postgres.Int64(event_id))).ORDER_BY(
		table.Participant.Organizer.DESC(),
		table.Participant.Accepted.DESC(),
		table.Participant.CreatedAt.DESC(),
	)
	err = query.QueryContext(ctx, s.DB, &event)
	return event, err
}

func (s *EventService) CreateEvent(ctx context.Context, user *types.User, input types.CreateEvent) (event types.Event, err error) {
	s.TX, err = s.DB.BeginTx(ctx, nil)
	defer s.TX.Rollback()

	qb := table.Event.INSERT(
		table.Event.Name,
		table.Event.Description,
		table.Event.Budget,
		table.Event.InvitationMessage,
		table.Event.DrawAt,
		table.Event.CloseAt,
	).MODEL(input).RETURNING(table.Event.AllColumns)
	if err = qb.QueryContext(ctx, s.TX, &event); err != nil {
		return types.Event{}, fmt.Errorf("could not create event")
	}

	// create participants for event in transaction scope
	s.ParticipantService.TX = s.TX
	participants, err := s.ParticipantService.BulkCreateParticipant(ctx, user, &event, input.Participants)
	if err != nil {
		return types.Event{}, err
	}

	if s.TX.Commit() != nil {
		return types.Event{}, fmt.Errorf("could not commit transaction")
	}
	// TODO: email participants?
	event.Participants = participants
	return event, nil
}
