package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	util "github.com/emoral435/swole-goal/utils"

	_ "github.com/lib/pq"
)

var testQueries *Queries

// CreateUserParam: generates a user parameters (email, password, username)
//
// returns: a CreateUserParams object with specified parameters
func CreateUserParam(e string, p string, u string) *CreateUserParams {
	return &CreateUserParams{Email: e, Password: p, Username: u}
}

// GenRandUser: generates random user parameters (email, password, username)
//
// returns: a CreateUserParams object with random values (email, password, username)
func GenRandUser() *CreateUserParams {
	return CreateUserParam(util.RandomEmail(15), util.RandomPassword(12), util.RandomUsername(12))
}

// CreateWorkoutParam: generates a user parameters (email, password, username)
//
// returns: a CreateWorkoutParams object with specified parameters
func CreateWorkoutParam(uid int64, t string, b string, ltv time.Time) *CreateWorkoutParams {
	return &CreateWorkoutParams{ID: uid, Title: t, Body: b, LastTime: ltv}
}

// GenRandWorkout: generates random user parameters (email, password, username)
//
// returns: a CreateUserParams object with random values (email, password, username)
func GenRandWorkout(uid int64) *CreateWorkoutParams {
	return CreateWorkoutParam(uid, util.RandomString(12), util.RandomString(12), time.Now())
}

// drivers to connect to databse to check
const (
	dbDriver = "postgres"
	// NOTE this is used for dev environments
	// dbSource = "postgresql://root:secret@host.docker.internal:5432/swole_goal?sslmode=disable"
	dbSource = "postgresql://root:secret@localhost:5432/swole_goal?sslmode=disable"
)

// TestMain: the main test driver. Opens the local database, generates the queries, then runs each test
//
// returns: nothing
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
