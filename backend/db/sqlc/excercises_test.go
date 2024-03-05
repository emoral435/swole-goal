package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCreateExercise: tests CreateExercise methods
//
// requires: GetUserWorkouts
func TestCreateExercise(t *testing.T) {
	TestCreateWorkout(t)
	newExercise := GenRandExercise(1)
	exeRow, err := testQueries.CreateExercise(context.Background(), *newExercise)

	require.NoError(t, err)
	require.Equal(t, newExercise.Title, exeRow.Title)
	require.Equal(t, newExercise.Description, exeRow.Description)
	require.Equal(t, newExercise.LastVolume, exeRow.LastVolume)
	require.Equal(t, int64(1), exeRow.WorkoutID)
}

// TestDeleteAllExercises: tests DeleteAllExercises methods
//
// requires: GetUserWorkouts
func TestDeleteAllExercises(t *testing.T) {
	TestCreateExercise(t)
	err := testQueries.DeleteAllExercises(context.Background(), 64)

	require.NoError(t, err)
}

// TestDeleteSingleExercise: tests DeleteSingleExercise methods
//
// requires: GetUserWorkouts
func TestDeleteSingleExercise(t *testing.T) {
	TestCreateExercise(t)
	err := testQueries.DeleteSingleExercise(context.Background(), 1)

	require.NoError(t, err)
}
