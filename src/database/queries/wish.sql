-- name: CreateWish :one
INSERT INTO wish (
    user_id,
    participant_id,
    product_id,
    event_id,
    quantity,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, COALESCE(sqlc.narg(quantity)::INTEGER, 1), now(), now()
) RETURNING *;

-- name: DeleteWish :one
DELETE FROM wish 
WHERE id = $1
RETURNING id;

-- name: GetWishByAllIDs :one
SELECT *
FROM wish 
WHERE
    id = $1 AND
    user_id = $2 AND
    participant_id = $3 AND
    event_id = $4;

-- name: GetAllWishesForUser :many
SELECT
    sqlc.embed(wish),
    sqlc.embed(product)
FROM wish
INNER JOIN product ON product.id = wish.product_id
WHERE
    wish.user_id = $1 AND
    wish.participant_id = $2 AND
    wish.event_id = $3
ORDER BY wish.created_at DESC;

-- name: GetWishWithProductID :one
SELECT * FROM wish
WHERE
    user_id = $1 AND 
    event_id = $2 AND
    participant_id = $3 AND
    product_id = $4;

-- name: UpdateWishQuantity :one
UPDATE wish
SET
    quantity = $2, 
    updated_at = NOW()
WHERE id = $1
RETURNING *;
