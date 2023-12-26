// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: event.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createEvent = `-- name: CreateEvent :one
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
RETURNING id, name, description, budget, invitation_message, draw_at, close_at, created_at, updated_at
`

type CreateEventParams struct {
	Name              string         `db:"name" json:"name"`
	Description       sql.NullString `db:"description" json:"description"`
	Budget            string         `db:"budget" json:"budget"`
	InvitationMessage string         `db:"invitation_message" json:"invitationMessage"`
	DrawAt            time.Time      `db:"draw_at" json:"drawAt"`
	CloseAt           time.Time      `db:"close_at" json:"closeAt"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.queryRow(ctx, q.createEventStmt, createEvent,
		arg.Name,
		arg.Description,
		arg.Budget,
		arg.InvitationMessage,
		arg.DrawAt,
		arg.CloseAt,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Budget,
		&i.InvitationMessage,
		&i.DrawAt,
		&i.CloseAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findAllEventsWithUser = `-- name: FindAllEventsWithUser :many
SELECT
    event.id, event.name, event.description, event.budget, event.invitation_message, event.draw_at, event.close_at, event.created_at, event.updated_at,
    p2.id, p2.name, p2.email, p2.address, p2.organizer, p2.participates, p2.accepted, p2.event_id, p2.user_id, p2.created_at, p2.updated_at,
    u.id, u.name, u.email, u.image_url, u.phone, u.admin, u.active, u.created_at, u.updated_at
FROM "event"
JOIN "participant" "p1" ON "p1"."event_id" = "event"."id"
JOIN "participant" "p2" ON "p2"."event_id" = "event"."id"
LEFT JOIN "user" "u" ON "u"."id" = "p2"."user_id"
WHERE 
    "p1"."user_id" = $1
ORDER BY
    "event"."draw_at" DESC,
    "event"."close_at" DESC
`

type FindAllEventsWithUserRow struct {
	Event       Event       `db:"event" json:"event"`
	Participant Participant `db:"participant" json:"participant"`
	User        User        `db:"user" json:"user"`
}

func (q *Queries) FindAllEventsWithUser(ctx context.Context, userID sql.NullInt64) ([]FindAllEventsWithUserRow, error) {
	rows, err := q.query(ctx, q.findAllEventsWithUserStmt, findAllEventsWithUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAllEventsWithUserRow
	for rows.Next() {
		var i FindAllEventsWithUserRow
		if err := rows.Scan(
			&i.Event.ID,
			&i.Event.Name,
			&i.Event.Description,
			&i.Event.Budget,
			&i.Event.InvitationMessage,
			&i.Event.DrawAt,
			&i.Event.CloseAt,
			&i.Event.CreatedAt,
			&i.Event.UpdatedAt,
			&i.Participant.ID,
			&i.Participant.Name,
			&i.Participant.Email,
			&i.Participant.Address,
			&i.Participant.Organizer,
			&i.Participant.Participates,
			&i.Participant.Accepted,
			&i.Participant.EventID,
			&i.Participant.UserID,
			&i.Participant.CreatedAt,
			&i.Participant.UpdatedAt,
			&i.User.ID,
			&i.User.Name,
			&i.User.Email,
			&i.User.ImageUrl,
			&i.User.Phone,
			&i.User.Admin,
			&i.User.Active,
			&i.User.CreatedAt,
			&i.User.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
