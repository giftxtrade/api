-- name: CreateWish :one
INSERT INTO wish (
    user_id,
    participant_id,
    product_id,
    event_id,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, now(), now()
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
