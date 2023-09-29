package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/fauzanfebrian/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount() (Account, error) {
	user, err := createRandomUser()
	if err != nil {
		return Account{}, err
	}

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	return testQueries.CreateAccount(context.Background(), arg)

}

func TestCreateAccount(t *testing.T) {
	account, err := createRandomAccount()

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.NotEmpty(t, account.Owner)
	require.NotEmpty(t, account.Balance)
	require.NotEmpty(t, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	account, _ := createRandomAccount()
	resAccount, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, resAccount)

	require.Equal(t, account.ID, resAccount.ID)
	require.Equal(t, account.Owner, resAccount.Owner)
	require.Equal(t, account.Balance, resAccount.Balance)
	require.Equal(t, account.Currency, resAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, resAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account, _ := createRandomAccount()

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: account.Balance,
	}

	accountUpdated, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, accountUpdated)

	require.Equal(t, account.ID, accountUpdated.ID)
	require.Equal(t, account.Owner, accountUpdated.Owner)
	require.Equal(t, account.Balance, accountUpdated.Balance)
	require.Equal(t, account.Currency, accountUpdated.Currency)
	require.WithinDuration(t, account.CreatedAt, accountUpdated.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account, _ := createRandomAccount()
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	resAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, resAccount)
}

func TestListAccounts(t *testing.T) {
	account, _ := createRandomAccount()

	arg := ListAccountsParams{
		Limit:  1,
		Offset: 0,
		Owner:  account.Owner,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 1)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
