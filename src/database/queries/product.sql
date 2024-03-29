-- name: FindProductById :one
SELECT * FROM "product"
WHERE "id" = $1;

-- name: FindProductByProductKey :one
SELECT * FROM "product"
WHERE "product_key" = $1;

-- name: CreateProduct :one
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
	$1, $2, $3, $4, $5, $6, $7, sqlc.narg(currency), $8, $9, $10
) RETURNING *;

-- name: FilterProducts :many
SELECT
  sqlc.embed(product),
  sqlc.embed(category),
  CEIL("product"."total_reviews" * "product"."rating") AS "weight"
FROM "product"
INNER JOIN "category" 
  ON "category"."id" = "product"."category_id"
WHERE
  (
    sqlc.narg(search)::TEXT IS NULL OR 
    "product"."product_ts" @@ to_tsquery('english', sqlc.narg(search)::TEXT)
  ) AND (
    "product"."price" BETWEEN sqlc.arg(min_price) AND sqlc.arg(max_price)
  )
ORDER BY
  CASE WHEN 
    sqlc.narg(sort_by_price)::BOOLEAN IS TRUE
  THEN
    "product"."price"
  END ASC,
  "weight" DESC
LIMIT $1
OFFSET $1 * (sqlc.arg(page)::INTEGER - 1);

-- name: UpdateProduct :one
UPDATE "product"
SET 
  "price" = coalesce(sqlc.narg('price'), "price"),
  "rating" = coalesce(sqlc.narg('rating'), "rating"),
  "total_reviews" = coalesce(sqlc.narg('total_reviews'), "total_reviews"),
  "title" = coalesce(sqlc.narg('title'), "title"),
  "image_url" = coalesce(sqlc.narg('image_url'), "image_url"),
  "description" = coalesce(sqlc.narg('description'), "description"),
  "updated_at" = now()
WHERE "product_key" = $1
RETURNING *;
