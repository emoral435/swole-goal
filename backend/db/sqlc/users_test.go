package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Email:    "gargyle.example@example.com",
		Password: "password",
		Username: "gargyle_rocks",
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
