package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCreateUser: tests CreateUser methods
//
// requires: NumUsers to work properly
func TestCreateUser(t *testing.T) {
	arg := GenRandUser()

	user, err := testQueries.CreateUser(context.Background(), *arg)
	numUsers, numErr := testQueries.NumUsers(context.Background())
	// if query failed
	require.NoError(t, err)
	require.NoError(t, numErr)
	require.NotEmpty(t, user)

	// checking the new user has the correct values in the table
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Username, user.Username)

	// print out the users id - should always be incrementing
	fmt.Println(user.ID)
	require.NotEqual(t, user.ID, numUsers-1)

	// ensure that the table is serializing the id's correctly
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}

// TestNumUsers: tests NumUsers
//
// requires: CreateUser to work properly
func TestNumUsers(t *testing.T) {
	// get the initial size of the users table
	numUsersBefore, numErrBefore := testQueries.NumUsers(context.Background())

	require.NoError(t, numErrBefore)
	require.NotEmpty(t, numUsersBefore)

	arg1 := GenRandUser()
	_, err1 := testQueries.CreateUser(context.Background(), *arg1)
	arg2 := GenRandUser()
	_, err2 := testQueries.CreateUser(context.Background(), *arg2)
	arg3 := GenRandUser()
	_, err3 := testQueries.CreateUser(context.Background(), *arg3)

	// this means the DI for CreateUser is not working
	if err1 != nil || err2 != nil || err3 != nil {
		require.Equal(t, 0, 1)
		return
	}

	numUsersAfter, numErrAfter := testQueries.NumUsers(context.Background())

	require.NoError(t, numErrAfter)
	require.NotEmpty(t, numUsersAfter)
	require.Equal(t, numUsersBefore+3, numUsersAfter)
}

// TestDeleteUser: tests DeleteUser method
//
// depends on TestCreateUser
func TestDeleteUser(t *testing.T) {
	// get the initial size of the users table
	numUsersBefore, numErrBefore := testQueries.NumUsers(context.Background())

	arg := GenRandUser()

	user, err := testQueries.CreateUser(context.Background(), *arg)

	// if we get some sort of error, retrun instantly, as CreateUser does not work
	// we can indicate this with returning 1
	if err != nil || numErrBefore != nil {
		require.Equal(t, 0, 1)
		return
	}

	// there should always be at least one user created, so the error should not be reported if working
	testErr := testQueries.DeleteUser(context.Background(), user.ID)

	numUsersAfter, numErrAfter := testQueries.NumUsers(context.Background())

	require.NoError(t, testErr)
	require.NoError(t, numErrAfter)
	require.Equal(t, numUsersBefore, numUsersAfter) // the table should now be the same starting size
}
