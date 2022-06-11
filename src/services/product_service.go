package services

import (
	"fmt"
	"log"
	"net/url"

	"github.com/giftxtrade/api/src/types"
	"github.com/google/uuid"
)

type ProductService struct {
	*Service
	CategoryService *CategoryService
}

func (service *ProductService) Create(create_product *types.CreateProduct) (*types.Product, error) {
	var err error
	if create_product, err = validate_create_product_input(create_product); err != nil {
		return nil, err
	}

	category, _, category_err := service.CategoryService.FindOrCreate(create_product.Category)
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
	parsed_url, err := url.ParseRequestURI(create_product.OriginalUrl)
	if err == nil {
		new_product.WebsiteOrigin = parsed_url.Host
	} else {
		return nil, err
	}

	err = service.DB.
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
		Where("products.product_key = ? OR products.id = ?", key, id).
		First(&product).
		Error
	return &product, err
}

// create a new product or update existing product with input
// boolean value is true if a new user is created, otherwise false
func (service *ProductService) CreateOrUpdate(create_product *types.CreateProduct) (*types.Product, bool, error) {
	product, err := service.Find(create_product.ProductKey)
	if err != nil {
		create_product, create_err := service.Create(create_product)
		if create_err == nil {
			return create_product, true, nil
		}
		return nil, false, create_err
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
		new_category, _, category_err := service.CategoryService.FindOrCreate(create_product.Category)
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
	return product, false, err
}

func (service *ProductService) Search(search string, limit int, offset int, minPrice float32, maxPrice float32, sort string) (*[]types.Product, error) {
	products := new([]types.Product)
	query := service.DB.
		Order("updated_at DESC").
		Limit(limit).
		Offset(offset)
	
	if minPrice > 0 && maxPrice >= minPrice {
		query.
			Where("price BETWEEN ? AND ?", minPrice, maxPrice)
	}
	if sort != "" {
		switch sort {
		case "rating":
			log.Println("sort by: ", sort)
			query.Order("rating DESC")
		case "price":
			log.Println("sort by: ", sort)
			query.Order("price DESC")
		case "totalReviews":
			log.Println("sort by: ", sort)
			query.Order("total_reviews DESC")
		}
	}
	err := query.
		Preload("Category").
		Find(products).
		Error
	return products, err
}

func validate_create_product_input(create_product *types.CreateProduct) (*types.CreateProduct, error) {
	if create_product.Title == "" {
		return nil, fmt.Errorf("title is required")
	}
	if create_product.ProductKey == "" {
		return nil, fmt.Errorf("productKey is required")
	}
	if create_product.Rating <= 0 {
		return nil, fmt.Errorf("rating is required")
	}
	if create_product.Rating > 5 {
		return nil, fmt.Errorf("rating should be between interval (0, 5]")
	}
	if create_product.Price <= 0 {
		return nil, fmt.Errorf("price is required")
	}
	if create_product.OriginalUrl == "" {
		return nil, fmt.Errorf("originalUrl is required")
	}
	if create_product.TotalReviews == 0 {
		return nil, fmt.Errorf("totalReviews is required")
	}
	if create_product.Category == "" {
		return nil, fmt.Errorf("category is required")
	}
	return create_product, nil
}