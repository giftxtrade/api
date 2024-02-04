package controllers

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/giftxtrade/api/src/database"
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
	
	products, err := ctr.Querier.FilterProducts(c.Context(), database.FilterProductsParams{
		Search: sql.NullString{
			Valid: filter.Search != nil && *filter.Search != "",
			String: *filter.Search,
		},
		Limit: filter.Limit,
		Page: filter.Page,
		MaxPrice: fmt.Sprintf("$%.2f", filter.MaxPrice),
		MinPrice: fmt.Sprintf("$%.2f", filter.MinPrice),
		SortByPrice: sql.NullBool{
			Valid: *filter.Sort == "price",
			Bool: *filter.Sort == "price",
		},
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
