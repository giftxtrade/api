package services

import (
	"fmt"
	"net/url"

	"github.com/giftxtrade/api/src/types"
	"github.com/google/uuid"
)

type ProductService struct {
	*Service
	CategoryService *CategoryService
}

func (service *ProductService) Create(create_product *types.CreateProduct, product *types.Product) error {
	if err := validate_create_product_input(create_product); err != nil {
		return err
	}

	var category types.Category
	_, category_err := service.CategoryService.FindOrCreate(create_product.Category, &category)
	if category_err != nil {
		return category_err
	}

	product.Title = create_product.Title
	product.Description = create_product.Description
	product.ProductKey = create_product.ProductKey
	product.ImageUrl = create_product.ImageUrl
	product.Rating = create_product.Rating
	product.Price = create_product.Price
	product.OriginalUrl = create_product.OriginalUrl
	product.TotalReviews = create_product.TotalReviews
	product.CategoryId = category.ID
	product.Category = category
	// add website origin
	parsed_url, err := url.ParseRequestURI(create_product.OriginalUrl)
	if err == nil {
		product.WebsiteOrigin = parsed_url.Host
	} else {
		return err
	}

	return service.DB.
		Table(service.TABLE).
		Create(product).
		Error
}

func (service *ProductService) Find(key string, product *types.Product) error {
	id, _ := uuid.Parse(key)
	return service.DB.
		Table(service.TABLE).
		Preload("Category").
		Where("products.product_key = ? OR products.id = ?", key, id).
		First(product).
		Error
}

// create a new product or update existing product with input
// boolean value is true if a new user is created, otherwise false
func (service *ProductService) CreateOrUpdate(create_product *types.CreateProduct, product *types.Product) (bool, error) {
	if service.Find(create_product.ProductKey, product) != nil {
		create_err := service.Create(create_product, product)
		if create_err == nil {
			return true, nil
		}
		return false, create_err
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
		var new_category types.Category
		_, category_err := service.CategoryService.FindOrCreate(create_product.Category, &new_category)
		if category_err == nil {
			product.CategoryId = new_category.ID
			product.Category = new_category
			changed = true
		}
	}

	var err error
	if changed {
		err = service.DB.
			Save(product).
			Error
	}
	return false, err
}

func (service *ProductService) Search(search string, limit int, page int, minPrice float32, maxPrice float32, sort string) (*[]types.Product, error) {
	offset := (page - 1) * limit
	products := new([]types.Product)
	query := service.DB.
		Limit(limit).
		Offset(offset)
	
	if minPrice > 0 && maxPrice >= minPrice {
		query.
			Where("price BETWEEN ? AND ?", minPrice, maxPrice)
	}
	switch sort {
	case "rating":
		query.Order("rating DESC")
	case "price":
		query.Order("price DESC")
	case "totalReviews":
		query.Order("total_reviews DESC")
	default:
		query.Order("updated_at DESC")
	}
	err := query.
		Preload("Category").
		Find(products).
		Error
	return products, err
}

func validate_create_product_input(create_product *types.CreateProduct) error {
	if create_product.Title == "" {
		return fmt.Errorf("title is required")
	}
	if create_product.ProductKey == "" {
		return fmt.Errorf("productKey is required")
	}
	if create_product.Rating <= 0 {
		return fmt.Errorf("rating is required")
	}
	if create_product.Rating > 5 {
		return fmt.Errorf("rating should be between interval (0, 5]")
	}
	if create_product.Price <= 0 {
		return fmt.Errorf("price is required")
	}
	if create_product.OriginalUrl == "" {
		return fmt.Errorf("originalUrl is required")
	}
	if create_product.TotalReviews == 0 {
		return fmt.Errorf("totalReviews is required")
	}
	if create_product.Category == "" {
		return fmt.Errorf("category is required")
	}
	return nil
}