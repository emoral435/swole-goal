package routes

import (
	"fmt"
	"net/http"

	db "github.com/emoral435/swole-goal/db/sqlc"
)

func CreateUser(w http.ResponseWriter, r *http.Request, store *db.Store) {

	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	username := r.URL.Query().Get("username")
	var bday string
	if birthday := r.URL.Query().Get("birthday"); len(birthday) == 0 {
		bday = "no birthday provided :("
	}
	fmt.Fprintf(w, "Method is %s\n", r.Method)
	fmt.Fprintf(w, "Email = %s\n", email)
	fmt.Fprintf(w, "password = %s\n", password)
	fmt.Fprintf(w, "username = %s\n", username)
	fmt.Fprintf(w, "birthday = %s\n", bday)

	arg := db.CreateUserParams{
		Email:    email,
		Password: password,
		Username: username,
	}

	store.CreateUser(r.Context(), arg)
}
