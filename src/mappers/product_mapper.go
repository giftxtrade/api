package mappers

import (
	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
)

func DbProductToProduct(product database.Product, category *database.Category) types.Product {
	result := types.Product{
		ID: product.ID,
		Title: product.Title,
		Description: product.Description.String,
		ProductKey: product.ProductKey,
		ImageUrl: product.ImageUrl,
		TotalReviews: product.TotalReviews,
		Rating: product.Rating,
		Price: product.Price,
		Currency: string(product.Currency),
		Url: product.Url,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		Origin: product.Origin,
		CategoryID: product.CategoryID.Int64,
	}
	if category != nil {
		result.CategoryID = category.ID
		result.Category = types.Category{
			ID: category.ID,
			Name: category.Name,
			Description: category.Description.String,
		}
	}
	return result
}
