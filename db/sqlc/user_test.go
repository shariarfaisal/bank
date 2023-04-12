package db

import (
	"context"
	"testing"
	"time"

	"github.com/shariarfaisal/bank/util"
	"github.com/stretchr/testify/require"
)

func randomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
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

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	randomUser(t)
}

func TestGetUser(t *testing.T) {

	u := randomUser(t)
	user, err := testQueries.GetUser(context.Background(), u.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, u.Username, user.Username)
	require.Equal(t, u.HashedPassword, user.HashedPassword)
	require.Equal(t, u.FullName, user.FullName)
	require.Equal(t, u.Email, user.Email)
	require.WithinDuration(t, u.CreatedAt, user.CreatedAt, time.Second)
	require.WithinDuration(t, u.PasswordChangedAt, user.PasswordChangedAt, time.Second)
}
