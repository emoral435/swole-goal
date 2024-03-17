package routes

import (
	"net/http"

	mw "github.com/emoral435/swole-goal/api/middleware"
	"github.com/emoral435/swole-goal/api/token"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

func ServeExercises(mux *http.ServeMux, ss *ServerStore) {
	eAPI := createExerciseAPIStruct(ss.TokenMaker, ss.Store, ss.Config) // implements IWorkout interface
	// create a workout
	mux.Handle("POST /swole/workout", mw.EnforceJSONHandler(mw.AuthMiddleware(eAPI.TokenMaker(), http.HandlerFunc(eAPI.CreateExercise))))
	// get all users workouts
	mux.Handle("GET /swole/workout", mw.EnforceJSONHandler(mw.AuthMiddleware(eAPI.TokenMaker(), http.HandlerFunc(eAPI.GetAllExercises))))
	// modify a users workout
	mux.Handle("PUT /swole/workout", mw.EnforceJSONHandler(mw.AuthMiddleware(eAPI.TokenMaker(), http.HandlerFunc(eAPI.ModifyExercise))))
	// delete a users workout given the workout ID
	mux.Handle("DELETE /swole/workout/{wid}", mw.EnforceJSONHandler(mw.AuthMiddleware(eAPI.TokenMaker(), http.HandlerFunc(eAPI.DeleteOneExercise))))
	// delete all user's workouts
	mux.Handle("DELETE /swole/workout", mw.EnforceJSONHandler(mw.AuthMiddleware(eAPI.TokenMaker(), http.HandlerFunc(eAPI.DeleteAllExercises))))
}

// CreateExercise creates a new exercise for the associated workout and authenticated user.
//
// no information is needed, but we should be able to return the exercise struct containing
// the exercise information in case we want to be able to display the new exercise information
func (api *ExerciseAPI) CreateExercise(res http.ResponseWriter, req *http.Request) {

}

// GetAllExercises returns a list of all the workouts exercises in the server store
func (api *ExerciseAPI) GetAllExercises(res http.ResponseWriter, req *http.Request) {

}

// ModifyExercise updates a workouts exercise
func (api *ExerciseAPI) ModifyExercise(res http.ResponseWriter, req *http.Request) {

}

// DeleteOneExercise deletes one workouts exercise
func (api *ExerciseAPI) DeleteOneExercise(res http.ResponseWriter, req *http.Request) {

}

// DeleteAllExercises deletes one users exercise
func (api *ExerciseAPI) DeleteAllExercises(res http.ResponseWriter, req *http.Request) {

}

// WorkoutAPI struct contains the API information and methods for making, updating, deleting, and all other API information related to user workouts.
type ExerciseAPI struct {
	tokenMaker token.Maker
	store      *db.Store
	config     util.Config
}

// createWorkoutAPI creates a new UserAPI struct instance
func createExerciseAPIStruct(t token.Maker, store *db.Store, config util.Config) IExercise {
	return &ExerciseAPI{
		t,
		store,
		config,
	}
}

func (e *ExerciseAPI) TokenMaker() token.Maker {
	return e.tokenMaker
}

func (e *ExerciseAPI) Store() *db.Store {
	return e.store
}

func (e *ExerciseAPI) Config() util.Config {
	return e.config

}
