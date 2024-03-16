package routes

import (
	"net/http"

	mw "github.com/emoral435/swole-goal/api/middleware"
	"github.com/emoral435/swole-goal/api/token"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

func ServeWorkouts(mux *http.ServeMux, ss *ServerStore) {
	wAPI := createWorkoutAPIStruct(ss.TokenMaker, ss.Store, ss.Config) // implements IWorkout interface
	// create a workout
	mux.Handle("POST /user", mw.EnforceJSONHandler(http.HandlerFunc(wAPI.CreateWorkout)))
	// TODO get all users workouts

	// TODO get a specific workout

	// TODO modify a users workout

	// TODO delete a users workout

	// TODO delete all user's workouts
}

func (api *WorkoutAPI) CreateWorkout(res http.ResponseWriter, req *http.Request) {
	// TODO
}

func (api *WorkoutAPI) GetAllWorkouts(res http.ResponseWriter, req *http.Request) {
	// TODO
}

func (api *WorkoutAPI) GetOneWorkout(res http.ResponseWriter, req *http.Request) {
	// TODO
}

func (api *WorkoutAPI) ModifyWorkout(res http.ResponseWriter, req *http.Request) {
	// TODO
}

func (api *WorkoutAPI) DeleteOneWorkout(res http.ResponseWriter, req *http.Request) {
	// TODO
}

func (api *WorkoutAPI) DeleteAllWorkouts(res http.ResponseWriter, req *http.Request) {
	// TODO
}

// WorkoutAPI struct contains the API information and methods for making, updating, deleting, and all other API information related to user workouts.
type WorkoutAPI struct {
	tokenMaker token.Maker
	store      *db.Store
	config     util.Config
}

// createWorkoutAPI creates a new UserAPI struct instance
func createWorkoutAPIStruct(t token.Maker, store *db.Store, config util.Config) IWorkout {
	return &WorkoutAPI{
		t,
		store,
		config,
	}
}

func (w *WorkoutAPI) TokenMaker() token.Maker {
	return w.tokenMaker
}

func (w *WorkoutAPI) Store() *db.Store {
	return w.store
}

func (w *WorkoutAPI) Config() util.Config {
	return w.config

}
