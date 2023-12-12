-- name: FindUserByEmail :one
SELECT * FROM "user"
WHERE email = $1;

-- name: FindUserById :one
SELECT * FROM "user"
WHERE id = $1;

-- name: FindUserByIdAndEmail :one
SELECT * FROM "user"
WHERE id = $1 AND email = $2;

-- name: FindUserByIdOrEmail :one
SELECT * FROM "user"
WHERE id = $1 OR email = $2;

-- name: CreateUser :one
INSERT INTO "user" (
    name,
    email,
    image_url,
    phone,
    admin,
    active
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;
