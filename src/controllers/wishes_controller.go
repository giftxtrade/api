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
		// check if product id is valid
		product_id := *input.ProductID
		p, err := ctr.Querier.FindProductById(c.Context(), product_id)
		if err != nil {
			return utils.FailResponse(c, "invalid product id")
		}
		product = &p
		create_wish_params.ProductID = sql.NullInt64{
			Int64: product_id,
			Valid: true,
		}

		// check if wish with `product_id` already exists and update `quantity`
		existing_wish, existing_wish_err := ctr.Querier.GetWishWithProductID(c.Context(), database.GetWishWithProductIDParams{
			UserID: auth.User.ID,
			EventID: event_id,
			ParticipantID: participant.ID,
			ProductID: sql.NullInt64{
				Valid: true,
				Int64: product_id,
			},
		})
		if existing_wish_err == nil {
			updated_wish, update_wish_err := ctr.Querier.UpdateWishQuantity(c.Context(), database.UpdateWishQuantityParams{
				ID: existing_wish.ID,
				Quantity: existing_wish.Quantity + 1,
			})
			if update_wish_err != nil {
				return utils.FailResponse(c, "could not update wish quantity")
			}
			return utils.DataResponse(c, mappers.DbWishToWish(updated_wish, product))
		}
	}

	wish, err := ctr.Querier.CreateWish(c.Context(), create_wish_params)
	if err != nil {
		return utils.FailResponse(c, "could not create wish")
	}
	return utils.DataResponseCreated(c, mappers.DbWishToWish(wish, product))
}

func (ctr *Controller) DeleteWish(c *fiber.Ctx) error {
	input, err := utils.ParseAndValidateBody[types.DeleteWish](ctr.Validator, c.Body())
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

	wish, err := ctr.Querier.GetWishByAllIDs(c.Context(), database.GetWishByAllIDsParams{
		ID: input.WishID,
		UserID: auth.User.ID,
		ParticipantID: participant.ID,
		EventID: event_id,
	})
	if err != nil {
		return utils.FailResponse(c, "could not find wish with the given inputs")
	}

	_, err = ctr.Querier.DeleteWish(c.Context(), wish.ID)
	if err != nil {
		return utils.FailResponse(c, "could not delete wish")
	}
	return utils.DataResponse(c, mappers.DbWishToWish(wish, nil))
}

func (ctr *Controller) GetWishes(c *fiber.Ctx) error {
	auth := GetAuthContext(c.UserContext())
	event_id := GetEventIdFromContext(c.UserContext())
	participant := GetParticipantFromContext(c.UserContext())
	wishes, err := ctr.Querier.GetAllWishesForUser(c.Context(), database.GetAllWishesForUserParams{
		UserID: auth.User.ID,
		EventID: event_id,
		ParticipantID: participant.ID,
	})
	if err != nil {
		return utils.FailResponse(c, "could not fetch wishes")
	}

	mapped_wishes := make([]types.Wish, len(wishes))
	for i, w := range wishes {
		mapped_wishes[i] = mappers.DbWishToWish(w.Wish, &w.Product)
	}
	return utils.DataResponse(c, mapped_wishes)
}
