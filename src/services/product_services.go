package services

import (
	"github.com/giftxtrade/api/src/types"
)

type ProductServices struct {
	Service
	CategoryServices *CategoryServices
}

func (service *ProductServices) Create(create_product *types.CreateProduct) types.Product {
	category := service.CategoryServices.FindOrCreate(create_product.Category)
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
		CategoryId: category.ID,
		Category: category,
	}
	service.DB.
		Table(service.TABLE).
		Create(&new_product).
		Joins(service.CategoryServices.TABLE)
	return new_product
}

func (service *ProductServices) Find(key string) types.Product {
	var product types.Product
	service.DB.
		Table(service.TABLE).
		Preload("Category").
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("products.id = ? OR products.product_key = ?", key, key).
		Find(&product)
	return product
}