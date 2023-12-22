-- name: CreateParticipant :one
INSERT INTO "participant" (
    "name",
    "email",
    "address",
    "organizer",
    "participates",
    "accepted",
    "event_id",
    "user_id"
) VALUES (
    $1, $2, sqlc.narg(address), $3, $4, $5, $6, sqlc.narg(user_id)
)
RETURNING *;
