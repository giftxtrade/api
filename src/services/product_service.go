package services

import (
	"context"
	"database/sql"
	"net/url"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
)

type ProductService struct {
	ServiceBase
}

func (service *ProductService) UpdateOrCreate(ctx context.Context, input types.CreateProduct) (database.Product, bool, error) {
	validation_err := service.Validator.Struct(input)
	if validation_err != nil {
		return database.Product{}, false, validation_err
	}

	found_product, err := service.
		Querier.
		FindProductByProductKey(ctx, input.ProductKey)
	// create new product
	if err != nil {
		parsed_url, url_parse_err := url.ParseRequestURI(input.OriginalUrl)
		if url_parse_err != nil {
			return database.Product{}, false, url_parse_err
		}

		category, category_err := service.FindOrCreateCategory(ctx, database.CreateCategoryParams{
			Name: input.Category,
		})
		if category_err != nil {
			return database.Product{}, false, category_err
		}
		product, err := service.Querier.CreateProduct(ctx, database.CreateProductParams{
			ProductKey: input.ProductKey,
			Title: input.Title,
			Description: sql.NullString{
				String: input.Description,
				Valid: input.Description != "",
			},
			ImageUrl: input.ImageUrl,
			TotalReviews: int32(input.TotalReviews),
			Rating: input.Rating,
			Price: input.Price,
			Url: input.OriginalUrl,
			Origin: parsed_url.Host,
			Currency: database.NullCurrencyType{
				CurrencyType: database.CurrencyTypeUSD,
				Valid: true,
			},
			CategoryID: sql.NullInt64{
				Int64: category.ID,
				Valid: true,
			},
		})
		return product, err == nil, err
	}
	
	// update existing product
	product, err := service.Querier.UpdateProduct(ctx, database.UpdateProductParams{
		ProductKey: input.ProductKey,
		Rating: sql.NullFloat64{
			Float64: float64(input.Rating),
			Valid: input.Rating != 0 && found_product.Rating != input.Rating,
		},
		TotalReviews: sql.NullInt32{
			Int32: int32(input.TotalReviews),
			Valid: input.TotalReviews != 0 && found_product.TotalReviews != int32(input.TotalReviews),
		},
		Price: sql.NullString{
			String: input.Price,
			Valid: input.Price != "" && found_product.Price != input.Price,
		},
		Title: sql.NullString{
			String: input.Title,
			Valid: input.Title != "" && found_product.Title != input.Title,
		},
		ImageUrl: sql.NullString{
			String: input.ImageUrl,
			Valid: input.ImageUrl != "" && found_product.ImageUrl != input.ImageUrl,
		},
		Description: sql.NullString{
			String: input.Description,
			Valid: input.Description != "" && found_product.Description.String != input.Description,
		},
	})
	return product, false, err
}

func (service *ProductService) FindOrCreateCategory(ctx context.Context, input database.CreateCategoryParams) (database.Category, error) {
	found_category, err := service.Querier.FindCategoryByName(ctx, input.Name)
	if err != nil {
		return service.Querier.CreateCategory(ctx, input)
	}
	return found_category, nil
}
