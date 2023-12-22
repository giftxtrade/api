package controllers

import (
	"database/sql"
	"fmt"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctr *Controller) CreateEvent(c *fiber.Ctx) error {
	auth_user := ParseAuthContext(c.Context())
	var input types.CreateEvent
	if c.BodyParser(&input) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}
	if err := ctr.Validator.Struct(input); err != nil {
		return utils.FailResponse(c, err.Error())
	}

	tx, err := ctr.DB.BeginTx(c.Context(), nil)
	if err != nil {
		tx.Rollback()
		return utils.FailResponse(c, "could not process", "error with database transaction")
	}
	q := ctr.Querier.WithTx(tx)
	defer q.Close()

	new_event, err := q.CreateEvent(c.Context(), database.CreateEventParams{
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
		return utils.FailResponse(c, "could not create event")
	}

	found_creator_participant := false
	var creator_participant types.CreateParticipant
	participants := make([]types.Participant, len(input.Participants))
	for i, p := range input.Participants {
		data := database.CreateParticipantParams{
			Name: p.Name,
			Email: p.Email,
			Organizer: p.Organizer,
			Participates: p.Participates,
			Accepted: false,
			EventID: new_event.ID,
			Address: sql.NullString{
				Valid: p.Address != "",
				String: p.Address,
			},
		}
		if p.Organizer && p.Email == auth_user.User.Email {
			found_creator_participant = true
			creator_participant = p
			data.UserID = sql.NullInt64{
				Valid: true,
				Int64: auth_user.User.ID,
			}
			data.Accepted = true
		}

		new_participant, err := q.CreateParticipant(c.Context(), data)
		if err != nil {
			tx.Rollback()
			return utils.FailResponse(c, "could not create participant", p.Name, p.Email)
		}
		participants[i] = types.Participant{
			ID: new_participant.ID,
			Name: new_participant.Name,
			Email: new_participant.Email,
			Address: new_participant.Address,
			Organizer: new_participant.Organizer,
			Participates: new_participant.Participates,
			Accepted: new_participant.Accepted,
			EventID: new_event.ID,
			UserID: new_participant.UserID.Int64,
		}
	}
	if !found_creator_participant {
		tx.Rollback()
		return utils.FailResponse(c, fmt.Sprintf(
			"%s (%s) must have the organizer field set to 'true'",
			creator_participant.Name, 
			creator_participant.Email,
		))
	}
	
	// commit all changes! create event, and all participants
	if tx.Commit() != nil {
		return utils.FailResponse(c, "could not commit transaction")
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
	return utils.DataResponseCreated(c, mapped_event)
}
