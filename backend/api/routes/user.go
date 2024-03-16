package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/araddon/dateparse"
	mw "github.com/emoral435/swole-goal/api/middleware"
	"github.com/emoral435/swole-goal/api/token"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

func ServerUsers(mux *http.ServeMux, ss *ServerStore) {
	uAPI := createUserAPIStruct(ss.TokenMaker, ss.Store, ss.Config) // implements IUser interface
	// creates a user using http headers
	mux.Handle("POST /user", mw.EnforceJSONHandler(http.HandlerFunc(uAPI.CreateUser)))

	// gets a user by their id
	mux.Handle("GET /user/id/{id}", mw.EnforceJSONHandler(mw.AuthMiddleware(uAPI.TokenMaker(), http.HandlerFunc(uAPI.GetUserFromID))))

	// gets a user using their email
	mux.Handle("GET /user/email/{email}", mw.EnforceJSONHandler(mw.AuthMiddleware(uAPI.TokenMaker(), http.HandlerFunc(uAPI.GetUserFromEmail))))

	// updates a users information, a user that correlates to their UID/email (probably will be using a form)
	mux.Handle("PUT /user/{id}", mw.EnforceJSONHandler(http.HandlerFunc(uAPI.UpdateUserInfo)))

	// deletes a single user
	mux.Handle("DELETE /user/{id}", mw.EnforceJSONHandler(mw.AuthMiddleware(uAPI.TokenMaker(), http.HandlerFunc(uAPI.DeleteUser))))

	// handles the authentication of a user with their JWT token
	mux.HandleFunc("POST /user/login", uAPI.LoginUser)
}

// CreateUser creates a new user, using their email, password, and username.
//
// This also stores their birthday and the time their account was created.
func (api *UserAPI) CreateUser(res http.ResponseWriter, req *http.Request) {

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

	user, err := api.Store().CreateUser(req.Context(), arg)

	// deal with bad request (params for creating user not satisfied)
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	util.ReturnValidJSONResponse(res, user)
}

// GetUserFromID returns user from the given ID string
func (api *UserAPI) GetUserFromID(res http.ResponseWriter, req *http.Request) {

	// get the id query from URL
	id, err := strconv.ParseInt(req.PathValue("id"), 10, 64)

	// deal with bad request (query is invalid)
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	user, err := api.Store().GetUser(req.Context(), id)

	// check if we got the user successfully
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	util.ReturnValidJSONResponse(res, user)
}

// GetUserFromEmail returns user from the given email path string
func (api *UserAPI) GetUserFromEmail(res http.ResponseWriter, req *http.Request) {

	// get the email query from URL
	email := req.PathValue("email")

	user, err := api.Store().GetUserEmail(req.Context(), email)

	// check if we got the user successfully
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	util.ReturnValidJSONResponse(res, user)
}

func (api *UserAPI) UpdateUserInfo(res http.ResponseWriter, req *http.Request) {

	// get the id query from URL
	uid, err := strconv.ParseInt(req.PathValue("id"), 10, 64)

	// deal with bad request (query is invalid)
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	currentUserInfo, userExists := api.Store().GetUser(req.Context(), uid)

	// deal with bad request (query is invalid)
	if userExists = util.CheckError(userExists, res, req); userExists != nil {
		return
	}

	// new user information
	bdayFormattedTime, isValid := parseBirthday(req.Header.Get("birthday"))
	newBirthday := sql.NullTime{
		Time:  bdayFormattedTime,
		Valid: isValid,
	}

	newPassword := req.Header.Get("password")
	newUserInfo := db.CreateNewUserInfo(req.Header.Get("email"), req.Header.Get("username"), newBirthday)

	if len(newPassword) > 0 {
		// UpdatePassword(res, req, ss, newPassword, uid)
		_, err := api.Store().UpdatePassword(req.Context(), *db.CreateUpdatePasswordParam(uid, newPassword))
		if err = util.CheckError(err, res, req); err != nil {
			return
		}
	}

	if len(newUserInfo.Username) > 0 {
		// UpdatePassword(res, req, ss, newPassword, uid)
		_, err := api.Store().UpdateUsername(req.Context(), *db.CreateUpdateUsernameParam(uid, newUserInfo.Username))
		if err = util.CheckError(err, res, req); err != nil {
			newUserInfo.Username = currentUserInfo.Username
			return
		}
	}

	if len(newUserInfo.Email) > 0 {
		// UpdatePassword(res, req, ss, newPassword, uid)
		_, err := api.Store().UpdateEmail(req.Context(), *db.CreateUpdateEmailParam(uid, newUserInfo.Email))
		if err = util.CheckError(err, res, req); err != nil {
			newUserInfo.Email = currentUserInfo.Email
			return
		}
	}

	if newUserInfo.Birthday.Valid {
		// UpdatePassword(res, req, ss, newPassword, uid)
		_, err := api.Store().UpdateBirthday(req.Context(), *db.CreateUpdateBirthdayParams(uid, newUserInfo.Birthday))
		if err = util.CheckError(err, res, req); err != nil {
			newUserInfo.Birthday = currentUserInfo.Birthday
			return
		}
	}

	util.ReturnValidJSONResponse(res, newUserInfo)
}

// parseBirthday generates a sql.Time struct, given the input string
//
// if the input string is not convertable using dateparse library, returns err, and thus, returns an empty time struct and a boolean indicating invalid time.
// Otherwise, returns the user-input string as a time.Time struct and a boolean indicating it was successfully parsed
func parseBirthday(bday string) (time.Time, bool) {
	newBday, err := dateparse.ParseStrict(bday)
	if err != nil {
		return time.Time{}, false
	}
	return newBday, true
}

// deletes a user, and all their information within the database
//
// this includes deleting their workouts, their exercises, and their sets
func (api *UserAPI) DeleteUser(res http.ResponseWriter, req *http.Request) {

	id, err := strconv.ParseInt(req.PathValue("id"), 10, 64)

	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	err = api.Store().DeleteUser(req.Context(), id)

	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(util.CreateSuccessResponse("success", http.StatusOK))
}

// LoginUser returns the access token for the user, provided their email and password and if it is a valid email/password combination
func (api *UserAPI) LoginUser(res http.ResponseWriter, req *http.Request) {

	user, err := api.Store().GetUserEmail(req.Context(), req.Header.Get("email"))

	// check if we got the user successfully
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	err = util.CompareHash(user.Password, req.Header.Get("password"))

	// check if the hash correlation was successful
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	accessToken, err := api.TokenMaker().CreateToken(
		user.Email,
		api.Config().AccessTokenDuration,
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

// loginUserResponse generates a struct that contains the access token for a user to utilize the API.
//
// This also contains the user as a item within the struct, which also has the users information
type loginUserResponse struct {
	AccessToken string  `json:"access_token"`
	User        db.User `json:"user"`
}

// UserAPI represents the API actions taken for each route for a user.
//
// Implements the IRouteStore interface.
type UserAPI struct {
	tokenMaker token.Maker
	store      *db.Store
	config     util.Config
}

// createUserAPI creates a new UserAPI struct instance
func createUserAPIStruct(t token.Maker, store *db.Store, config util.Config) IUser {
	return &UserAPI{
		t,
		store,
		config,
	}
}

func (u *UserAPI) TokenMaker() token.Maker {
	return u.tokenMaker
}

func (u *UserAPI) Store() *db.Store {
	return u.store
}

func (u *UserAPI) Config() util.Config {
	return u.config

}
