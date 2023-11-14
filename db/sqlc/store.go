package db

import (
	"context"
	"database/sql"
	"fmt"
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
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, queryFn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	queryErr := queryFn(New(tx))

	if queryErr != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", queryErr, rbErr)
		}
		return queryErr
	}

	return tx.Commit()
}
