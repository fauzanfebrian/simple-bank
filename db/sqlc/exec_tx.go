package db

import (
	"context"
	"fmt"
)

func (store *SQLStore) execTx(ctx context.Context, queryFn func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx)

	if err != nil {
		return err
	}

	queryErr := queryFn(New(tx))

	if queryErr != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", queryErr, rbErr)
		}
		return queryErr
	}

	return tx.Commit(ctx)
}
