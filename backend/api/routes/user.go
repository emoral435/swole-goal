package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

func ServerUsers(mux *http.ServeMux, store *db.Store) {
	// creates a user using http headers
	mux.HandleFunc("POST /user", func(res http.ResponseWriter, req *http.Request) {
		CreateUser(res, req, store)
	})

	// gets a user using their id
	mux.HandleFunc("GET /user/id/{id}", func(res http.ResponseWriter, req *http.Request) {
		GetUserFromID(res, req, store)
	})

	// gets a user using their email
	mux.HandleFunc("GET /user/email/{email}", func(res http.ResponseWriter, req *http.Request) {
		GetUserFromEmail(res, req, store)
	})

	// updates a users information, a user that correlates to their UID/email (probably will be using a form)
	mux.HandleFunc("PUT /user/{id}", func(res http.ResponseWriter, req *http.Request) {
		UpdateUserInfo(res, req, store)
	})

	// deletes a single user
	mux.HandleFunc("DELETE /user/{id}", func(res http.ResponseWriter, req *http.Request) {
		DeleteUser(res, req, store)
	})
}

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

// GetUserFromEmail returns user from the given email path string
func GetUserFromEmail(res http.ResponseWriter, req *http.Request, store *db.Store) {
	res.Header().Set("Content-Type", "application/json")
	// get the email query from URL
	email := req.PathValue("email")

	user, err := store.GetUserEmail(req.Context(), email)

	// check if we got the user successfully
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	// send back the correct response
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(user)
}

func UpdateUserInfo(res http.ResponseWriter, req *http.Request, store *db.Store) {
	res.Header().Set("Content-Type", "application/json")
	// get the id query from URL
	id, err := strconv.ParseInt(req.PathValue("id"), 10, 64)

	// deal with bad request (query is invalid)
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	// what we need in the url query
	arg := db.UpdatePasswordParams{
		ID:       id,
		Password: req.Header.Get("password"),
	}

	if len(arg.Password) > 0 {
	}

}

// deletes a user, and all their information within the database
//
// this includes deleting their workouts, their exercises, and their sets
func DeleteUser(res http.ResponseWriter, req *http.Request, store *db.Store) {
	res.Header().Set("Content-Type", "application/json")

	id, err := strconv.ParseInt(req.PathValue("id"), 10, 64)

	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	err = store.DeleteUser(req.Context(), id)

	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(util.CreateSuccessResponse("success", http.StatusOK))
}
