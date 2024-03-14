package routes

import "net/http"

func ServeWorkouts(mux *http.ServeMux, ss *ServerStore) {
	// TODO create a workout
	mux.Handle()
	// TODO get all users workouts

	// TODO get a specific workout

	// TODO modify a users workout

	// TODO delete a users workout

	// TODO delete all user's workouts
}

// CreateWorkout creates a new workout for the associated and authenticated user.
//
// no information is needed, but we should be able to return the workout struct containing
// the workout information in case we want to be able to make a form popup on the creation of the workout.
func (ss *ServerStore) CreateWorkout(res *http.ResponseWriter, req *http.Request) {
	// TODO
}

// GetAllWorkouts returns a list of all workout's in the server store
func (ss *ServerStore) GetAllWorkouts(res *http.ResponseWriter, req *http.Request) {
	// TODO
}

// GetOneWorkout returns a workout based on workout id within the database
func (ss *ServerStore) GetOneWorkout(res *http.ResponseWriter, req *http.Request) {
	// TODO
}

// ModifyWorkout updates a users workout
func (ss *ServerStore) ModifyWorkout(res *http.ResponseWriter, req *http.Request) {
	// TODO
}

// DeleteOneWorkout deletes one users workout
func (ss *ServerStore) DeleteOneWorkout(res *http.ResponseWriter, req *http.Request) {
	// TODO
}

// DeleteAllWorkouts deletes one users workout
func (ss *ServerStore) DeleteAllWorkouts(res *http.ResponseWriter, req *http.Request) {
	// TODO
}
