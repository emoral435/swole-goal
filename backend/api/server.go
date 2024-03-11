package server

import (
	"database/sql"
	"fmt"
	"net/http"

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
	routes.ServerUsers(mux, store)
}
