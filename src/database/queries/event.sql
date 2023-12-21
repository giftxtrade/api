-- name: CreateEvent :one
INSERT INTO "event" (
    "name",
    "description",
    "budget",
    "invitation_message",
    "draw_at",
    "close_at"
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;
