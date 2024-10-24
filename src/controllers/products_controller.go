package controllers

import (
	"strconv"
	"strings"

	"github.com/giftxtrade/api/src/mappers"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

// [GET] /products
func (ctr Controller) FindAllProducts(c *fiber.Ctx) error {
	search_query := c.Query("search")
	filter := types.ProductFilter{
		Search: &search_query,
		Limit: int32(c.QueryInt("limit")),
		Page: int32(c.QueryInt("page")),
		MinPrice: float32(c.QueryFloat("minPrice")),
		MaxPrice: float32(c.QueryFloat("maxPrice")),
	}
	if sort := c.Query("sort"); sort != "" {
		value := ""
		switch sort {
		case "price":
			value = "price"
		case "rating":
			value = "rating"
		default:
			return utils.FailResponse(c, "invalid value for param 'sort'")
		}
		filter.Sort = &value
	}
	if err := ctr.Validator.Struct(filter); err != nil {
		return utils.FailResponse(c, err.Error())
	}
	
	products, err := ctr.Service.ProductService.Search(c.Context(), filter)
	if err != nil {
		errors := strings.Split(err.Error(), "\n")
		return utils.FailResponse(c, errors...)
	}
	return utils.DataResponse(c, products)
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
	mapped_product := mappers.DbProductToProduct(product, nil)
	if created {
		return utils.DataResponseCreated(c, mapped_product)
	}
	return utils.DataResponse(c, mapped_product)
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
	return utils.DataResponse(c, mappers.DbProductToProduct(product, nil))
}
