// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: link.sql

package database

import (
	"context"
	"time"
)

const createLink = `-- name: CreateLink :one
INSERT INTO "link" (
    code,
    expiration_date,
    event_id
) VALUES (
    $1, $2, $3
) RETURNING id, code, expiration_date, event_id, created_at, updated_at
`

type CreateLinkParams struct {
	Code           string    `db:"code" json:"code"`
	ExpirationDate time.Time `db:"expiration_date" json:"expirationDate"`
	EventID        int64     `db:"event_id" json:"eventId"`
}

func (q *Queries) CreateLink(ctx context.Context, arg CreateLinkParams) (Link, error) {
	row := q.queryRow(ctx, q.createLinkStmt, createLink, arg.Code, arg.ExpirationDate, arg.EventID)
	var i Link
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.ExpirationDate,
		&i.EventID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findLinkByCode = `-- name: FindLinkByCode :one
SELECT id, code, expiration_date, event_id, created_at, updated_at FROM "link" WHERE code = $1
`

func (q *Queries) FindLinkByCode(ctx context.Context, code string) (Link, error) {
	row := q.queryRow(ctx, q.findLinkByCodeStmt, findLinkByCode, code)
	var i Link
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.ExpirationDate,
		&i.EventID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findLinkWithEventByCode = `-- name: FindLinkWithEventByCode :one
SELECT
    link.id, link.code, link.expiration_date, link.event_id, link.created_at, link.updated_at,
    event.id, event.name, event.description, event.budget, event.invitation_message, event.draw_at, event.close_at, event.created_at, event.updated_at
FROM "link"
JOIN "event" ON "event"."id" = "link"."event_id"
WHERE code = $1
`

type FindLinkWithEventByCodeRow struct {
	Link  Link  `db:"link" json:"link"`
	Event Event `db:"event" json:"event"`
}

func (q *Queries) FindLinkWithEventByCode(ctx context.Context, code string) (FindLinkWithEventByCodeRow, error) {
	row := q.queryRow(ctx, q.findLinkWithEventByCodeStmt, findLinkWithEventByCode, code)
	var i FindLinkWithEventByCodeRow
	err := row.Scan(
		&i.Link.ID,
		&i.Link.Code,
		&i.Link.ExpirationDate,
		&i.Link.EventID,
		&i.Link.CreatedAt,
		&i.Link.UpdatedAt,
		&i.Event.ID,
		&i.Event.Name,
		&i.Event.Description,
		&i.Event.Budget,
		&i.Event.InvitationMessage,
		&i.Event.DrawAt,
		&i.Event.CloseAt,
		&i.Event.CreatedAt,
		&i.Event.UpdatedAt,
	)
	return i, err
}
