package db

import (
	"context"
	"testing"
	"time"

	"github.com/fauzanfebrian/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry() (Entry, error) {
	account, _ := createRandomAccount()

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	return testQueries.CreateEntry(context.Background(), arg)
}

func TestCreateEntry(t *testing.T) {
	transfer, err := createRandomEntry()

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.NotEmpty(t, transfer.AccountID)
	require.NotEmpty(t, transfer.Amount)

	require.NotEmpty(t, transfer.ID)
	require.NotEmpty(t, transfer.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	transfer, _ := createRandomEntry()
	resEntry, err := testQueries.GetEntry(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, resEntry)

	require.Equal(t, transfer.ID, resEntry.ID)
	require.Equal(t, transfer.AccountID, resEntry.AccountID)

	require.WithinDuration(t, transfer.CreatedAt, resEntry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	transfer, _ := createRandomEntry()

	argFrom := ListEntriesParams{
		AccountID: transfer.AccountID,
		Limit:     1,
		Offset:    0,
	}

	transfersFrom, err := testQueries.ListEntries(context.Background(), argFrom)

	require.NoError(t, err)
	require.Len(t, transfersFrom, 1)

	for _, transfer := range transfersFrom {
		require.NotEmpty(t, transfer)
	}

}
