package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDb)

	baseFromAccount, _ := createRandomAccount()
	baseToAccount, _ := createRandomAccount()
	fmt.Println(">> base balance:", baseFromAccount.Balance, baseToAccount.Balance)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTXParams{
				FromAccountID: baseFromAccount.ID,
				ToAccountID:   baseToAccount.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, baseFromAccount.ID, transfer.FromAccountID)
		require.Equal(t, baseToAccount.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotEmpty(t, transfer.ID)
		require.NotEmpty(t, transfer.CreatedAt)

		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, baseFromAccount.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotEmpty(t, fromEntry.ID)
		require.NotEmpty(t, fromEntry.CreatedAt)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, baseToAccount.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotEmpty(t, toEntry.ID)
		require.NotEmpty(t, toEntry.CreatedAt)

		// check account
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, baseFromAccount.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, baseToAccount.ID, toAccount.ID)

		fmt.Println(">> each tx:", fromAccount.Balance, toAccount.Balance)

		diff1 := baseFromAccount.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - baseToAccount.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updatedFromAccount, err := testQueries.GetAccount(context.Background(), baseFromAccount.ID)
	require.NoError(t, err)

	updatedToAccount, err := testQueries.GetAccount(context.Background(), baseToAccount.ID)
	require.NoError(t, err)

	fmt.Println(">> after updated balance:", updatedFromAccount.Balance, updatedToAccount.Balance)
	require.Equal(t, baseFromAccount.Balance-int64(n)*amount, updatedFromAccount.Balance)
	require.Equal(t, baseToAccount.Balance+int64(n)*amount, updatedToAccount.Balance)
}
