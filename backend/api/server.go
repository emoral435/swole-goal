package server

import (
	"database/sql"
	"fmt"
	"log"
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
		log.Fatalf("We recieved and error: %s", err.Error())
	}
}

func serveRoutes(mux *http.ServeMux, store *db.Store) {
	// creates a user
	mux.HandleFunc("POST /user", func(write http.ResponseWriter, read *http.Request) {
		routes.CreateUser(write, read, store)
	})
}
