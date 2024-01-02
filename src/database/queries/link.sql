-- name: CreateLink :one
INSERT INTO "link" (
    code,
    expiration_date,
    event_id
) VALUES (
    $1, $2, $3
) RETURNING *;
