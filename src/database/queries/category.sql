-- name: FindCategoryByName :one
SELECT * FROM "category"
WHERE "name" = $1;

-- name: CreateCategory :one
INSERT INTO "category" (
    "name",
    "description",
    "category_url"
) VALUES(
    $1,
    sqlc.narg(description),
    sqlc.narg(category_url)
)
RETURNING *;
