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

-- name: AcceptEventInvite :one
UPDATE "participant"
SET
    "accepted" = TRUE,
    "user_id" = $1,
    "updated_at" = now()
WHERE 
    "email" = $2
        AND
    "event_id" = $3
RETURNING *;

-- name: DeclineEventInvite :one
DELETE FROM "participant"
WHERE "email" = $1 AND "event_id" = $2
RETURNING *;

-- name: FindParticipantFromEventIdAndUser :one
SELECT * FROM "participant"
WHERE "event_id" = $1 AND "user_id" = $2;

-- name: FindParticipantWithIdAndEventId :one
SELECT * FROM "participant"
WHERE "event_id" = $1 AND "id" = sqlc.arg(participant_id);

-- name: UpdateParticipantStatus :one
UPDATE "participant"
SET
    "organizer" = COALESCE(sqlc.narg(organizer), "organizer"),
    "participates" = COALESCE(sqlc.narg(participates), "participates"),
    "updated_at" = now()
WHERE "event_id" = $1 AND "id" = sqlc.arg(participant_id)
RETURNING *;
