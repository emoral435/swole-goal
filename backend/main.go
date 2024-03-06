package main

import (
	"database/sql"
	"fmt"
	"log"

	server "github.com/emoral435/swole-goal/api"
	util "github.com/emoral435/swole-goal/utils"
	_ "github.com/lib/pq" // needed to connect to database -> https://stackoverflow.com/a/52791919/19919302
)

// the main server function for swole-goal
func main() {
	fmt.Println("Building REST API for swole goal...")

	// get env details
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load env: ", err)
		return
	}

	// connect to database server
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
		return
	}

	// serve our api with our database connection
	server.Serve(conn)
}
