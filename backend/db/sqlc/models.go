// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"
)

type Exercise struct {
	ID        int64 `json:"id"`
	WorkoutID int64 `json:"workout_id"`
	// The body section this exercise hits - chest, back, etc.
	Type string `json:"type"`
	// What is the exercise called?
	Title string `json:"title"`
	// description of the exercies - good for reminders
	Desc    sql.NullString `json:"desc"`
	Set1    sql.NullInt64  `json:"set1"`
	Weight1 sql.NullInt64  `json:"weight1"`
	Set2    sql.NullInt64  `json:"set2"`
	Weight2 sql.NullInt64  `json:"weight2"`
	Set3    sql.NullInt64  `json:"set3"`
	Weight3 sql.NullInt64  `json:"weight3"`
	Set4    sql.NullInt64  `json:"set4"`
	Weight4 sql.NullInt64  `json:"weight4"`
	// tracks what the overall volume was the last time this exercise was performed
	LastVolume int64 `json:"last_volume"`
}

type User struct {
	ID int64 `json:"id"`
	// email to sign in - also to send reminders
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Username  string       `json:"username"`
	CreatedAt time.Time    `json:"created_at"`
	Birthday  sql.NullTime `json:"birthday"`
}

type Workout struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	// Description of workout
	Body string `json:"body"`
	// Timestamp of the last time completed
	Last time.Time     `json:"last"`
	Exe1 sql.NullInt64 `json:"exe1"`
	Exe2 sql.NullInt64 `json:"exe2"`
	Exe3 sql.NullInt64 `json:"exe3"`
	Exe4 sql.NullInt64 `json:"exe4"`
	Exe5 sql.NullInt64 `json:"exe5"`
	Exe6 sql.NullInt64 `json:"exe6"`
	Exe7 sql.NullInt64 `json:"exe7"`
}
