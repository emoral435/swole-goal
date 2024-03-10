package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

// CreateUser creates a new user, using their email, password, and username.
//
// This also stores their birthday and the time their account was created.
func CreateUser(res http.ResponseWriter, req *http.Request, store *db.Store) {
	res.Header().Set("Content-Type", "application/json")

	// what we need in the header
	arg := db.CreateUserParams{
		Email:    req.Header.Get("email"),
		Password: req.Header.Get("password"),
		Username: req.Header.Get("username"),
	}

	user, err := store.CreateUser(req.Context(), arg)

	// deal with bad request (params for creating user not satisfied)
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	// return user in the form of JSON
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(user)
}

// GetUserFromID returns user from the given ID string
func GetUserFromID(res http.ResponseWriter, req *http.Request, store *db.Store) {
	res.Header().Set("Content-Type", "application/json")
	// get the id query from URL
	id, err := strconv.ParseInt(req.PathValue("id"), 10, 64)

	// deal with bad request (query is invalid)
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	user, err := store.GetUser(req.Context(), id)

	// check if we got the user successfully
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	// send back the correct response
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(user)
}

func DeleteUser(res http.ResponseWriter, req *http.Request, store *db.Store, uid int64) {
	res.Header().Set("Content-Type", "application/json")

	// get -> user -> all workouts -> all exercises -> all sets
	// delete -> all sets -> all workouts -> all users

	err := store.DeleteUser(req.Context(), uid)

	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(util.CreateSuccessResponse("success", http.StatusOK))
}
