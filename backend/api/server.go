package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/emoral435/swole-goal/api/routes"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

// Serve serves the API for our application
//
// defines the server muxiplexer interface, and using package router, defines routes and handlers
func Serve(connection *sql.DB, config util.Config) {
	muxRouter := http.NewServeMux()

	store := db.NewStore(connection)

	// defines all our routes and the handlers for each route and METHOD
	serveRoutes(muxRouter, store)

	fmt.Println("Server started!")
	// starts the server
	if err := http.ListenAndServe(config.ServerAddress, muxRouter); err != nil {
		fmt.Printf("Something went wrong!")
		return
	}
}

func serveRoutes(mux *http.ServeMux, store *db.Store) {
	// just for me hehe
	mux.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(util.SuccessResponse{Message: "Hello world!", Status: http.StatusOK})
	})

	// creates a user using http headers
	mux.HandleFunc("POST /user", func(res http.ResponseWriter, req *http.Request) {
		routes.CreateUser(res, req, store)
	})

	// gets a user using their id
	mux.HandleFunc("GET /user/id/{id}", func(res http.ResponseWriter, req *http.Request) {
		routes.GetUserFromID(res, req, store)
	})

	// gets a user using their email
	mux.HandleFunc("GET /user/email/{email}", func(res http.ResponseWriter, req *http.Request) {
		routes.GetUserFromEmail(res, req, store)
	})

	// deletes a single user
	mux.HandleFunc("DELETE /user/{id}", func(res http.ResponseWriter, req *http.Request) {
		id, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
		routes.DeleteUser(res, req, store, id)
	})
}
