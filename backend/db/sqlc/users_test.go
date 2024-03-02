package db

import (
	"context"
	"testing"

	util "github.com/emoral435/swole-goal/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Email:    util.RandomEmail(10),
		Password: util.RandomPassword(15),
		Username: util.RandomUsername(10),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	// if query failed
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// checking the new user has the correct values in the table
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Username, user.Username)

	// ensure that the table is serializing the id's correctly
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}
