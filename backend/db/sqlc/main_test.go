package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	// NOTe this is used for dev environments
	// dbSource = "postgresql://root:secret@host.docker.internal:5432/swole_goal?sslmode=disable"
	dbSource = "postgresql://root:secret@localhost:5432/swole_goal?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
