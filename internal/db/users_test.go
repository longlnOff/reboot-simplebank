package db

import (
	"context"
	"testing"

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