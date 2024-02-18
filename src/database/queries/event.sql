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

-- name: UpdateEvent :one
UPDATE "event"
SET
    "name" = COALESCE(sqlc.narg(name), "name"),
    "description" = COALESCE(sqlc.narg(description), "description"),
    "budget" = COALESCE(sqlc.narg(budget), "budget"),
    "draw_at" = COALESCE(sqlc.narg(draw_at), "draw_at"),
    "close_at" = COALESCE(sqlc.narg(close_at), "close_at"),
    "updated_at" = now()
WHERE "event"."id" = $1
RETURNING *;

-- name: DeleteEvent :one
DELETE FROM "event"
WHERE "event"."id" = $1
RETURNING *;

-- name: FindAllEventsWithUser :many
SELECT
    sqlc.embed(event),
    sqlc.embed(p)
FROM "event"
JOIN "participant" "p1" ON "p1"."event_id" = "event"."id"
JOIN "participant_user" "p" ON "p"."event_id" = "event"."id"
WHERE 
    "p1"."user_id" = $1
ORDER BY
    "event"."draw_at" ASC,
    "event"."close_at" ASC,
    "p"."id" ASC;

-- name: FindEventSimple :one
SELECT * FROM "event" WHERE "event"."id" = $1;

-- name: FindEventById :many
SELECT
    sqlc.embed(event_link),
    sqlc.embed(p)
FROM "event_link"
JOIN "participant_user" "p" ON "p"."event_id" = "event_link"."id"
WHERE "event_link"."id" = $1
ORDER BY 
    "p"."organizer" DESC,
	"p"."accepted" DESC,
    "p"."created_at" DESC;

-- name: FindEventInvites :many
SELECT "event".*
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
WHERE 
    "participant"."accepted" = FALSE
        AND
    "participant"."email" = $1;


--
-- event verification queries
--

-- name: VerifyEventWithEmailOrUser :one
SELECT "event"."id"
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
WHERE
    "event"."id" = sqlc.arg(event_id)
        AND
    (
        "participant"."user_id" = sqlc.narg(user_id) OR "participant"."email" = sqlc.narg(email)
    );

-- name: VerifyEventWithParticipantId :one
SELECT "event"."id"
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
WHERE
    "event"."id" = sqlc.arg(event_id)
        AND
    "participant"."id" = sqlc.arg(participant_id);

-- name: VerifyEventForUserAsParticipant :one
SELECT "event"."id"
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
JOIN "user" ON "user"."id" = "participant"."user_id"
WHERE
    "event"."id" = sqlc.arg(event_id)
        AND
    "participant"."participates" = TRUE
        AND
    "user"."id" = sqlc.arg(user_id);

-- name: VerifyEventForUserAsOrganizer :one
SELECT "event"."id"
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
JOIN "user" ON "user"."id" = "participant"."user_id"
WHERE
    "event"."id" = sqlc.arg(event_id)
        AND
    "participant"."organizer" = TRUE
        AND
    "user"."id" = sqlc.arg(user_id);
