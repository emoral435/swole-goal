package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	mw "github.com/emoral435/swole-goal/api/middleware"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

func ServerUsers(mux *http.ServeMux, ss *ServerStore) {
	// creates a user using http headers
	mux.Handle("POST /user", mw.EnforceJSONHandler(http.HandlerFunc(ss.CreateUser)))

	// gets a user by their id
	mux.Handle("GET /user/id/{id}", mw.EnforceJSONHandler(mw.AuthMiddleware(ss.TokenMaker, http.HandlerFunc(ss.GetUserFromID))))

	// // gets a user using their email
	// mux.HandleFunc("GET /user/email/{email}", func(res http.ResponseWriter, req *http.Request) {
	// 	GetUserFromEmail(res, req, store)
	// })

	// // updates a users information, a user that correlates to their UID/email (probably will be using a form)
	// mux.HandleFunc("PUT /user/{id}", func(res http.ResponseWriter, req *http.Request) {
	// 	UpdateUserInfo(res, req, store)
	// })

	// finalUpdateUserHandler := http.HandlerFunc(UpdateUserInfo)
	// mux.Handle("/", mw.EnforceJSONHandler(mw.AuthMiddleware(serverStore.TokenMaker, finalUpdateUserHandler)))

	// // deletes a single user
	// mux.HandleFunc("DELETE /user/{id}", func(res http.ResponseWriter, req *http.Request) {
	// 	DeleteUser(res, req, store)
	// })

	// // handles the authentication of a user with their JWT token
	mux.HandleFunc("POST /user/login", ss.LoginUser)
}

// CreateUser creates a new user, using their email, password, and username.
//
// This also stores their birthday and the time their account was created.
func (ss *ServerStore) CreateUser(res http.ResponseWriter, req *http.Request) {

	hashedPassword, err := util.HashPassword(req.Header.Get("password"))

	// deal with bad request (params for creating user not satisfied)
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	// what we need in the header
	arg := db.CreateUserParams{
		Email:    req.Header.Get("email"),
		Password: hashedPassword,
		Username: req.Header.Get("username"),
	}

	user, err := ss.Store.CreateUser(req.Context(), arg)

	// deal with bad request (params for creating user not satisfied)
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	// return user in the form of JSON
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(user)
}

// GetUserFromID returns user from the given ID string
func (ss *ServerStore) GetUserFromID(res http.ResponseWriter, req *http.Request) {

	// get the id query from URL
	id, err := strconv.ParseInt(req.PathValue("id"), 10, 64)

	// deal with bad request (query is invalid)
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	user, err := ss.Store.GetUser(req.Context(), id)

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
	//
	// // get the id query from URL
	// uid, err := strconv.ParseInt(req.PathValue("id"), 10, 64)

	// // deal with bad request (query is invalid)
	// if err = util.CheckError(err, res, req); err != nil {
	// 	return
	// }

	// new user information
	// newEmail := req.Header.Get("email")
	// newPassword := req.Header.Get("password")
	// newUsername := req.Header.Get("username")
	// newBirthday := req.Header.Get("birthday")

	// if len(newPassword) > 0 {
	// 	UpdatePassword(res, req, store, newPassword, uid)
	// }

	// // what we need in the url query
	// pswrdParams := db.UpdatePasswordParams{
	// 	ID:       id,
	// 	Password: req.Header.Get("password"),
	// }

	// bdayParams := db.UpdateBirthdayParams{
	// 	ID:       id,
	// 	Birthday: req.Header.Get("password"),
	// }

	// if len(arg.Password) > 0 {
	// }

}

func UpdatePassword(res http.ResponseWriter, req *http.Request, store *db.Store, newPassowrd string, uid int64) {
	// what we need in the url query
	// pswrdParams := db.UpdatePasswordParams{
	// 	ID:       id,
	// 	Password: req.Header.Get("password"),
	// }
}

// deletes a user, and all their information within the database
//
// this includes deleting their workouts, their exercises, and their sets
func DeleteUser(res http.ResponseWriter, req *http.Request, store *db.Store) {

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

// LoginUser returns the access token for the user, provided their email and password and if it is a valid email/password combination
func (ss *ServerStore) LoginUser(res http.ResponseWriter, req *http.Request) {

	user, err := ss.Store.GetUserEmail(req.Context(), req.Header.Get("email"))

	// check if we got the user successfully
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	err = util.CompareHash(user.Password, req.Header.Get("password"))

	// check if the hash correlation was successful
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	accessToken, err := ss.TokenMaker.CreateToken(
		user.Email,
		ss.Config.AccessTokenDuration,
	)

	// check if we got the user successfully
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	rsp := loginUserResponse{ // generates the response we want the client to recieve
		AccessToken: accessToken,
		User:        user,
	}

	// return user in the form of JSON
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(rsp)
}

type loginUserResponse struct {
	AccessToken string  `json:"access_token"`
	User        db.User `json:"user"`
}
