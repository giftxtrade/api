package controllers

import (
	"database/sql"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/mappers"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctr *Controller) ManageParticipantUpdate(c *fiber.Ctx) error {
	auth := GetAuthContext(c.UserContext())
	event_id := GetEventIdFromContext(c.UserContext())
	participant := c.UserContext().Value(PARTICIPANT_OB_KEY).(database.Participant)

	// parse/validate body
	input, err := utils.ParseAndValidateBody[types.PatchParticipant](ctr.Validator, c.Body())
	if err != nil {
		return utils.FailResponse(c, err.Error())
	}
	
	patch_data := database.UpdateParticipantStatusParams{
		EventID: event_id,
		ParticipantID: participant.ID,
	}
	if input.Organizer != nil {
		if participant.UserID.Int64 == auth.User.ID {
			return utils.FailResponse(c, "event organizer cannot modify their own 'organizer' status")
		}
		patch_data.Organizer = sql.NullBool{
			Valid: *input.Organizer != participant.Organizer,
			Bool: *input.Organizer,
		}
	}
	if input.Participates != nil {
		patch_data.Participates = sql.NullBool{
			Valid: *input.Participates != participant.Participates,
			Bool: *input.Participates,
		}
	}
	patched_participant, err := ctr.Querier.UpdateParticipantStatus(c.Context(), patch_data)
	if err != nil {
		utils.FailResponse(c, "could not update participant")
	}
	return utils.DataResponse(c, mappers.DbParticipantToParticipant(patched_participant, nil, nil))
}

func (ctr *Controller) ManageParticipantRemoval(c *fiber.Ctx) error {
	auth := GetAuthContext(c.UserContext())
	event_id := GetEventIdFromContext(c.UserContext())
	participant := c.UserContext().Value(PARTICIPANT_OB_KEY).(database.Participant)
	if participant.UserID.Int64 == auth.User.ID {
		utils.FailResponse(c, "event organizer cannot remove themselves")
	}
	_, err := ctr.Querier.DeleteParticipantByIdAndEventId(c.Context(), database.DeleteParticipantByIdAndEventIdParams{
		EventID: event_id,
		ParticipantID: participant.ID,
	})
	if err != nil {
		return utils.FailResponse(c, "could not remove participant")
	}
	return utils.DataResponse(c, mappers.DbParticipantToParticipant(participant, nil, nil))
}

func (ctr *Controller) GetParticipantById(c *fiber.Ctx) error {
	participant := c.UserContext().Value(PARTICIPANT_OB_KEY).(database.Participant)
	rows, err := ctr.Querier.FindParticipantUserWithFullEventById(c.Context(), participant.ID)
	if err != nil {
		return utils.FailResponseNotFound(c, "could not find participant with the id", err.Error())
	}
	if len(rows) <= 0 {
		return utils.FailResponse(c, "result error")
	}
	mapped_participant := mappers.DbParticipantUserToParticipant(rows[0].ParticipantUser, &rows[0].Event)
	participants := make([]types.Participant, len(rows))
	for i, row := range rows {
		participants[i] = mappers.DbParticipantUserToParticipant(row.ParticipantUser_2, nil)
	}
	mapped_participant.Event.Participants = participants
	return utils.DataResponse(c, mapped_participant)
}
