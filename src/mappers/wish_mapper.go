package mappers

import (
	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
)

func DbWishToWish(wish database.Wish, product *database.Product) types.Wish {
	result := types.Wish{
		ID: wish.ID,
		UserID: wish.UserID,
		ParticipantID: wish.ParticipantID,
		EventID: wish.EventID,
		Quantity: wish.Quantity,
	}
	if product != nil {
		result.ProductID = product.ID
		product := DbProductToProduct(*product, nil)
		result.Product = &product
	}
	return result
}
