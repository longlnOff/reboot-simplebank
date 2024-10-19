package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/longln/reboot-simplebank/internal/utils"
	"github.com/stretchr/testify/require"
)


func createRandomUser(t *testing.T) User {
	hashedPassword := "secret"
	// TODO: implement hash password
	arg := CreateUserParams{
		UserName: utils.RandomString(10),
		HashedPassword: hashedPassword,
		FullName: utils.RandomString(10),
		Email: utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserName, user.UserName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user.UserName)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user2.UserName, user.UserName)
	require.Equal(t, user2.HashedPassword, user.HashedPassword)
	require.Equal(t, user2.FullName, user.FullName)
	require.Equal(t, user2.Email, user.Email)
	require.WithinDuration(t, user2.PasswordChangedAt, user.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user2.CreatedAt, user.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.UserName)
	require.NoError(t, err)
	user2, err := testQueries.GetUser(context.Background(), user.UserName)
	require.Error(t, err)
	require.Empty(t, user2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}