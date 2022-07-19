package controllers

import (
	"strings"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctx Controller) FindAllProducts(c *fiber.Ctx) error {
	var filter types.ProductFilter
	if c.BodyParser(&filter) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}
	
	products, err := ctx.
		Service.
		ProductService.
		Search(filter)
	if err != nil {
		errors := strings.Split(err.Error(), "\n")
		return utils.FailResponse(c, errors...)
	}
	return utils.DataResponse(c, products)
}

func (ctx Controller) CreateProduct(c *fiber.Ctx) error {
	var create_product types.CreateProduct
	if c.BodyParser(&create_product) != nil {
		return utils.FailResponse(c, "could not parse body data")
	}

	var new_product types.Product
	_, err := ctx.Service.ProductService.CreateOrUpdate(&create_product, &new_product)
	if err != nil {
		return utils.FailResponse(c, strings.Split(err.Error(), "\n")...)
	}
	return utils.DataResponse(c, new_product)
}

func (ctx Controller) FindProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product types.Product
	if ctx.Service.ProductService.Find(id, &product) != nil {
		return utils.FailResponse(c, "product not found")
	}
	return utils.DataResponse(c, product)
}