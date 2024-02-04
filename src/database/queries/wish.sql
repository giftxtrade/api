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
