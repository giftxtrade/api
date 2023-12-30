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

const deleteEvent = `-- name: DeleteEvent :one
DELETE FROM "event"
WHERE "event"."id" = $1
RETURNING id, name, description, budget, invitation_message, draw_at, close_at, created_at, updated_at
`

func (q *Queries) DeleteEvent(ctx context.Context, id int64) (Event, error) {
	row := q.queryRow(ctx, q.deleteEventStmt, deleteEvent, id)
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
    p.id, p.name, p.email, p.address, p.organizer, p.participates, p.accepted, p.event_id, p.user_id, p.created_at, p.updated_at, p.user_name, p.user_email, p.user_image_url
FROM "event"
JOIN "participant" "p1" ON "p1"."event_id" = "event"."id"
JOIN "participant_user" "p" ON "p"."event_id" = "event"."id"
WHERE 
    "p1"."user_id" = $1
ORDER BY
    "event"."draw_at" ASC,
    "event"."close_at" ASC,
    "p"."id" ASC
`

type FindAllEventsWithUserRow struct {
	Event           Event           `db:"event" json:"event"`
	ParticipantUser ParticipantUser `db:"participant_user" json:"participantUser"`
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
			&i.ParticipantUser.ID,
			&i.ParticipantUser.Name,
			&i.ParticipantUser.Email,
			&i.ParticipantUser.Address,
			&i.ParticipantUser.Organizer,
			&i.ParticipantUser.Participates,
			&i.ParticipantUser.Accepted,
			&i.ParticipantUser.EventID,
			&i.ParticipantUser.UserID,
			&i.ParticipantUser.CreatedAt,
			&i.ParticipantUser.UpdatedAt,
			&i.ParticipantUser.UserName,
			&i.ParticipantUser.UserEmail,
			&i.ParticipantUser.UserImageUrl,
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

const findEventById = `-- name: FindEventById :many
SELECT
    event.id, event.name, event.description, event.budget, event.invitation_message, event.draw_at, event.close_at, event.created_at, event.updated_at,
    p.id, p.name, p.email, p.address, p.organizer, p.participates, p.accepted, p.event_id, p.user_id, p.created_at, p.updated_at, p.user_name, p.user_email, p.user_image_url
FROM "event"
JOIN "participant_user" "p" ON "p"."event_id" = "event"."id"
WHERE "event"."id" = $1
`

type FindEventByIdRow struct {
	Event           Event           `db:"event" json:"event"`
	ParticipantUser ParticipantUser `db:"participant_user" json:"participantUser"`
}

func (q *Queries) FindEventById(ctx context.Context, id int64) ([]FindEventByIdRow, error) {
	rows, err := q.query(ctx, q.findEventByIdStmt, findEventById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindEventByIdRow
	for rows.Next() {
		var i FindEventByIdRow
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
			&i.ParticipantUser.ID,
			&i.ParticipantUser.Name,
			&i.ParticipantUser.Email,
			&i.ParticipantUser.Address,
			&i.ParticipantUser.Organizer,
			&i.ParticipantUser.Participates,
			&i.ParticipantUser.Accepted,
			&i.ParticipantUser.EventID,
			&i.ParticipantUser.UserID,
			&i.ParticipantUser.CreatedAt,
			&i.ParticipantUser.UpdatedAt,
			&i.ParticipantUser.UserName,
			&i.ParticipantUser.UserEmail,
			&i.ParticipantUser.UserImageUrl,
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

const findEventInvites = `-- name: FindEventInvites :many
SELECT event.id, event.name, event.description, event.budget, event.invitation_message, event.draw_at, event.close_at, event.created_at, event.updated_at
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
WHERE 
    "participant"."accepted" = FALSE
        AND
    "participant"."email" = $1
`

func (q *Queries) FindEventInvites(ctx context.Context, email string) ([]Event, error) {
	rows, err := q.query(ctx, q.findEventInvitesStmt, findEventInvites, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Budget,
			&i.InvitationMessage,
			&i.DrawAt,
			&i.CloseAt,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateEvent = `-- name: UpdateEvent :one
UPDATE "event"
SET
    "name" = COALESCE($2, "name"),
    "description" = COALESCE($3, "description"),
    "budget" = COALESCE($4, "budget"),
    "draw_at" = COALESCE($5, "draw_at"),
    "close_at" = COALESCE($6, "close_at"),
    "updated_at" = now()
WHERE "event"."id" = $1
RETURNING id, name, description, budget, invitation_message, draw_at, close_at, created_at, updated_at
`

type UpdateEventParams struct {
	ID          int64          `db:"id" json:"id"`
	Name        sql.NullString `db:"name" json:"name"`
	Description sql.NullString `db:"description" json:"description"`
	Budget      sql.NullString `db:"budget" json:"budget"`
	DrawAt      sql.NullTime   `db:"draw_at" json:"drawAt"`
	CloseAt     sql.NullTime   `db:"close_at" json:"closeAt"`
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) (Event, error) {
	row := q.queryRow(ctx, q.updateEventStmt, updateEvent,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Budget,
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

const verifyEventForUserAsOrganizer = `-- name: VerifyEventForUserAsOrganizer :one
SELECT "event"."id"
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
JOIN "user" ON "user"."id" = "participant"."user_id"
WHERE
    "event"."id" = $1
        AND
    "participant"."organizer" = TRUE
        AND
    "user"."id" = $2
`

type VerifyEventForUserAsOrganizerParams struct {
	EventID int64 `db:"event_id" json:"eventId"`
	UserID  int64 `db:"user_id" json:"userId"`
}

func (q *Queries) VerifyEventForUserAsOrganizer(ctx context.Context, arg VerifyEventForUserAsOrganizerParams) (int64, error) {
	row := q.queryRow(ctx, q.verifyEventForUserAsOrganizerStmt, verifyEventForUserAsOrganizer, arg.EventID, arg.UserID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const verifyEventForUserAsParticipant = `-- name: VerifyEventForUserAsParticipant :one
SELECT "event"."id"
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
JOIN "user" ON "user"."id" = "participant"."user_id"
WHERE
    "event"."id" = $1
        AND
    "participant"."participates" = TRUE
        AND
    "user"."id" = $2
`

type VerifyEventForUserAsParticipantParams struct {
	EventID int64 `db:"event_id" json:"eventId"`
	UserID  int64 `db:"user_id" json:"userId"`
}

func (q *Queries) VerifyEventForUserAsParticipant(ctx context.Context, arg VerifyEventForUserAsParticipantParams) (int64, error) {
	row := q.queryRow(ctx, q.verifyEventForUserAsParticipantStmt, verifyEventForUserAsParticipant, arg.EventID, arg.UserID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const verifyEventWithEmailOrUser = `-- name: VerifyEventWithEmailOrUser :one

SELECT "event"."id"
FROM "event"
JOIN "participant" ON "participant"."event_id" = "event"."id"
WHERE
    "event"."id" = $1
        AND
    (
        "participant"."user_id" = $2 OR "participant"."email" = $3
    )
`

type VerifyEventWithEmailOrUserParams struct {
	EventID int64          `db:"event_id" json:"eventId"`
	UserID  sql.NullInt64  `db:"user_id" json:"userId"`
	Email   sql.NullString `db:"email" json:"email"`
}

// event verification queries
func (q *Queries) VerifyEventWithEmailOrUser(ctx context.Context, arg VerifyEventWithEmailOrUserParams) (int64, error) {
	row := q.queryRow(ctx, q.verifyEventWithEmailOrUserStmt, verifyEventWithEmailOrUser, arg.EventID, arg.UserID, arg.Email)
	var id int64
	err := row.Scan(&id)
	return id, err
}
