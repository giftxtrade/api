package controllers

import (
	"strconv"
	"strings"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

// [GET] /products
func (ctr Controller) FindAllProducts(c *fiber.Ctx) error {
	var filter types.ProductFilter
	if c.BodyParser(&filter) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}
	if err := ctr.Validator.Struct(filter); err != nil {
		return utils.FailResponse(c, err.Error())
	}
	
	products, err := ctr.Querier.FilterProducts(c.Context(), database.FilterProductsParams{
		Search: filter.Search,
		Limit: filter.Limit,
		Page: filter.Page,
	})
	if err != nil {
		errors := strings.Split(err.Error(), "\n")
		return utils.FailResponse(c, errors...)
	}
	mapped_products := make([]types.Product, len(products))
	for i, p := range products {
		mapped_products[i] = types.Product{
			ID: p.Product.ID,
			Title: p.Product.Title,
			Description: p.Product.Description.String,
			ProductKey: p.Product.ProductKey,
			ImageUrl: p.Product.ImageUrl,
			TotalReviews: p.Product.TotalReviews,
			Rating: p.Product.Rating,
			Price: p.Product.Price,
			Currency: string(p.Product.Currency),
			Url: p.Product.Url,
			CategoryID: p.Product.CategoryID.Int64,
			Category: types.Category{
				ID: p.Category.ID,
				Name: p.Category.Name,
				Description: p.Category.Description.String,
			},
			CreatedAt: p.Product.CreatedAt,
			UpdatedAt: p.Product.UpdatedAt,
			Origin: p.Product.Origin,
		}
	}
	return utils.DataResponse(c, mapped_products)
}

// [POST] /products
func (ctr Controller) CreateProduct(c *fiber.Ctx) error {
	var create_product types.CreateProduct
	if c.BodyParser(&create_product) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}

	product, created, err := ctr.Service.ProductService.UpdateOrCreate(c.Context(), create_product)
	if err != nil {
		return utils.FailResponse(c, "could not create/update product")
	}
	if created {
		return utils.DataResponseCreated(c, product)
	}
	return utils.DataResponse(c, product)
}

// [GET] /products/:id
func (ctr Controller) FindProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.FailResponse(c, "invalid product id")
	}

	product, err := ctr.Querier.FindProductById(c.Context(), int64(id))
	if err != nil {
		return utils.FailResponse(c, "product not found")
	}
	return utils.DataResponse(c, product)
}
