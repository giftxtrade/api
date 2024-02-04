package controllers

import (
	"database/sql"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/mappers"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctr *Controller) CreateWish(c *fiber.Ctx) error {
	input, err := utils.ParseAndValidateBody[types.CreateWish](ctr.Validator, c.Body())
	if err != nil {
		return utils.FailResponse(c, err.Error())
	}

	auth := GetAuthContext(c.UserContext())
	event_id := GetEventIdFromContext(c.UserContext())
	participant, err := ctr.Querier.FindParticipantFromEventIdAndUser(c.Context(), database.FindParticipantFromEventIdAndUserParams{
		EventID: event_id,
		UserID: sql.NullInt64{
			Int64: auth.User.ID,
			Valid: true,
		},
	})
	if err != nil {
		return utils.FailResponse(c, "participant does not exist on the event")
	}
	create_wish_params := database.CreateWishParams{
		UserID: auth.User.ID,
		EventID: event_id,
		ParticipantID: participant.ID,
	}
	var product *database.Product = nil
	if input.ProductID != nil {
		p, err := ctr.Querier.FindProductById(c.Context(), *input.ProductID)
		if err != nil {
			return utils.FailResponse(c, "invalid product id")
		}
		product = &p
		create_wish_params.ProductID = sql.NullInt64{
			Int64: *input.ProductID,
			Valid: true,
		}
	}
	wish, err := ctr.Querier.CreateWish(c.Context(), create_wish_params)
	if err != nil {
		return utils.FailResponse(c, "could not create wish")
	}
	return utils.DataResponseCreated(c, mappers.DbWishToWish(wish, product))
}
