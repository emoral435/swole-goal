package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCreateWorkout: tests CreateWorkout methods
//
// requires: GetUserWorkouts
func TestCreateExercise(t *testing.T) {
	arg := GenRandUser()
	user, err := testQueries.CreateUser(context.Background(), *arg)
	if err != nil {
		return
	}
	newWorkout := GenRandWorkout(user.ID)
	workout, err := testQueries.CreateWorkout(context.Background(), *newWorkout)
	// if query failed
	require.NoError(t, err)
	require.NotEmpty(t, workout)

	// checking the new user has the correct values in the table
	require.Equal(t, workout.UserID, user.ID)

	checkWorkout, err := testQueries.GetUserWorkouts(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, checkWorkout)

	require.Equal(t, workout.ID, checkWorkout[0].ID)
}
