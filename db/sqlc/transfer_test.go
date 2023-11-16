package db

import (
	"context"
	"testing"
	"time"

	"github.com/fauzanfebrian/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer() (Transfer, error) {
	accountSource, _ := createRandomAccount()
	accountDest, _ := createRandomAccount()

	arg := CreateTransferParams{
		FromAccountID: accountSource.ID,
		ToAccountID:   accountDest.ID,
		Amount:        util.RandomMoney(),
	}

	return testStore.CreateTransfer(context.Background(), arg)
}

func TestCreateTransfer(t *testing.T) {
	transfer, err := createRandomTransfer()

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.NotEmpty(t, transfer.FromAccountID)
	require.NotEmpty(t, transfer.ToAccountID)
	require.NotEmpty(t, transfer.Amount)

	require.NotEmpty(t, transfer.ID)
	require.NotEmpty(t, transfer.CreatedAt)
}

func TestGetTransfer(t *testing.T) {
	transfer, _ := createRandomTransfer()
	resTransfer, err := testStore.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, resTransfer)

	require.Equal(t, transfer.ID, resTransfer.ID)
	require.Equal(t, transfer.FromAccountID, resTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, resTransfer.ToAccountID)
	require.WithinDuration(t, transfer.CreatedAt, resTransfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	transfer, _ := createRandomTransfer()

	argFrom := ListTransfersParams{
		FromAccountID: transfer.FromAccountID,
		Limit:         1,
		Offset:        0,
	}

	transfersFrom, err := testStore.ListTransfers(context.Background(), argFrom)

	require.NoError(t, err)
	require.Len(t, transfersFrom, 1)

	for _, transfer := range transfersFrom {
		require.NotEmpty(t, transfer)
	}

	argTo := ListTransfersParams{
		ToAccountID: transfer.ToAccountID,
		Limit:       1,
		Offset:      0,
	}

	transfersTo, err := testStore.ListTransfers(context.Background(), argTo)

	require.NoError(t, err)
	require.Len(t, transfersTo, 1)

	for _, transfer := range transfersTo {
		require.NotEmpty(t, transfer)
	}

}
