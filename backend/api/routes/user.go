package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/araddon/dateparse"
	mw "github.com/emoral435/swole-goal/api/middleware"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

func ServerUsers(mux *http.ServeMux, ss *ServerStore) {
	// creates a user using http headers
	mux.Handle("POST /user", mw.EnforceJSONHandler(http.HandlerFunc(ss.CreateUser)))

	// gets a user by their id
	mux.Handle("GET /user/id/{id}", mw.EnforceJSONHandler(mw.AuthMiddleware(ss.TokenMaker, http.HandlerFunc(ss.GetUserFromID))))

	// gets a user using their email
	mux.Handle("GET /user/email/{email}", mw.EnforceJSONHandler(mw.AuthMiddleware(ss.TokenMaker, http.HandlerFunc(ss.GetUserFromEmail))))

	// updates a users information, a user that correlates to their UID/email (probably will be using a form)
	mux.Handle("PUT /user/{id}", mw.EnforceJSONHandler(http.HandlerFunc(ss.UpdateUserInfo)))

	// deletes a single user
	mux.Handle("DELETE /user/{id}", mw.EnforceJSONHandler(mw.AuthMiddleware(ss.TokenMaker, http.HandlerFunc(ss.DeleteUser))))

	// handles the authentication of a user with their JWT token
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
func (ss *ServerStore) GetUserFromEmail(res http.ResponseWriter, req *http.Request) {

	// get the email query from URL
	email := req.PathValue("email")

	user, err := ss.Store.GetUserEmail(req.Context(), email)

	// check if we got the user successfully
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	// send back the correct response
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(user)
}

func (ss *ServerStore) UpdateUserInfo(res http.ResponseWriter, req *http.Request) {

	// get the id query from URL
	uid, err := strconv.ParseInt(req.PathValue("id"), 10, 64)

	// deal with bad request (query is invalid)
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	currentUserInfo, userExists := ss.Store.GetUser(req.Context(), uid)

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
		_, err := ss.Store.UpdatePassword(req.Context(), *db.CreateUpdatePasswordParam(uid, newPassword))
		if err = util.CheckError(err, res, req); err != nil {
			return
		}
	}

	if len(newUserInfo.Username) > 0 {
		// UpdatePassword(res, req, ss, newPassword, uid)
		_, err := ss.Store.UpdateUsername(req.Context(), *db.CreateUpdateUsernameParam(uid, newUserInfo.Username))
		if err = util.CheckError(err, res, req); err != nil {
			newUserInfo.Username = currentUserInfo.Username
			return
		}
	}

	if len(newUserInfo.Email) > 0 {
		// UpdatePassword(res, req, ss, newPassword, uid)
		_, err := ss.Store.UpdateEmail(req.Context(), *db.CreateUpdateEmailParam(uid, newUserInfo.Email))
		if err = util.CheckError(err, res, req); err != nil {
			newUserInfo.Email = currentUserInfo.Email
			return
		}
	}

	if newUserInfo.Birthday.Valid {
		// UpdatePassword(res, req, ss, newPassword, uid)
		_, err := ss.Store.UpdateBirthday(req.Context(), *db.CreateUpdateBrithdayParams(uid, newUserInfo.Birthday))
		if err = util.CheckError(err, res, req); err != nil {
			newUserInfo.Birthday = currentUserInfo.Birthday
			return
		}
	}

	// send back the correct response
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(newUserInfo)
}

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
func (ss *ServerStore) DeleteUser(res http.ResponseWriter, req *http.Request) {

	id, err := strconv.ParseInt(req.PathValue("id"), 10, 64)

	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	err = ss.Store.DeleteUser(req.Context(), id)

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
