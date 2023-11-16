package db

import (
	"context"
	"testing"
	"time"

	"github.com/fauzanfebrian/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry() (Entry, error) {
	account, _ := createRandomAccount()

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	return testStore.CreateEntry(context.Background(), arg)
}

func TestCreateEntry(t *testing.T) {
	entry, err := createRandomEntry()

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotEmpty(t, entry.AccountID)
	require.NotEmpty(t, entry.Amount)

	require.NotEmpty(t, entry.ID)
	require.NotEmpty(t, entry.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	entry, _ := createRandomEntry()
	resEntry, err := testStore.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, resEntry)

	require.Equal(t, entry.ID, resEntry.ID)
	require.Equal(t, entry.AccountID, resEntry.AccountID)

	require.WithinDuration(t, entry.CreatedAt, resEntry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	entry, _ := createRandomEntry()

	argFrom := ListEntriesParams{
		AccountID: entry.AccountID,
		Limit:     1,
		Offset:    0,
	}

	entrysFrom, err := testStore.ListEntries(context.Background(), argFrom)

	require.NoError(t, err)
	require.Len(t, entrysFrom, 1)

	for _, entry := range entrysFrom {
		require.NotEmpty(t, entry)
	}

}
