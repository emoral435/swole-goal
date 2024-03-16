package routes

import (
	"net/http"
	"strconv"
	"time"

	mw "github.com/emoral435/swole-goal/api/middleware"
	"github.com/emoral435/swole-goal/api/token"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

func ServeWorkouts(mux *http.ServeMux, ss *ServerStore) {
	wAPI := createWorkoutAPIStruct(ss.TokenMaker, ss.Store, ss.Config) // implements IWorkout interface
	// create a workout
	mux.Handle("POST /swole/workout", mw.EnforceJSONHandler(mw.AuthMiddleware(wAPI.TokenMaker(), http.HandlerFunc(wAPI.CreateWorkout))))
	// get all users workouts
	mux.Handle("GET /swole/workout", mw.EnforceJSONHandler(mw.AuthMiddleware(wAPI.TokenMaker(), http.HandlerFunc(wAPI.GetAllWorkouts))))
	// get a specific workout
	mux.Handle("GET /swole/workout/{id}", mw.EnforceJSONHandler(mw.AuthMiddleware(wAPI.TokenMaker(), http.HandlerFunc(wAPI.GetOneWorkout))))
	// modify a users workout
	mux.Handle("PUT /swole/workout/{id}", mw.EnforceJSONHandler(mw.AuthMiddleware(wAPI.TokenMaker(), http.HandlerFunc(wAPI.ModifyWorkout))))
	// delete a users workout given the workout ID
	mux.Handle("DELETE /swole/workout/{id}", mw.EnforceJSONHandler(mw.AuthMiddleware(wAPI.TokenMaker(), http.HandlerFunc(wAPI.DeleteOneWorkout))))
	// delete all user's workouts
	mux.Handle("DELETE /swole/workout", mw.EnforceJSONHandler(mw.AuthMiddleware(wAPI.TokenMaker(), http.HandlerFunc(wAPI.DeleteAllWorkouts))))
}

func (api *WorkoutAPI) CreateWorkout(res http.ResponseWriter, req *http.Request) {
	queries, header := req.URL.Query(), req.Header
	// for create workout params struct
	uid, title, body := header.Get("uid"), queries.Get("title"), queries.Get("body")
	// invalid user uid
	if len(uid) <= 0 {
		util.ReturnErrorJSONResponse(res, "Invalid user id for the corresponding workout", 400)
		return
	}
	// check if we recieved the neccessary params for workout. Otherwise, we can always leave blank
	wParams, err := MakeCreateWorkoutParams(uid, title, body)
	numWorkouts, errNum := api.Store().GetNumWorkouts(req.Context(), wParams.ID)

	// cap the number of workouts one user can have to be AT MOST 7
	if numWorkouts > 7 || errNum != nil {
		util.ReturnErrorJSONResponse(res, "Error occured: user can only have 7 exercises max", 400)
		return
	}

	// making the params had trouble
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	// this actually makers the workout within the database
	newWorkout, err := api.Store().Queries.CreateWorkout(req.Context(), *wParams)

	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	util.ReturnValidJSONResponse(res, newWorkout)
}

// MakeCreateWorkoutParams function returns the CreateWorkoutParams struct object
func MakeCreateWorkoutParams(uid string, title string, body string) (*db.CreateWorkoutParams, error) {
	parsedUID, err := strconv.ParseInt(uid, 10, 64)

	if err != nil {
		return nil, err
	}

	return &db.CreateWorkoutParams{
		ID:       parsedUID,
		Title:    title,
		Body:     body,
		LastTime: time.Now(),
	}, nil
}

func (api *WorkoutAPI) GetAllWorkouts(res http.ResponseWriter, req *http.Request) {
	uid, err := strconv.ParseInt(req.Header.Get("uid"), 10, 64)

	// invalid user uid
	if err != nil {
		util.ReturnErrorJSONResponse(res, "Invalid user id for fetching all workouts", 400)
		return
	}

	allWorkouts, err := api.Store().Queries.GetUserWorkouts(req.Context(), uid)

	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	util.ReturnValidJSONResponse(res, allWorkouts)
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
