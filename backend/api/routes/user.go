package routes

import (
	"net/http"

	db "github.com/emoral435/swole-goal/db/sqlc"
)

func CreateUser(res http.ResponseWriter, req *http.Request, store *db.Store) {

	arg := db.CreateUserParams{
		Email:    req.Header.Get("email"),
		Password: req.Header.Get("password"),
		Username: req.Header.Get("username"),
	}

	_, err := store.CreateUser(req.Context(), arg)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Email already in use."))
	}
}
