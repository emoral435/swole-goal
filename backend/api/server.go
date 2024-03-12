package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emoral435/swole-goal/api/routes"
	"github.com/emoral435/swole-goal/api/token"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

type Server struct {
	config     util.Config
	tokenMaker token.Maker
}

type loginUserResponse struct {
	AccessToken string  `json:"access_token"`
	User        db.User `json:"user"`
}

func (s *Server) loginUser(res http.ResponseWriter, req *http.Request, store *db.Store) {
	res.Header().Set("Content-Type", "application/json")

	user, err := store.GetUserEmail(req.Context(), req.Header.Get("email"))

	// check if we got the user successfully
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	err = util.CompareHash(user.Password, req.Header.Get("password"))
	// check if we got the user successfully
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	accessToken, err := s.tokenMaker.CreateToken(
		user.Email,
		s.config.AccessTokenDuration,
	)

	// check if we got the user successfully
	if err = util.CheckError(err, res, req); err != nil {
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        user,
	}

	// return user in the form of JSON
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(rsp)
}

// Serve serves the API for our application
//
// defines the server muxiplexer interface, and using package router, defines routes and handlers
func Serve(connection *sql.DB, config util.Config) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)

	if err != nil {
		fmt.Println("Bad token...")
		return
	}

	server := &Server{
		tokenMaker: tokenMaker,
		config:     config,
	}

	muxRouter := http.NewServeMux()

	store := db.NewStore(connection)

	// defines all our routes and the handlers for each route and METHOD
	serveRoutes(muxRouter, store, server)

	fmt.Println("Server started!")
	// starts the server
	if err := http.ListenAndServe(config.ServerAddress, muxRouter); err != nil {
		fmt.Printf("Something went wrong!")
		return
	}
}

func serveRoutes(mux *http.ServeMux, store *db.Store, s *Server) {
	mux.HandleFunc("POST /user/login", func(res http.ResponseWriter, req *http.Request) {
		s.loginUser(res, req, store)
	})
	// just for me hehe
	routes.ServerUsers(mux, store)
}
