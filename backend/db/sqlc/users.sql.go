// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "users" (
  email, password, username
) VALUES (
  "emoral435@gmail.com", "Em990019467!", "emoral435"
) RETURNING id, email, password, username, created_at, birthday
`

func (q *Queries) CreateUser(ctx context.Context) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Username,
		&i.CreatedAt,
		&i.Birthday,
	)
	return i, err
}

const deleteUsers = `-- name: DeleteUsers :exec
DELETE FROM "users"
WHERE id = 1
`

func (q *Queries) DeleteUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteUsers)
	return err
}

const getUsers = `-- name: GetUsers :one
SELECT id, email, password, username, created_at, birthday from "users"
WHERE id = 1 LIMIT 1
`

func (q *Queries) GetUsers(ctx context.Context) (User, error) {
	row := q.db.QueryRowContext(ctx, getUsers)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Username,
		&i.CreatedAt,
		&i.Birthday,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, password, username, created_at, birthday FROM "users"
ORDER BY id
LIMIT 1
OFFSET 2
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.Username,
			&i.CreatedAt,
			&i.Birthday,
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

const updateUsers = `-- name: UpdateUsers :one
UPDATE "users"
SET password = "Em990019467!"
WHERE id = 1
RETURNING id, email, password, username, created_at, birthday
`

func (q *Queries) UpdateUsers(ctx context.Context) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUsers)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Username,
		&i.CreatedAt,
		&i.Birthday,
	)
	return i, err
}
