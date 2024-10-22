package controllers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/database/jet/postgres/public/table"
	"github.com/giftxtrade/api/src/mappers"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/gofiber/fiber/v2"
)

const EVENT_LINK_CODE_LEN = 15

func (ctr *Controller) CreateEvent(c *fiber.Ctx) error {
	auth_user := GetAuthContext(c.UserContext())
	var input types.CreateEvent
	if c.BodyParser(&input) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}
	if err := ctr.Validator.Struct(input); err != nil {
		return utils.FailResponse(c, err.Error())
	}

	event, err := ctr.Service.EventService.CreateEvent(c.Context(), &auth_user.User, input)
	if err != nil {
		return utils.FailResponse(c, "could not create event", err.Error())
	}
	return utils.DataResponseCreated(c, event)
}

func (ctr *Controller) GetEvents(c *fiber.Ctx) error {
	auth_user := GetAuthContext(c.UserContext())

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
		table.Participant.AS("p1").UserID.EQ(postgres.Int(auth_user.User.ID)),
	).
	ORDER_BY(
		table.Event.DrawAt.ASC(),
		table.Event.CloseAt.ASC(),
		table.Participant.ID.From(participant_user_sub_query).ASC(),
	)

	var dest []types.Event
	err := query.QueryContext(c.Context(), ctr.DB, &dest)
	if err != nil {
		fmt.Println(query.DebugSql(), err)
		return utils.FailResponse(c, "could not return events")
	}
	return utils.DataResponse(c, dest)
}

func (ctr *Controller) GetEventById(c *fiber.Ctx) error {
	auth := GetAuthContext(c.UserContext())
	event_id := GetEventIdFromContext(c.UserContext())

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
					table.Wish.UserID.EQ(postgres.Int(auth.User.ID)),
				),
			).
			LEFT_JOIN(table.Product, table.Wish.ProductID.EQ(table.Product.ID)),
	).WHERE(table.Event.ID.EQ(postgres.Int64(event_id))).ORDER_BY(
		table.Participant.Organizer.DESC(),
		table.Participant.Accepted.DESC(),
		table.Participant.CreatedAt.DESC(),
	)

	var event types.Event
	err := query.QueryContext(c.Context(), ctr.DB, &event)
	if err != nil {
		return utils.FailResponse(c, "could not load event")
	}
	return utils.DataResponse(c, event)
}

func (ctr *Controller) GetInvites(c *fiber.Ctx) error {
	auth := GetAuthContext(c.UserContext())
	rows, err := ctr.Querier.FindEventInvites(c.Context(), auth.User.Email)
	if err != nil {
		return utils.FailResponse(c, "could not fetch invites")
	}
	return utils.DataResponse(c, mappers.DbEventsToEventsSimple(rows))
}

func (ctr *Controller) AcceptEventInvite(c *fiber.Ctx) error {
	auth := GetAuthContext(c.UserContext())
	event_id := GetEventIdFromContext(c.UserContext())

	tx, err := ctr.DB.BeginTx(c.Context(), nil)
	if err != nil {
		tx.Rollback()
		return utils.FailResponse(c, "transaction error. please try again")
	}
	q := ctr.Querier.WithTx(tx)
	defer q.Close()

	participant, err := q.AcceptEventInvite(c.Context(), database.AcceptEventInviteParams{
		EventID: event_id,
		UserID: sql.NullInt64{
			Valid: true,
			Int64: auth.User.ID,
		},
		Email: auth.User.Email,
	})
	if err != nil {
		tx.Rollback()
		return utils.FailResponse(c, "could not accept invite for event")
	}

	event_rows, err := q.FindEventById(c.Context(), participant.EventID)
	if err != nil {
		tx.Rollback()
		return utils.FailResponse(c, "could not fetch event")
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return utils.FailResponse(c, "could not save changes")
	}
	event := mappers.DbFindEventByIdToEvent(event_rows)
	return utils.DataResponse(c, event)
}

func (ctr *Controller) DeclineEventInvite(c *fiber.Ctx) error {
	auth := GetAuthContext(c.UserContext())
	event_id := GetEventIdFromContext(c.UserContext())
	_, err := ctr.Querier.DeclineEventInvite(c.Context(), database.DeclineEventInviteParams{
		EventID: event_id,
		Email: auth.User.Email,
	})
	if err != nil {
		return utils.FailResponse(c, "could not decline event invitation. please try again.")
	}
	return utils.DataResponse(c, types.DeleteStatus{
		Deleted: true,
	})
}

// [PATCH] events/:event_id - Organizer Auth
func (ctr *Controller) UpdateProduct(c *fiber.Ctx) error {
	event_id := GetEventIdFromContext(c.UserContext())
	var input types.UpdateEvent
	if c.BodyParser(&input) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}
	if err := ctr.Validator.Struct(input); err != nil {
		return utils.FailResponse(c, "validation errors with input", err.Error())
	}

	_, err := ctr.Querier.UpdateEvent(c.Context(), database.UpdateEventParams{
		ID: event_id,
		Name: sql.NullString{
			Valid: input.Name != "",
			String: input.Name,
		},
		Description: sql.NullString{
			Valid: input.Description != "",
			String: input.Description,
		},
		Budget: sql.NullString{
			Valid: input.Budget != 0,
			String: fmt.Sprintf("%f", input.Budget),
		},
		DrawAt: sql.NullTime{
			Valid: !input.DrawAt.IsZero(),
			Time: input.DrawAt,
		},
		CloseAt: sql.NullTime{
			Valid: !input.CloseAt.IsZero(),
			Time: input.CloseAt,
		},
	})
	if err != nil {
		return utils.FailResponse(c, "could not update event")
	}

	event_row, err := ctr.Querier.FindEventById(c.Context(), event_id)
	if err != nil {
		return utils.FailResponse(c, "could not return event")
	}
	event := mappers.DbFindEventByIdToEvent(event_row)
	return utils.DataResponse(c, event)
}

// [DELETE] /events/:event_id - Uses organizer auth
func (ctr *Controller) DeleteEvent(c *fiber.Ctx) error {
	event_id := GetEventIdFromContext(c.UserContext())
	_, err := ctr.Querier.DeleteEvent(c.Context(), event_id)
	if err != nil {
		return utils.FailResponse(c, "event could not be deleted. please try again.")
	}
	return utils.DataResponse(c, types.DeleteStatus{
		Deleted: true,
	})
}

// [GET] /events/:event_id/get-link - Uses event participant auth
func (ctr *Controller) GetEventLink(c *fiber.Ctx) error {
	event_id := GetEventIdFromContext(c.UserContext())
	event, _ := ctr.Querier.FindEventSimple(c.Context(), event_id)
	code, _ := utils.GenerateRandomUrlEncodedString(EVENT_LINK_CODE_LEN)
	link, err := ctr.Querier.CreateLink(c.Context(), database.CreateLinkParams{
		EventID: event_id,
		Code: code,
		ExpirationDate: event.DrawAt,
	})
	if err != nil {
		return utils.FailResponse(c, "could not create link for event")
	}
	return utils.DataResponseCreated(c, mappers.DbLinkToLink(link, nil))
}

func (Ctr *Controller) get_invite_code(c *fiber.Ctx) (code string, error error) {
	invite_code := c.Params("invite_code")
	if len(invite_code) != EVENT_LINK_CODE_LEN {
		return "", fmt.Errorf("invalid invite code")
	}
	return invite_code, nil
}

func (ctr *Controller) VerifyEventLinkCode(c *fiber.Ctx) error {
	invite_code, err := ctr.get_invite_code(c)
	if err != nil {
		return utils.FailResponse(c, err.Error())
	}
	res, err := ctr.Querier.FindLinkWithEventByCode(c.Context(), invite_code)
	if err != nil {
		return utils.FailResponse(c, "invite code expired or invalid")
	}
	return utils.DataResponse(c, mappers.DbLinkToLink(res.Link, &res.Event))
}

func (ctr *Controller) JoinEventViaInviteCode(c *fiber.Ctx) error {
	invite_code, err := ctr.get_invite_code(c)
	if err != nil {
		return utils.FailResponse(c, err.Error())
	}
	res, err := ctr.Querier.FindLinkByCode(c.Context(), invite_code)
	if err != nil || res.ExpirationDate.Before(time.Now()) {
		return utils.FailResponse(c, "invite code is expired or invalid")
	}

	auth := GetAuthContext(c.UserContext())
	_, err = ctr.Querier.VerifyEventWithEmailOrUser(c.Context(), database.VerifyEventWithEmailOrUserParams{
		EventID: res.EventID,
		UserID: sql.NullInt64{
			Valid: true,
			Int64: auth.User.ID,
		},
		Email: sql.NullString{
			Valid: true,
			String: auth.User.Email,
		},
	})
	// auth user is already a participant in the event
	if err == nil {
		p, _ := ctr.Querier.FindParticipantFromEventIdAndUser(c.Context(), database.FindParticipantFromEventIdAndUserParams{
			EventID: res.EventID,
			UserID: sql.NullInt64{
				Valid: true,
				Int64: auth.User.ID,
			},
		})
		return utils.DataResponse(c, mappers.DbParticipantToParticipant(p, nil, nil))
	}

	participant, err := ctr.Querier.CreateParticipant(c.Context(), database.CreateParticipantParams{
		Name: auth.User.Name,
		Email: auth.User.Email,
		Organizer: false,
		Participates: true,
		Accepted: true,
		EventID: res.EventID,
		UserID: sql.NullInt64{
			Valid: true,
			Int64: auth.User.ID,
		},
	})
	if err != nil {
		return utils.FailResponse(c, "could not join event")
	}
	return utils.DataResponseCreated(c, mappers.DbParticipantToParticipant(participant, nil, nil))
}
