package routes

import (
	"net/http"

	"github.com/emoral435/swole-goal/api/token"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

type IExercise interface {
	TokenMaker() token.Maker
	Store() *db.Store
	Config() util.Config

	// CreateExercise creates a new exercise for the associated workout and authenticated user.
	//
	// no information is needed, but we should be able to return the exercise struct containing
	// the exercise information in case we want to be able to display the new exercise information
	CreateExercise(res http.ResponseWriter, req *http.Request)
	// GetAllExercises returns a list of all the workouts exercises in the server store
	GetAllExercises(res http.ResponseWriter, req *http.Request)
	// ModifyExercise updates a workouts exercise
	ModifyExercise(res http.ResponseWriter, req *http.Request)
	// DeleteOneExercise deletes one workouts exercise
	DeleteOneExercise(res http.ResponseWriter, req *http.Request)
	// DeleteAllExercises deletes one users exercise
	DeleteAllExercises(res http.ResponseWriter, req *http.Request)
}
