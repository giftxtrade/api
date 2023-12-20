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
func (ctx Controller) FindAllProducts(c *fiber.Ctx) error {
	var filter types.ProductFilter
	if c.BodyParser(&filter) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}
	
	products, err := ctx.Querier.FilterProducts(c.Context(), database.FilterProductsParams{
		Search: filter.Search,
		Limit: filter.Limit,
		Page: filter.Page,
	})
	if err != nil {
		errors := strings.Split(err.Error(), "\n")
		return utils.FailResponse(c, errors...)
	}
	return utils.DataResponse(c, products)
}

// [POST] /products
func (ctx Controller) CreateProduct(c *fiber.Ctx) error {
	var create_product types.CreateProduct
	if c.BodyParser(&create_product) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}

	product, created, err := ctx.Service.ProductService.UpdateOrCreate(c.Context(), create_product)
	if err != nil {
		return utils.FailResponse(c, "could not create/update product")
	}
	if created {
		return utils.DataResponseCreated(c, product)
	}
	return utils.DataResponse(c, product)
}

// [GET] /products/:id
func (ctx Controller) FindProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.FailResponse(c, "invalid product id")
	}

	product, err := ctx.Querier.FindProductById(c.Context(), int64(id))
	if err != nil {
		return utils.FailResponse(c, "product not found")
	}
	return utils.DataResponse(c, product)
}