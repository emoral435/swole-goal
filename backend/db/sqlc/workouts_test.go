package db

import (
	"context"
	"testing"

	util "github.com/emoral435/swole-goal/utils"
	"github.com/stretchr/testify/require"
)

// TestCreateWorkout: tests CreateWorkout methods
//
// requires: GetUserWorkouts
func TestCreateWorkout(t *testing.T) {
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

// TestGetUserWorkouts: tests GetUserWorkouts methods
//
// requires: CreateUser
func TestGetUserWorkouts(t *testing.T) {
	arg := GenRandUser()
	user, err := testQueries.CreateUser(context.Background(), *arg)
	if err != nil {
		return
	}
	newWorkout := GenRandWorkout(user.ID)
	workout, err := testQueries.CreateWorkout(context.Background(), *newWorkout)
	if err != nil {
		return
	}
	require.NotNil(t, workout)
	userWorkouts, err := testQueries.GetUserWorkouts(context.Background(), user.ID)
	// if query failed
	require.NoError(t, err)
	require.NotEmpty(t, userWorkouts)
	require.Equal(t, len(userWorkouts), 1)
}

// TestGetWorkouts: tests GetWorkouts methods
//
// requires: CreateUser
func TestGetWorkouts(t *testing.T) {
	arg := GenRandUser()
	user, err := testQueries.CreateUser(context.Background(), *arg)
	if err != nil {
		return
	}
	newWorkout := GenRandWorkout(user.ID)
	workout, err := testQueries.CreateWorkout(context.Background(), *newWorkout)
	if err != nil {
		return
	}
	require.NotNil(t, workout)
	userWorkouts, err := testQueries.GetWorkout(context.Background(), workout.ID)
	// if query failed
	require.NoError(t, err)
	require.NotEmpty(t, userWorkouts)
}

// TestDeleteAllWorkouts: tests DeleteAllWorkouts methods
//
// requires: CreateUser, CreateWorkout, GetUserWorkouts
func TestDeleteAllWorkouts(t *testing.T) {
	arg := GenRandUser()
	user, err := testQueries.CreateUser(context.Background(), *arg)
	if err != nil {
		return
	}
	w1 := GenRandWorkout(user.ID)
	w2 := GenRandWorkout(user.ID)
	_, err1 := testQueries.CreateWorkout(context.Background(), *w1)
	_, err2 := testQueries.CreateWorkout(context.Background(), *w2)
	if err1 != nil || err2 != nil {
		return
	}

	testQueries.DeleteAllWorkouts(context.Background(), user.ID)
	checkWorkout, err := testQueries.GetUserWorkouts(context.Background(), user.ID)

	require.NoError(t, err)
	require.Equal(t, len(checkWorkout), 0)
}

// TestDeleteSingleWorkout: tests DeleteSingleWorkout methods
//
// requires: CreateUser, CreateWorkout, GetUserWorkouts
func TestDeleteSingleWorkout(t *testing.T) {
	arg := GenRandUser()
	user, err := testQueries.CreateUser(context.Background(), *arg)
	if err != nil {
		return
	}
	w1 := GenRandWorkout(user.ID)
	w2 := GenRandWorkout(user.ID)
	workout1, err1 := testQueries.CreateWorkout(context.Background(), *w1)
	_, err2 := testQueries.CreateWorkout(context.Background(), *w2)
	if err1 != nil || err2 != nil {
		return
	}

	testQueries.DeleteSingleWorkout(context.Background(), workout1.ID)
	checkWorkout, err := testQueries.GetUserWorkouts(context.Background(), user.ID)

	require.NoError(t, err)
	require.Equal(t, len(checkWorkout), 1)
}

// TestUpdateWorkoutBody: tests UpdateWorkoutBody methods
//
// requires: CreateUser, CreateWorkout, GetUserWorkouts
func TestUpdateWorkoutBody(t *testing.T) {
	arg := GenRandUser()
	user, err := testQueries.CreateUser(context.Background(), *arg)
	if err != nil {
		return
	}
	w1 := GenRandWorkout(user.ID)
	workout, err := testQueries.CreateWorkout(context.Background(), *w1)
	if err != nil {
		return
	}
	workoutParams := UpdateWorkoutBodyParams{
		ID:   workout.ID,
		Body: util.RandomString(10),
	}
	lWorkout, err := testQueries.UpdateWorkoutBody(context.Background(), workoutParams)
	checkWorkout, _ := testQueries.GetUserWorkouts(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, len(checkWorkout), 1)
	require.NotEqual(t, lWorkout.Body, w1.Body)
}

// TestUpdateWorkoutLast: tests UpdateWorkoutLast methods
//
// requires: CreateUser, CreateWorkout, GetUserWorkouts
func TestUpdateWorkoutLast(t *testing.T) {
	arg := GenRandUser()
	user, err := testQueries.CreateUser(context.Background(), *arg)
	if err != nil {
		return
	}
	w1 := GenRandWorkout(user.ID)
	workout, err := testQueries.CreateWorkout(context.Background(), *w1)
	if err != nil {
		return
	}
	workoutParams := UpdateWorkoutLastParams{
		ID:       workout.ID,
		LastTime: user.Birthday.Time,
	}
	lWorkout, err := testQueries.UpdateWorkoutLast(context.Background(), workoutParams)
	checkWorkout, _ := testQueries.GetUserWorkouts(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, len(checkWorkout), 1)
	require.NotEqual(t, lWorkout.LastTime, w1.LastTime)
}

// TestUpdateWorkoutTitle: tests UpdateWorkoutTitle methods
//
// requires: CreateUser, CreateWorkout, GetUserWorkouts
func TestUpdateWorkoutTitle(t *testing.T) {
	arg := GenRandUser()
	user, err := testQueries.CreateUser(context.Background(), *arg)
	if err != nil {
		return
	}
	w1 := GenRandWorkout(user.ID)
	workout, err := testQueries.CreateWorkout(context.Background(), *w1)
	if err != nil {
		return
	}
	workoutParams := UpdateWorkoutTitleParams{
		ID:    workout.ID,
		Title: util.RandomString(10),
	}
	lWorkout, err := testQueries.UpdateWorkoutTitle(context.Background(), workoutParams)
	checkWorkout, _ := testQueries.GetUserWorkouts(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, len(checkWorkout), 1)
	require.NotEqual(t, lWorkout.Title, w1.Title)
}
