package services

import (
	"net/url"

	"github.com/giftxtrade/api/src/types"
	"github.com/google/uuid"
)

type ProductService struct {
	Service
	CategoryServices *CategoryService
}

func (service *ProductService) Create(create_product *types.CreateProduct) (*types.Product, error) {
	category, _, category_err := service.CategoryServices.FindOrCreate(create_product.Category)
	if category_err != nil {
		return nil, category_err
	}
	new_product := types.Product{
		Title: create_product.Title,
		Description: create_product.Description,
		ProductKey: create_product.ProductKey,
		ImageUrl: create_product.ImageUrl,
		Rating: create_product.Rating,
		Price: create_product.Price,
		OriginalUrl: create_product.OriginalUrl,
		TotalReviews: create_product.TotalReviews,
		CategoryId: category.ID,
		Category: *category,
	}

	// add website origin
	if create_product.OriginalUrl != "" {
		parsed_url, err := url.ParseRequestURI(create_product.OriginalUrl)
		if err == nil {
			new_product.WebsiteOrigin = parsed_url.Host
		} else {
			return nil, err
		}
	}
	err := service.DB.
		Table(service.TABLE).
		Create(&new_product).
		Error
	return &new_product, err
}

func (service *ProductService) Find(key string) (*types.Product, error) {
	id, _ := uuid.Parse(key)
	var product types.Product
	err := service.DB.
		Table(service.TABLE).
		Preload("Category").
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("products.product_key = ? OR products.id = ?", key, id).
		First(&product).
		Error
	return &product, err
}

func (service *ProductService) CreateOrUpdate(create_product *types.CreateProduct) (*types.Product, error) {
	product, err := service.Find(create_product.ProductKey)
	if err != nil {
		return service.Create(create_product)
	}
	
	// product already exists, so update...
	changed := false
	if create_product.Title != product.Title {
		product.Title = create_product.Title
		changed = true
	}
	if create_product.Description != product.Description {
		product.Description = create_product.Description
		changed = true
	}
	if create_product.ImageUrl != product.ImageUrl {
		product.ImageUrl = create_product.ImageUrl
		changed = true
	}
	if create_product.Price != product.Price {
		product.Price = create_product.Price
		changed = true
	}
	if create_product.Rating != product.Rating {
		product.Rating = create_product.Rating
		changed = true
	}
	if create_product.TotalReviews != product.TotalReviews {
		product.TotalReviews = create_product.TotalReviews
		changed = true
	}
	if create_product.Category != product.Category.Name {
		new_category, _, category_err := service.CategoryServices.FindOrCreate(create_product.Category)
		if category_err == nil {
			product.CategoryId = new_category.ID
			product.Category = *new_category
			changed = true
		}
	}

	if changed {
		err = service.DB.
			Save(product).
			Error
	}
	return product, err
}