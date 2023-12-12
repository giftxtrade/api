// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user.sql

package database

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (
    name,
    email,
    image_url,
    phone,
    admin,
    active
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING id, name, email, image_url, phone, admin, active, created_at, updated_at
`

type CreateUserParams struct {
	Name     string         `db:"name" json:"name"`
	Email    string         `db:"email" json:"email"`
	ImageUrl string         `db:"image_url" json:"imageUrl"`
	Phone    sql.NullString `db:"phone" json:"phone"`
	Admin    bool           `db:"admin" json:"admin"`
	Active   bool           `db:"active" json:"active"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.Name,
		arg.Email,
		arg.ImageUrl,
		arg.Phone,
		arg.Admin,
		arg.Active,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.ImageUrl,
		&i.Phone,
		&i.Admin,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, name, email, image_url, phone, admin, active, created_at, updated_at FROM "user"
WHERE email = $1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.findUserByEmailStmt, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.ImageUrl,
		&i.Phone,
		&i.Admin,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserById = `-- name: FindUserById :one
SELECT id, name, email, image_url, phone, admin, active, created_at, updated_at FROM "user"
WHERE id = $1
`

func (q *Queries) FindUserById(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.findUserByIdStmt, findUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.ImageUrl,
		&i.Phone,
		&i.Admin,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByIdAndEmail = `-- name: FindUserByIdAndEmail :one
SELECT id, name, email, image_url, phone, admin, active, created_at, updated_at FROM "user"
WHERE id = $1 AND email = $2
`

type FindUserByIdAndEmailParams struct {
	ID    int64  `db:"id" json:"id"`
	Email string `db:"email" json:"email"`
}

func (q *Queries) FindUserByIdAndEmail(ctx context.Context, arg FindUserByIdAndEmailParams) (User, error) {
	row := q.queryRow(ctx, q.findUserByIdAndEmailStmt, findUserByIdAndEmail, arg.ID, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.ImageUrl,
		&i.Phone,
		&i.Admin,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByIdOrEmail = `-- name: FindUserByIdOrEmail :one
SELECT id, name, email, image_url, phone, admin, active, created_at, updated_at FROM "user"
WHERE id = $1 OR email = $2
`

type FindUserByIdOrEmailParams struct {
	ID    int64  `db:"id" json:"id"`
	Email string `db:"email" json:"email"`
}

func (q *Queries) FindUserByIdOrEmail(ctx context.Context, arg FindUserByIdOrEmailParams) (User, error) {
	row := q.queryRow(ctx, q.findUserByIdOrEmailStmt, findUserByIdOrEmail, arg.ID, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.ImageUrl,
		&i.Phone,
		&i.Admin,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
