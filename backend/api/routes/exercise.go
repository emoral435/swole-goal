package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	mw "github.com/emoral435/swole-goal/api/middleware"
	"github.com/emoral435/swole-goal/api/token"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

func ServeExercises(mux *http.ServeMux, ss *ServerStore) {
	eAPI := createExerciseAPIStruct(ss.TokenMaker, ss.Store, ss.Config) // implements IWorkout interface
	// create a workout
	mux.Handle("POST /swole/exercise", mw.EnforceJSONHandler(mw.AuthMiddleware(eAPI.TokenMaker(), http.HandlerFunc(eAPI.CreateExercise))))
	// get all users workouts
	mux.Handle("GET /swole/exercise", mw.EnforceJSONHandler(mw.AuthMiddleware(eAPI.TokenMaker(), http.HandlerFunc(eAPI.GetAllExercises))))
	// modify a users workout
	mux.Handle("PUT /swole/exercise", mw.EnforceJSONHandler(mw.AuthMiddleware(eAPI.TokenMaker(), http.HandlerFunc(eAPI.ModifyExercise))))
	// delete a users workout given the workout ID
	mux.Handle("DELETE /swole/exercise/{wid}", mw.EnforceJSONHandler(mw.AuthMiddleware(eAPI.TokenMaker(), http.HandlerFunc(eAPI.DeleteOneExercise))))
	// delete all user's workouts
	mux.Handle("DELETE /swole/exercise", mw.EnforceJSONHandler(mw.AuthMiddleware(eAPI.TokenMaker(), http.HandlerFunc(eAPI.DeleteAllExercises))))
}

// CreateExercise creates a new exercise for the associated workout and authenticated user.
//
// no information is needed, but we should be able to return the exercise struct containing
// the exercise information in case we want to be able to display the new exercise information
func (api *ExerciseAPI) CreateExercise(res http.ResponseWriter, req *http.Request) {
	queries, header := req.URL.Query(), req.Header
	// for create workout params struct
	eid, exType, title, desc := header.Get("eid"), queries.Get("type"), queries.Get("title"), queries.Get("description")
	// invalid user wid
	if len(eid) <= 0 {
		util.ReturnErrorJSONResponse(res, "Invalid exercise id for the corresponding workout", 400)
		return
	}
	// check if we recieved the neccessary params for workout. Otherwise, we can always leave blank
	eParams, err := MakeCreateExerciseParams(eid, exType, title, desc)

	// making the params had trouble
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	// this actually makers the workout within the database
	newExercise, err := api.Store().Queries.CreateExercise(req.Context(), *eParams)

	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	util.ReturnValidJSONResponse(res, newExercise)
}

// MakeDesc returns a sql.NullString representation of the description. If the string is empty, we return "No description."
func MakeDesc(description string) *sql.NullString {
	if len(description) > 0 {
		return &sql.NullString{
			String: description,
			Valid:  true,
		}
	}
	return &sql.NullString{
		String: "No description given.",
		Valid:  false,
	}
}

// MakeCreateExerciseParams returns a pointer to a struct for CreateExerciseParams
func MakeCreateExerciseParams(wid string, exType string, title string, desc string) (*db.CreateExerciseParams, error) {
	parsedWID, err := strconv.ParseInt(wid, 10, 64)

	if err != nil {
		return nil, err
	}

	return &db.CreateExerciseParams{
		ID:          parsedWID,
		Type:        exType,
		Title:       title,
		Description: *MakeDesc(desc),
	}, nil
}

// GetAllExercises returns a list of all the workouts exercises in the server store
func (api *ExerciseAPI) GetAllExercises(res http.ResponseWriter, req *http.Request) {
	wid, err := strconv.ParseInt(req.Header.Get("wid"), 10, 64) // get a users id from the header

	// invalid user uid
	if err != nil {
		util.ReturnErrorJSONResponse(res, "Invalid workout id for fetching all exercises", 400)
		return
	}

	allExercises, err := api.Store().Queries.GetWorkoutsExercise(req.Context(), wid) // this actually gets all the workouts

	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	util.ReturnValidJSONResponse(res, allExercises)
}

// ModifyExercise updates a workouts exercise
func (api *ExerciseAPI) ModifyExercise(res http.ResponseWriter, req *http.Request) {
	queries := req.URL.Query()
	eid, errE := strconv.ParseInt(queries.Get("eid"), 10, 64)
	// invalid user uid
	if errE != nil {
		util.ReturnErrorJSONResponse(res, "Invalid exercise/workout id for fetching single exercise", 400)
		return
	}
	exType, title, desc, LastVolume := queries.Get("type"), queries.Get("title"), queries.Get("description"), queries.Get("lastVolume")

	if len(title) > 0 {
		_, errTitle := modifyTitleExercise(req, title, api.Store(), eid)
		if errTitle = util.CheckError(errTitle, res, req); errTitle != nil {
			return
		}
	}

	if len(exType) > 0 {
		_, errType := modifyType(req, exType, api.Store(), eid)
		if errType = util.CheckError(errType, res, req); errType != nil {
			return
		}
	}

	if len(desc) > 0 {
		_, errLastTime := modifyDesc(req, desc, api.Store(), eid)
		if errLastTime = util.CheckError(errLastTime, res, req); errLastTime != nil {
			return
		}
	}

	if len(LastVolume) > 0 {
		parsedVolume, err := strconv.ParseInt(LastVolume, 10, 64)
		if err = util.CheckError(err, res, req); err != nil {
			parsedVolume = int64(0)
		}
		_, errLastTime := modifyExerciseLast(req, parsedVolume, api.Store(), eid)
		if errLastTime = util.CheckError(errLastTime, res, req); errLastTime != nil {
			return
		}
	}

	newExercise, err := api.Store().GetExercise(req.Context(), eid)

	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	util.ReturnValidJSONResponse(res, newExercise)
}

func modifyTitleExercise(req *http.Request, title string, store *db.Store, eid int64) (db.Exercise, error) {
	return store.Queries.UpdateExerciseTitle(req.Context(), db.UpdateExerciseTitleParams{ID: eid, Title: title})
}

func modifyType(req *http.Request, exType string, store *db.Store, eid int64) (db.Exercise, error) {
	return store.Queries.UpdateExerciseType(req.Context(), db.UpdateExerciseTypeParams{ID: eid, Type: exType})
}

func modifyDesc(req *http.Request, desc string, store *db.Store, eid int64) (db.Exercise, error) {
	newDesc := MakeDesc(desc)
	return store.Queries.UpdateExerciseDescription(req.Context(), db.UpdateExerciseDescriptionParams{ID: eid, Description: *newDesc})
}

func modifyExerciseLast(req *http.Request, lastVolume int64, store *db.Store, eid int64) (db.Exercise, error) {
	return store.Queries.UpdateExerciseLast(req.Context(), db.UpdateExerciseLastParams{ID: eid, LastVolume: lastVolume})
}

// DeleteOneExercise deletes one workouts exercise
func (api *ExerciseAPI) DeleteOneExercise(res http.ResponseWriter, req *http.Request) {
	eid, errE := strconv.ParseInt(req.PathValue("eid"), 10, 64)
	// invalid user uid
	if errE != nil {
		util.ReturnErrorJSONResponse(res, "Invalid exercise id for fetching single workout", 400)
		return
	}

	errAfterDeletion := api.Store().Queries.DeleteSingleExercise(req.Context(), eid)

	if errAfterDeletion = util.CheckError(errAfterDeletion, res, req); errAfterDeletion != nil {
		return
	}

	util.ReturnValidJSONResponse(res, util.CreateSuccessResponse("Successfully deleted single user workouts", 200))
}

// DeleteAllExercises deletes one users exercise
func (api *ExerciseAPI) DeleteAllExercises(res http.ResponseWriter, req *http.Request) {
	wid, errW := strconv.ParseInt(req.Header.Get("wid"), 10, 64)
	// invalid workout uid
	if errW != nil {
		util.ReturnErrorJSONResponse(res, "Invalid workout id for deleting all workouts", 400)
		return
	}

	errAfterDeletion := api.Store().Queries.DeleteAllExercises(req.Context(), wid)

	if errAfterDeletion = util.CheckError(errAfterDeletion, res, req); errAfterDeletion != nil {
		return
	}

	util.ReturnValidJSONResponse(res, util.CreateSuccessResponse("Successfully deleted all workout exercises", 200))
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
