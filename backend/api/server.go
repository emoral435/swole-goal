package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/emoral435/swole-goal/api/routes"
	util "github.com/emoral435/swole-goal/utils"
)

// Serve serves the API for our application
//
// defines the server muxiplexer interface, and using package router, defines routes and handlers
func Serve(connection *sql.DB, config util.Config, serverStore *routes.ServerStore) {
	muxRouter := http.NewServeMux()

	// defines all our routes and the handlers for each route and METHOD
	serveRoutes(muxRouter, serverStore)

	// successful start - everything connecting went well
	fmt.Println("Server starting!")

	// starts the server
	if err := http.ListenAndServe(config.ServerAddress, muxRouter); err != nil {
		fmt.Printf("Something went wrong!")
		return
	}
}

// the place where all routes of the API are registered
func serveRoutes(mux *http.ServeMux, serverStore *routes.ServerStore) {
	routes.ServerUsers(mux, serverStore)    // found in routes/user.go
	routes.ServeWorkouts(mux, serverStore)  // found in routes/workout.go
	routes.ServeExercises(mux, serverStore) // found in routes/exercise.go
}
