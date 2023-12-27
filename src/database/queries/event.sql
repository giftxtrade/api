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
    sqlc.embed(p) AS "participant"
FROM "event"
JOIN "participant" "p1" ON "p1"."event_id" = "event"."id"
JOIN "participant_user" "p" ON "p"."event_id" = "event"."id"
WHERE 
    "p1"."user_id" = $1
ORDER BY
    "event"."draw_at" DESC,
    "event"."close_at" DESC;

-- name: FindEventForUser :one
SELECT "event"."id"
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
JOIN "user" ON "user"."id" = "participant"."user_id"
WHERE
    "event"."id" = sqlc.arg(event_id)
        AND
    "user"."id" = sqlc.arg(user_id);

-- name: FindEventForUserAsParticipant :one
SELECT "event"."id"
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
JOIN "user" ON "user"."id" = "participant"."user_id"
WHERE
    "event"."id" = $1
        AND
    "participant"."participates" = TRUE
        AND
    "user"."id" = $2;

-- name: FindEventForUserAsOrganizer :one
SELECT "event"."id"
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
JOIN "user" ON "user"."id" = "participant"."user_id"
WHERE
    "event"."id" = $1
        AND
    "participant"."organizer" = TRUE
        AND
    "user"."id" = $2;
