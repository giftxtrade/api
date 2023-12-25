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

-- name: FindAllEventsWithUser :many
SELECT
    sqlc.embed(event),
    sqlc.embed(p2),
    sqlc.embed(u)
FROM "event"
JOIN "participant" "p1" ON "p1"."event_id" = "event"."id"
JOIN "participant" "p2" ON "p2"."event_id" = "event"."id"
JOIN "user" "u" ON "u"."id" = "p2"."user_id"
WHERE 
    "p1"."user_id" = $1
ORDER BY
    "event"."draw_at" DESC,
    "event"."close_at" DESC;
