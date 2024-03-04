package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	util "github.com/emoral435/swole-goal/utils"

	"github.com/stretchr/testify/require"
)

// TestCreateUserParam: tests our helper struct maker function
//
// requires: the utils random.go functions to work
func TestCreateUserParam(t *testing.T) {
	email := "email@example.com"
	password := "password"
	username := "username"
	userStruct := CreateUserParam(email, password, username)
	require.Equal(t, userStruct.Email, email)
	require.Equal(t, userStruct.Password, password)
	require.Equal(t, userStruct.Username, username)
}

// TestGenRandUser: tests our helper struct maker function
//
// requires: the utils random.go functions to work
func TestGenRandUser(t *testing.T) {
	user := GenRandUser()
	require.NotNil(t, user)
}

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
		return
	}

	// there should always be at least one user created, so the error should not be reported if working
	testErr := testQueries.DeleteUser(context.Background(), user.ID)

	numUsersAfter, numErrAfter := testQueries.NumUsers(context.Background())

	require.NoError(t, testErr)
	require.NoError(t, numErrAfter)
	require.Equal(t, numUsersBefore, numUsersAfter) // the table should now be the same starting size
}

// TestGetUser: tests GetUser method
//
// depends on CreateUser
func TestGetUser(t *testing.T) {
	arg := GenRandUser()

	user, err := testQueries.CreateUser(context.Background(), *arg)
	// if we get some sort of error, retrun instantly, as CreateUser does not work
	// we can indicate this with returning 1
	if err != nil {
		return
	}

	userFromID, Err := testQueries.GetUser(context.Background(), user.ID)

	// check if the user that we JUST made can be retrieved using the same
	require.NoError(t, Err)
	require.NotNil(t, userFromID)
	require.Equal(t, userFromID.ID, user.ID)
}

// TestGetUserEmail: tests GetUserEmail method
//
// depends on CreateUser
func TestGetUserEmail(t *testing.T) {
	arg := GenRandUser()

	user, err := testQueries.CreateUser(context.Background(), *arg)
	// if we get some sort of error, retrun instantly, as CreateUser does not work
	// we can indicate this with returning 1
	if err != nil {
		return
	}

	userFromEmail, emailErr := testQueries.GetUserEmail(context.Background(), user.Email)

	// check if the user that we JUST made can be retrieved using the same email
	require.NoError(t, emailErr)
	require.NotNil(t, userFromEmail)
	require.Equal(t, userFromEmail.Email, user.Email)
}

// TestListUsers: tests ListUsers method
//
// depends on CreateUser, NumUsers
func TestListUsers(t *testing.T) {
	arg := GenRandUser()

	_, userErr := testQueries.CreateUser(context.Background(), *arg)

	// get the initial size of the users table
	numUsers, numErr := testQueries.NumUsers(context.Background())
	// if we get some sort of error, retrun instantly, as CreateUser does not work
	// we can indicate this with returning 1
	if numErr != nil || userErr != nil {
		return
	}

	users, err := testQueries.ListUsers(context.Background())

	require.NoError(t, err)
	require.NotNil(t, users)
	require.Equal(t, int64(len(users)), numUsers)
}

// TestUpdateBirthday: tests UpdateBirthday method
//
// depends on CreateUser, NumUsers
func TestUpdateBirthday(t *testing.T) {
	arg := GenRandUser()

	user, userErr := testQueries.CreateUser(context.Background(), *arg)

	// if we get some sort of error, retrun instantly, as CreateUser does not work
	// we can indicate this with returning 1
	if userErr != nil {
		return
	}

	// creates the new birthday
	newBirthday := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	insertBirthday := UpdateBirthdayParams{
		ID:       user.ID,
		Birthday: newBirthday,
	}

	userBDAY, bdayErr := testQueries.UpdateBirthday(context.Background(), insertBirthday)

	// the inserted birthday should now not be the same
	require.NoError(t, bdayErr)
	require.NotNil(t, userBDAY)
	require.NotEqual(t, userBDAY.Birthday, user.Birthday)
}

// TestUpdateEmail: tests UpdateEmail method
//
// depends on CreateUser, NumUsers
func TestUpdateEmail(t *testing.T) {
	arg := GenRandUser()

	user, userErr := testQueries.CreateUser(context.Background(), *arg)

	// if we get some sort of error, retrun instantly, as CreateUser does not work
	// we can indicate this with returning 1
	if userErr != nil {
		return
	}

	insertEmail := UpdateEmailParams{
		ID:    user.ID,
		Email: util.RandomString(15),
	}

	userEmail, emailErr := testQueries.UpdateEmail(context.Background(), insertEmail)

	// the inserted Email should now not be the same
	require.NoError(t, emailErr)
	require.NotNil(t, userEmail)
	require.NotEqual(t, userEmail.Email, user.Email)
	require.Equal(t, userEmail.Email, insertEmail.Email)
}

// TestUpdatePassword: tests UpdatePassword method
//
// depends on CreateUser, NumUsers
func TestUpdatePassword(t *testing.T) {
	arg := GenRandUser()

	user, userErr := testQueries.CreateUser(context.Background(), *arg)

	// if we get some sort of error, retrun instantly, as CreateUser does not work
	// we can indicate this with returning 1
	if userErr != nil {
		return
	}

	insertPassword := UpdatePasswordParams{
		ID:       user.ID,
		Password: "newPassword",
	}

	userPassword, passwordErr := testQueries.UpdatePassword(context.Background(), insertPassword)

	// the inserted Password should now not be the same
	require.NoError(t, passwordErr)
	require.NotNil(t, userPassword)
	require.NotEqual(t, userPassword.Password, user.Password)
	require.Equal(t, userPassword.Password, insertPassword.Password)
}

// TestUpdatePasswordEmail: tests UpdatePasswordEmail method
//
// depends on CreateUser, NumUsers
func TestUpdatePasswordEmail(t *testing.T) {
	arg := GenRandUser()

	user, userErr := testQueries.CreateUser(context.Background(), *arg)

	// if we get some sort of error, retrun instantly, as CreateUser does not work
	// we can indicate this with returning 1
	if userErr != nil {
		return
	}

	insertEmail := UpdatePasswordEmailParams{
		Email:    user.Email,
		Password: "newPassword",
	}

	userPassword, passwordErr := testQueries.UpdatePasswordEmail(context.Background(), insertEmail)

	// the inserted Password should now not be the same
	require.NoError(t, passwordErr)
	require.NotNil(t, userPassword)
	require.NotEqual(t, userPassword.Password, user.Password)
	require.Equal(t, userPassword.Email, insertEmail.Email)
}

// TestUpdateUsername: tests UpdateUsername method
//
// depends on CreateUser, NumUsers
func TestUpdateUsername(t *testing.T) {
	arg := GenRandUser()

	user, userErr := testQueries.CreateUser(context.Background(), *arg)

	// if we get some sort of error, retrun instantly, as CreateUser does not work
	// we can indicate this with returning 1
	if userErr != nil {
		return
	}

	insertUsername := UpdateUsernameParams{
		ID:       user.ID,
		Username: util.RandomString(15),
	}

	userUsername, usernameErr := testQueries.UpdateUsername(context.Background(), insertUsername)

	// the inserted Username should now not be the same
	require.NoError(t, usernameErr)
	require.NotNil(t, userUsername)
	require.NotEqual(t, userUsername.Username, user.Username)
	require.Equal(t, userUsername.Username, insertUsername.Username)
}

// TestUpdateUsernameEmail: tests UpdateUsernameEmail method
//
// depends on CreateUser, NumUsers
func TestUpdateUsernameEmail(t *testing.T) {
	arg := GenRandUser()

	user, userErr := testQueries.CreateUser(context.Background(), *arg)

	// if we get some sort of error, retrun instantly, as CreateUser does not work
	// we can indicate this with returning 1
	if userErr != nil {
		return
	}

	insertUsername := UpdateUsernameEmailParams{
		Email:    user.Email,
		Username: util.RandomString(15),
	}

	userUsername, usernameErr := testQueries.UpdateUsernameEmail(context.Background(), insertUsername)

	// the inserted Username should now not be the same
	require.NoError(t, usernameErr)
	require.NotNil(t, userUsername)
	require.NotEqual(t, userUsername.Username, user.Username)
	require.Equal(t, userUsername.Username, insertUsername.Username)
}
