package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store providea all function to execute db queries and tx
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTXParams) (TransferTxResult, error)
	CreateUserTx(ctx context.Context, arg CreateUserTXParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTXParams) (VerifyEmailTxResult, error)
}

type SQLStore struct {
	*Queries
	connPool *pgxpool.Pool
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
