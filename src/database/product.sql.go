// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: product.sql

package database

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO "product" (
  "title",
  "description",
  "product_key",
  "image_url",
  "total_reviews",
  "rating",
  "price",
  "currency",
  "url",
  "origin",
  "category_id"
) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $11, $8, $9, $10
) RETURNING id, title, description, product_key, image_url, total_reviews, rating, price, currency, url, category_id, created_at, updated_at, product_ts, origin
`

type CreateProductParams struct {
	Title        string           `db:"title" json:"title"`
	Description  sql.NullString   `db:"description" json:"description"`
	ProductKey   string           `db:"product_key" json:"productKey"`
	ImageUrl     string           `db:"image_url" json:"imageUrl"`
	TotalReviews int32            `db:"total_reviews" json:"totalReviews"`
	Rating       float32          `db:"rating" json:"rating"`
	Price        string           `db:"price" json:"price"`
	Url          string           `db:"url" json:"url"`
	Origin       string           `db:"origin" json:"origin"`
	CategoryID   sql.NullInt64    `db:"category_id" json:"categoryId"`
	Currency     NullCurrencyType `db:"currency" json:"currency"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.queryRow(ctx, q.createProductStmt, createProduct,
		arg.Title,
		arg.Description,
		arg.ProductKey,
		arg.ImageUrl,
		arg.TotalReviews,
		arg.Rating,
		arg.Price,
		arg.Url,
		arg.Origin,
		arg.CategoryID,
		arg.Currency,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.ProductKey,
		&i.ImageUrl,
		&i.TotalReviews,
		&i.Rating,
		&i.Price,
		&i.Currency,
		&i.Url,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductTs,
		&i.Origin,
	)
	return i, err
}

const filterProducts = `-- name: FilterProducts :many
SELECT
  product.id, product.title, product.description, product.product_key, product.image_url, product.total_reviews, product.rating, product.price, product.currency, product.url, product.category_id, product.created_at, product.updated_at, product.product_ts, product.origin,
  category.id, category.name, category.description, category.category_url, category.created_at, category.updated_at,
  CEIL("product"."total_reviews" * "product"."rating") AS "weight"
FROM "product"
INNER JOIN "category" 
  ON "category"."id" = "product"."category_id"
WHERE
  (
    $2::TEXT IS NULL OR 
    "product"."product_ts" @@ to_tsquery('english', $2::TEXT)
  ) AND (
    "product"."price" BETWEEN $3 AND $4
  )
ORDER BY
  CASE WHEN 
    $5::BOOLEAN IS TRUE
  THEN
    "product"."price"
  END ASC,
  "weight" DESC
LIMIT $1
OFFSET $1 * ($6::INTEGER - 1)
`

type FilterProductsParams struct {
	Limit       int32          `db:"limit" json:"limit"`
	Search      sql.NullString `db:"search" json:"search"`
	MinPrice    string         `db:"min_price" json:"minPrice"`
	MaxPrice    string         `db:"max_price" json:"maxPrice"`
	SortByPrice sql.NullBool   `db:"sort_by_price" json:"sortByPrice"`
	Page        int32          `db:"page" json:"page"`
}

type FilterProductsRow struct {
	Product  Product  `db:"product" json:"product"`
	Category Category `db:"category" json:"category"`
	Weight   float64  `db:"weight" json:"weight"`
}

func (q *Queries) FilterProducts(ctx context.Context, arg FilterProductsParams) ([]FilterProductsRow, error) {
	rows, err := q.query(ctx, q.filterProductsStmt, filterProducts,
		arg.Limit,
		arg.Search,
		arg.MinPrice,
		arg.MaxPrice,
		arg.SortByPrice,
		arg.Page,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FilterProductsRow
	for rows.Next() {
		var i FilterProductsRow
		if err := rows.Scan(
			&i.Product.ID,
			&i.Product.Title,
			&i.Product.Description,
			&i.Product.ProductKey,
			&i.Product.ImageUrl,
			&i.Product.TotalReviews,
			&i.Product.Rating,
			&i.Product.Price,
			&i.Product.Currency,
			&i.Product.Url,
			&i.Product.CategoryID,
			&i.Product.CreatedAt,
			&i.Product.UpdatedAt,
			&i.Product.ProductTs,
			&i.Product.Origin,
			&i.Category.ID,
			&i.Category.Name,
			&i.Category.Description,
			&i.Category.CategoryUrl,
			&i.Category.CreatedAt,
			&i.Category.UpdatedAt,
			&i.Weight,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findProductById = `-- name: FindProductById :one
SELECT id, title, description, product_key, image_url, total_reviews, rating, price, currency, url, category_id, created_at, updated_at, product_ts, origin FROM "product"
WHERE "id" = $1
`

func (q *Queries) FindProductById(ctx context.Context, id int64) (Product, error) {
	row := q.queryRow(ctx, q.findProductByIdStmt, findProductById, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.ProductKey,
		&i.ImageUrl,
		&i.TotalReviews,
		&i.Rating,
		&i.Price,
		&i.Currency,
		&i.Url,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductTs,
		&i.Origin,
	)
	return i, err
}

const findProductByProductKey = `-- name: FindProductByProductKey :one
SELECT id, title, description, product_key, image_url, total_reviews, rating, price, currency, url, category_id, created_at, updated_at, product_ts, origin FROM "product"
WHERE "product_key" = $1
`

func (q *Queries) FindProductByProductKey(ctx context.Context, productKey string) (Product, error) {
	row := q.queryRow(ctx, q.findProductByProductKeyStmt, findProductByProductKey, productKey)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.ProductKey,
		&i.ImageUrl,
		&i.TotalReviews,
		&i.Rating,
		&i.Price,
		&i.Currency,
		&i.Url,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductTs,
		&i.Origin,
	)
	return i, err
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE "product"
SET 
  "price" = coalesce($2, "price"),
  "rating" = coalesce($3, "rating"),
  "total_reviews" = coalesce($4, "total_reviews"),
  "title" = coalesce($5, "title"),
  "image_url" = coalesce($6, "image_url"),
  "description" = coalesce($7, "description"),
  "updated_at" = now()
WHERE "product_key" = $1
RETURNING id, title, description, product_key, image_url, total_reviews, rating, price, currency, url, category_id, created_at, updated_at, product_ts, origin
`

type UpdateProductParams struct {
	ProductKey   string          `db:"product_key" json:"productKey"`
	Price        sql.NullString  `db:"price" json:"price"`
	Rating       sql.NullFloat64 `db:"rating" json:"rating"`
	TotalReviews sql.NullInt32   `db:"total_reviews" json:"totalReviews"`
	Title        sql.NullString  `db:"title" json:"title"`
	ImageUrl     sql.NullString  `db:"image_url" json:"imageUrl"`
	Description  sql.NullString  `db:"description" json:"description"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.queryRow(ctx, q.updateProductStmt, updateProduct,
		arg.ProductKey,
		arg.Price,
		arg.Rating,
		arg.TotalReviews,
		arg.Title,
		arg.ImageUrl,
		arg.Description,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.ProductKey,
		&i.ImageUrl,
		&i.TotalReviews,
		&i.Rating,
		&i.Price,
		&i.Currency,
		&i.Url,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductTs,
		&i.Origin,
	)
	return i, err
}
