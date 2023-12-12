-- name: GetAllUsers :many
SELECT * FROM "user"
LIMIT $1
OFFSET $2;
