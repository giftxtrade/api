-- name: CreateLink :one
INSERT INTO "link" (
    code,
    expiration_date,
    event_id
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: FindLinkByCode :one
SELECT * FROM "link" WHERE code = $1;

-- name: FindLinkWithEventByCode :one
SELECT
    sqlc.embed(link),
    sqlc.embed(event)
FROM "link"
JOIN "event" ON "event"."id" = "link"."event_id"
WHERE code = $1;
