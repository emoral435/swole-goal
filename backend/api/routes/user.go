package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

func CreateUser(res http.ResponseWriter, req *http.Request, store *db.Store) {
	res.Header().Set("Content-Type", "application/json")

	arg := db.CreateUserParams{
		Email:    req.Header.Get("email"),
		Password: req.Header.Get("password"),
		Username: req.Header.Get("username"),
	}

	user, err := store.CreateUser(req.Context(), arg)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "Please try again, "+err.Error())
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(user)
}

func DeleteUser(res http.ResponseWriter, req *http.Request, store *db.Store, uid int64) {
	res.Header().Set("Content-Type", "application/json")

	// get -> user -> all workouts -> all exercises -> all sets
	// delete -> all sets -> all workouts -> all users

	err := store.DeleteUser(req.Context(), uid)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "Please try again, "+err.Error())
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(util.ResponseMessage{Message: "User deleted successfully"})
}
