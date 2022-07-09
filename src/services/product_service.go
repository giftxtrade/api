package services

import (
	"net/url"

	"github.com/giftxtrade/api/src/types"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductService struct {
	*Service
	CategoryService *CategoryService
}

func (service *ProductService) Create(create_product *types.CreateProduct, product *types.Product) error {
	validate := validator.New()
	if err := validate.Struct(create_product); err != nil {
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