package routes

import (
	"net/http"

	"github.com/emoral435/swole-goal/api/token"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

type IWorkout interface {
	TokenMaker() token.Maker
	Store() *db.Store
	Config() util.Config

	// CreateWorkout creates a new workout for the associated and authenticated user.
	//
	// no information is needed, but we should be able to return the workout struct containing
	// the workout information in case we want to be able to make a form popup on the creation of the workout.
	CreateWorkout(res http.ResponseWriter, req *http.Request)
	// GetAllWorkouts returns a list of all workout's in the server store
	GetAllWorkouts(res http.ResponseWriter, req *http.Request)
	// GetOneWorkout returns a workout based on workout id within the database
	GetOneWorkout(res http.ResponseWriter, req *http.Request)
	// ModifyWorkout updates a users workout
	ModifyWorkout(res http.ResponseWriter, req *http.Request)
	// DeleteOneWorkout deletes one users workout
	DeleteOneWorkout(res http.ResponseWriter, req *http.Request)
	// DeleteAllWorkouts deletes one users workout
	DeleteAllWorkouts(res http.ResponseWriter, req *http.Request)
}
