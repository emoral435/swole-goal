package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/emoral435/swole-goal/api/routes"
	db "github.com/emoral435/swole-goal/db/sqlc"
)

// Serve serves the API for our application
//
// defines the server muxiplexer interface, and using package router, defines routes and handlers
func Serve(connection *sql.DB) {
	muxRouter := http.NewServeMux()

	store := db.NewStore(connection)

	// defines all our routes and the handlers for each route and METHOD
	serveRoutes(muxRouter, store)

	fmt.Println("Server started!")
	// starts the server
	if err := http.ListenAndServe("localhost:9090", muxRouter); err != nil {
		fmt.Printf("Something went wrong!")
		return
	}
}

func serveRoutes(mux *http.ServeMux, store *db.Store) {
	// just for me hehe
	mux.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Server started."))
	})

	// creates a user using http headers
	mux.HandleFunc("POST /user", func(res http.ResponseWriter, req *http.Request) {
		routes.CreateUser(res, req, store)
	})

	// gets a single user
	mux.HandleFunc("GET /user/{id}", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "server hath started")
		fmt.Fprintf(res, "server hath started, with an id of %s", req.PathValue("id"))
	})
}
