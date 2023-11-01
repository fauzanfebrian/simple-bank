package db

import "context"

type TransferTXParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTXParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg))
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, addMoneyParams{
				account1Id: arg.FromAccountID,
				amount1:    -arg.Amount,
				account2Id: arg.ToAccountID,
				amount2:    arg.Amount,
			})
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, addMoneyParams{
				account1Id: arg.ToAccountID,
				amount1:    arg.Amount,
				account2Id: arg.FromAccountID,
				amount2:    -arg.Amount,
			})
		}

		return err
	})

	return result, err
}

type addMoneyParams struct {
	account1Id int64
	account2Id int64
	amount1    int64
	amount2    int64
}

func addMoney(ctx context.Context, q *Queries, arg addMoneyParams) (
	account1 Account,
	account2 Account,
	err error,
) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     arg.account1Id,
		Amount: arg.amount1,
	})

	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     arg.account2Id,
		Amount: arg.amount2,
	})

	return
}
