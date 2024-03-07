// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: workouts.sql

package db

import (
	"context"
	"time"
)

const createWorkout = `-- name: CreateWorkout :one
INSERT INTO "workouts" (
    user_id, title, body, last_time
) VALUES (
    (SELECT "users".id FROM "users" WHERE "users".id = $1), $2, $3, $4
) RETURNING id, user_id, title, body, last_time
`

type CreateWorkoutParams struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	LastTime time.Time `json:"last_time"`
}

// CreateWorkout: returns a new workout, provided uid, title, body, and last time modified/used
//
// returns: the new workout row
func (q *Queries) CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (Workout, error) {
	row := q.db.QueryRowContext(ctx, createWorkout,
		arg.ID,
		arg.Title,
		arg.Body,
		arg.LastTime,
	)
	var i Workout
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.LastTime,
	)
	return i, err
}

const deleteAllWorkouts = `-- name: DeleteAllWorkouts :exec
DELETE FROM "workouts"
WHERE user_id = $1
`

// DeleteAllWorkouts: deletes All user's workouts
//
// returns: nothing! see https://docs.sqlc.dev/en/stable/reference/query-annotations.html for exec
func (q *Queries) DeleteAllWorkouts(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteAllWorkouts, userID)
	return err
}

const deleteSingleWorkout = `-- name: DeleteSingleWorkout :exec
DELETE FROM "workouts"
WHERE id = $1
`

// DeleteSingleWorkout: deletes a single user's workout
//
// returns: nothing! see https://docs.sqlc.dev/en/stable/reference/query-annotations.html for exec
func (q *Queries) DeleteSingleWorkout(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSingleWorkout, id)
	return err
}

const getUserWorkouts = `-- name: GetUserWorkouts :many
SELECT id, user_id, title, body, last_time FROM "workouts"
WHERE user_id = $1
`

// GetUserWorkouts: returns a users workouts, provided their uid
//
// returns: the user's corresponding workout rows
func (q *Queries) GetUserWorkouts(ctx context.Context, userID int64) ([]Workout, error) {
	rows, err := q.db.QueryContext(ctx, getUserWorkouts, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Workout
	for rows.Next() {
		var i Workout
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Body,
			&i.LastTime,
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

const getWorkout = `-- name: GetWorkout :one
SELECT id, user_id, title, body, last_time FROM "workouts"
WHERE id = $1 LIMIT 1
`

// GetWorkout: returns an existing workout, given workout id
//
// returns: the corresponding workout row
func (q *Queries) GetWorkout(ctx context.Context, id int64) (Workout, error) {
	row := q.db.QueryRowContext(ctx, getWorkout, id)
	var i Workout
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.LastTime,
	)
	return i, err
}

const updateWorkoutBody = `-- name: UpdateWorkoutBody :one
UPDATE "workouts"
SET body = $2
WHERE id = $1
RETURNING id, user_id, title, body, last_time
`

type UpdateWorkoutBodyParams struct {
	ID   int64  `json:"id"`
	Body string `json:"body"`
}

// UpdateBody: updates workout's body text given its workouts id
//
// returns: the workouts new row
func (q *Queries) UpdateWorkoutBody(ctx context.Context, arg UpdateWorkoutBodyParams) (Workout, error) {
	row := q.db.QueryRowContext(ctx, updateWorkoutBody, arg.ID, arg.Body)
	var i Workout
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.LastTime,
	)
	return i, err
}

const updateWorkoutLast = `-- name: UpdateWorkoutLast :one
UPDATE "workouts"
SET last_time = $2
WHERE id = $1
RETURNING id, user_id, title, body, last_time
`

type UpdateWorkoutLastParams struct {
	ID       int64     `json:"id"`
	LastTime time.Time `json:"last_time"`
}

// UpdateLastWorkout: updates workout's last workout time given its id
//
// returns: the workout's new corresponding row
func (q *Queries) UpdateWorkoutLast(ctx context.Context, arg UpdateWorkoutLastParams) (Workout, error) {
	row := q.db.QueryRowContext(ctx, updateWorkoutLast, arg.ID, arg.LastTime)
	var i Workout
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.LastTime,
	)
	return i, err
}

const updateWorkoutTitle = `-- name: UpdateWorkoutTitle :one
UPDATE "workouts"
SET title = $2
WHERE id = $1
RETURNING id, user_id, title, body, last_time
`

type UpdateWorkoutTitleParams struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// UpdateWorkoutTitle: updates workouts title given its id
//
// returns: the workout's new corresponding row
func (q *Queries) UpdateWorkoutTitle(ctx context.Context, arg UpdateWorkoutTitleParams) (Workout, error) {
	row := q.db.QueryRowContext(ctx, updateWorkoutTitle, arg.ID, arg.Title)
	var i Workout
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.LastTime,
	)
	return i, err
}