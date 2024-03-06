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

// GenRandWorkout: generates random workouts parameters
//
// returns: a CreateUserParams object with random values
func GenRandWorkout(uid int64) *CreateWorkoutParams {
	return CreateWorkoutParam(uid, util.RandomString(12), util.RandomString(12), time.Now())
}

// CreateExerciseParam: generates a user parameters
//
// returns: a CreateExerciseParams object with specified parameters
func CreateExerciseParam(uid int64, t string, d sql.NullString, ltv int64) *CreateExerciseParams {
	return &CreateExerciseParams{ID: uid, Type: t, Description: d, LastVolume: ltv}
}

// GenRandExercise: generates random user parameters
//
// returns: a CreateUserParams object with random values
func GenRandExercise(uid int64) *CreateExerciseParams {
	return CreateExerciseParam(uid, util.RandomString(12), util.RandomStringNull(12), util.RandomInt(0, 10))
}

// TestMain: the main test driver. Opens the local database, generates the queries, then runs each test
//
// returns: nothing
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("Cannot load env: ", err)
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
