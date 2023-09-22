package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/fauzanfebrian/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser() (User, error) {
	username := util.RandomOwner()
	email := fmt.Sprintf("%s@simplebank.com", username)

	password, err := util.HashPassword("12345678")
	if err != nil {
		return User{}, err
	}

	arg := CreateUserParams{
		Username:       username,
		FullName:       username,
		Email:          email,
		HashedPassword: password,
	}

	return testQueries.CreateUser(context.Background(), arg)
}

func TestCreateUser(t *testing.T) {
	account, err := createRandomUser()

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.NotEmpty(t, account.Username)
	require.NotEmpty(t, account.FullName)
	require.NotEmpty(t, account.Email)
	require.NotEmpty(t, account.HashedPassword)

	require.NotZero(t, account.PasswordChangedAt)
	require.NotZero(t, account.CreatedAt)
}

func TestGetUser(t *testing.T) {
	account, _ := createRandomUser()
	resUser, err := testQueries.GetUser(context.Background(), account.Username)

	require.NoError(t, err)
	require.NotEmpty(t, resUser)

	require.Equal(t, account.Username, resUser.Username)
	require.Equal(t, account.FullName, resUser.FullName)
	require.Equal(t, account.Email, resUser.Email)
	require.WithinDuration(t, account.PasswordChangedAt, resUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, account.CreatedAt, resUser.CreatedAt, time.Second)
}
