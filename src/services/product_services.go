package services

import "github.com/giftxtrade/api/src/types"

type ProductServices struct {
	Service
	CategoryServices *CategoryServices
}

func (service *ProductServices) Create(create_product *types.CreateProduct) types.Product {
	new_product := types.Product{
		Title: create_product.Title,
		Description: create_product.Description,
		ProductKey: create_product.ProductKey,
		ImageUrl: create_product.ImageUrl,
		Rating: create_product.Rating,
		Price: create_product.Price,
		OriginalUrl: create_product.OriginalUrl,
		WebsiteOrigin: create_product.WebsiteOrigin,
		TotalReviews: create_product.TotalReviews,
		CategoryId: service.CategoryServices.FindOrCreate(create_product.Category).ID,
	}
	service.DB.
		Table(service.TABLE).
		Create(&new_product)
	return new_product
}