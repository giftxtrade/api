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
  "modified",
  "url",
  "origin",
  "category_id"
) VALUES (
	$1, $2, $3, $4, $5, $6, $7, sqlc.narg(currency), $8, $9, $10, $11
) RETURNING *;

-- name: FilterProducts :many
SELECT
  sqlc.embed(product),
  sqlc.embed(category)
FROM "product"
INNER JOIN "category" 
  ON "category"."id" = "product"."category_id"
WHERE
  "product"."product_ts" @@ to_tsquery('english', sqlc.arg(search))
ORDER BY
  "product"."total_reviews" DESC,
  "product"."rating" DESC
LIMIT $1
OFFSET $2;

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
