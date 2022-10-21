package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"pro.qbitty/simplebank/util"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.PasswordChangedAt)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	user1, err := testQueries.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user1.Username, user.Username)
	require.Equal(t, user1.HashedPassword, user.HashedPassword)
	require.Equal(t, user1.FullName, user.FullName)
	require.Equal(t, user1.Email, user.Email)

	require.WithinDuration(t, user1.PasswordChangedAt, user.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user.CreatedAt, time.Second)
}
