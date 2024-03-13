package main

import (
	"database/sql"
	"fmt"
	"log"

	server "github.com/emoral435/swole-goal/api"
	"github.com/emoral435/swole-goal/api/routes"
	"github.com/emoral435/swole-goal/api/token"
	db "github.com/emoral435/swole-goal/db/sqlc"
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

	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)

	if err != nil {
		log.Fatal("Bad token...")
		return
	}

	store := db.NewStore(conn)

	s := routes.CreateServerStore(tokenMaker, config, store)

	// serve our api with our database connection
	server.Serve(conn, config, s)
}
