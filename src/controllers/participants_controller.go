package controllers

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/mappers"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctr *Controller) ManageParticipantUpdate(c *fiber.Ctx) error {
	auth := GetAuthContext(c.UserContext())
	event_id := GetEventIdFromContext(c.UserContext())
	participant_id, err := strconv.ParseInt(c.Query("participantId"), 10, 64)
	if err != nil {
		return utils.FailResponse(c, "invalid participant id")
	}

	// verify if participant exists in event
	participant, err := ctr.Querier.FindParticipantWithIdAndEventId(c.Context(), database.FindParticipantWithIdAndEventIdParams{
		EventID: event_id,
		ParticipantID: participant_id,
	})
	if err != nil {
		return utils.FailResponse(c, "participant does not exist on the event")
	}

	// parse/validate body
	var input types.PatchParticipant
	if json.Unmarshal(c.Body(), &input) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}
	if err := ctr.Validator.StructCtx(c.Context(), input); err != nil {
		return utils.FailResponse(c, err.Error())
	}
	
	patch_data := database.UpdateParticipantStatusParams{
		EventID: event_id,
		ParticipantID: participant_id,
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
