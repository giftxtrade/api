package services

import (
	"net/url"

	"github.com/giftxtrade/api/src/types"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductService struct {
	ServiceBase
	CategoryService CategoryService
}

func (service *ProductService) Create(input *types.CreateProduct, output *types.Product) error {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return err
	}

	var category types.Category
	_, category_err := service.CategoryService.FindOrCreate(input.Category, &category)
	if category_err != nil {
		return category_err
	}

	output.Title = input.Title
	output.Description = input.Description
	output.ProductKey = input.ProductKey
	output.ImageUrl = input.ImageUrl
	output.Rating = input.Rating
	output.Price = input.Price
	output.OriginalUrl = input.OriginalUrl
	output.TotalReviews = input.TotalReviews
	output.CategoryId = category.ID
	output.Category = category
	// add website origin
	parsed_url, err := url.ParseRequestURI(input.OriginalUrl)
	if err == nil {
		output.WebsiteOrigin = parsed_url.Host
	} else {
		return err
	}

	return service.DB.
		Table(service.TABLE).
		Create(output).
		Error
}

func (service *ProductService) Find(key string, output *types.Product) error {
	id, _ := uuid.Parse(key)
	return service.DB.
		Table(service.TABLE).
		Preload("Category").
		Where("products.product_key = ? OR products.id = ?", key, id).
		First(output).
		Error
}

// create a new product or update existing product with input
// boolean value is true if a new user is created, otherwise false
func (service *ProductService) CreateOrUpdate(input *types.CreateProduct, output *types.Product) (bool, error) {
	if service.Find(input.ProductKey, output) != nil {
		create_err := service.Create(input, output)
		if create_err == nil {
			return true, nil
		}
		return false, create_err
	}
	
	// product already exists, so update...
	changed := false
	if input.Title != output.Title {
		output.Title = input.Title
		changed = true
	}
	if input.Description != output.Description {
		output.Description = input.Description
		changed = true
	}
	if input.ImageUrl != output.ImageUrl {
		output.ImageUrl = input.ImageUrl
		changed = true
	}
	if input.Price != output.Price {
		output.Price = input.Price
		changed = true
	}
	if input.Rating != output.Rating {
		output.Rating = input.Rating
		changed = true
	}
	if input.TotalReviews != output.TotalReviews {
		output.TotalReviews = input.TotalReviews
		changed = true
	}
	if input.Category != output.Category.Name {
		var new_category types.Category
		_, category_err := service.CategoryService.FindOrCreate(input.Category, &new_category)
		if category_err == nil {
			output.CategoryId = new_category.ID
			output.Category = new_category
			changed = true
		}
	}

	var err error
	if changed {
		err = service.DB.
			Save(output).
			Error
	}
	return false, err
}

func (service *ProductService) Search(filter types.ProductFilter) (*[]types.Product, error) {
	validate := validator.New()
	if err := validate.Struct(filter); err != nil {
		return nil, err
	}

	products := new([]types.Product)
	offset := (filter.Page - 1) * filter.Limit
	query := service.DB.
		Limit(filter.Limit).
		Offset(offset)
	
	if filter.MinPrice > 0 || filter.MaxPrice > 0 {
		query.
			Where("price BETWEEN ? AND ?", filter.MinPrice, filter.MaxPrice)
	}

	switch filter.Sort {
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