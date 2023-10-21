package db

import (
	"context"
	"database/sql"
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

func TestUpdateOnlyFullName(t *testing.T) {
	oldUser, _ := createRandomUser()

	newFullName := util.RandomOwner()
	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		FullName: sql.NullString{
			String: newFullName,
			Valid:  true,
		},
		Username: oldUser.Username,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FullName, newUser.FullName)
	require.Equal(t, newFullName, newUser.FullName)
	require.Equal(t, oldUser.Email, newUser.Email)
	require.Equal(t, oldUser.HashedPassword, newUser.HashedPassword)
}

func TestUpdateOnlyEmail(t *testing.T) {
	oldUser, _ := createRandomUser()

	newEmail := util.RandomEmail()
	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
		Username: oldUser.Username,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, newUser.Email)
	require.Equal(t, newEmail, newUser.Email)
	require.Equal(t, oldUser.FullName, newUser.FullName)
	require.Equal(t, oldUser.HashedPassword, newUser.HashedPassword)
}

func TestUpdateOnlyHashedPassword(t *testing.T) {
	oldUser, _ := createRandomUser()

	newEmail := util.RandomEmail()
	newFullName := util.RandomOwner()
	newHashedPassword, err := util.HashPassword(util.RandomString(6))

	require.NoError(t, err)

	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		HashedPassword: sql.NullString{
			String: newHashedPassword,
			Valid:  true,
		},
		FullName: sql.NullString{
			String: newFullName,
			Valid:  true,
		},
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
		Username: oldUser.Username,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, newUser.Email)
	require.Equal(t, newEmail, newUser.Email)
	require.NotEqual(t, oldUser.FullName, newUser.FullName)
	require.Equal(t, newFullName, newUser.FullName)
	require.NotEqual(t, oldUser.HashedPassword, newUser.HashedPassword)
	require.Equal(t, newHashedPassword, newUser.HashedPassword)
}

func TestUpdate(t *testing.T) {
	oldUser, _ := createRandomUser()

	newHashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	newUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		HashedPassword: sql.NullString{
			String: newHashedPassword,
			Valid:  true,
		},
		Username: oldUser.Username,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, newUser.HashedPassword)
	require.Equal(t, newHashedPassword, newUser.HashedPassword)
	require.Equal(t, oldUser.FullName, newUser.FullName)
	require.Equal(t, oldUser.Email, newUser.Email)
}
