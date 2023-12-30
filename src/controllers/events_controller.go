package controllers

import (
	"database/sql"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/mappers"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctr *Controller) CreateEvent(c *fiber.Ctx) error {
	auth_user := ParseAuthContext(c.UserContext())
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
	auth_user := ParseAuthContext(c.UserContext())
	events, err := ctr.Querier.FindAllEventsWithUser(c.Context(), sql.NullInt64{
		Valid: true,
		Int64: auth_user.User.ID,
	})
	if err != nil {
		return utils.FailResponse(c, "could not return events", err.Error())
	}

	mapped_events := mappers.DbFindAllEventsWithUserRowToEvent(events)
	return utils.DataResponse(c, mapped_events)
}

func (ctr *Controller) GetEventById(c *fiber.Ctx) error {
	event_id := c.UserContext().Value(EVENT_ID_PARAM_KEY).(int64)
	event_rows, err := ctr.Querier.FindEventById(c.Context(), event_id)
	if err != nil {
		return utils.FailResponse(c, "could not load event")
	}

	event := mappers.DbFindEventByIdToEvent(event_rows)
	return utils.DataResponse(c, event)
}

func (ctr *Controller) GetInvites(c *fiber.Ctx) error {
	auth := ParseAuthContext(c.UserContext())
	rows, err := ctr.Querier.FindEventInvites(c.Context(), auth.User.Email)
	if err != nil {
		return utils.FailResponse(c, "could not fetch invites")
	}
	return utils.DataResponse(c, mappers.DbEventsToEventsSimple(rows))
}

func (ctr *Controller) AcceptEventInvite(c *fiber.Ctx) error {
	auth := ParseAuthContext(c.UserContext())
	event_id := c.UserContext().Value(EVENT_ID_PARAM_KEY).(int64)

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
	auth := ParseAuthContext(c.UserContext())
	event_id := c.UserContext().Value(EVENT_ID_PARAM_KEY).(int64)
	_, err := ctr.Querier.DeclineEventInvite(c.Context(), database.DeclineEventInviteParams{
		EventID: event_id,
		UserID: sql.NullInt64{
			Valid: true,
			Int64: auth.User.ID,
		},
	})
	if err != nil {
		return utils.FailResponse(c, "could not decline event invitation. please try again.")
	}
	return utils.DataResponse(c, struct{Success bool}{
		Success: true,
	})
}
