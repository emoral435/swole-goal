package routes

import (
	"net/http"

	"github.com/emoral435/swole-goal/api/token"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

type IUser interface {
	TokenMaker() token.Maker
	Store() *db.Store
	Config() util.Config

	// CreateUser creates a new user within the database at "POST /user"
	CreateUser(res http.ResponseWriter, req *http.Request)
	// GetUserFromID returns in JSON the user from their ID in the database at "GET /user/id/{id}"
	GetUserFromID(res http.ResponseWriter, req *http.Request)
	// GetUserFromEmail returns user from the given email path string at "GET /user/email/{email}"
	GetUserFromEmail(res http.ResponseWriter, req *http.Request)
	// UpdateUserInfo updates user information (all fields) in the database at "PUT /user/{id}"
	UpdateUserInfo(res http.ResponseWriter, req *http.Request)
	// DeleteUser deletes user from the database, at "DELETE /user/{id}"
	DeleteUser(res http.ResponseWriter, req *http.Request)
	// LoginUser checks if the user correctly logged into the database, and correctly returns a token for the session at "POST /user/login"
	LoginUser(res http.ResponseWriter, req *http.Request)
}
