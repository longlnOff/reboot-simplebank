package db

import (
	"context"
	"database/sql"
	"fmt"
)

type StoreSQL struct {
	*Queries				// struct composite --> to perform all database operations
	DB *sql.DB				// to perform database transaction
}

func NewStore(db *sql.DB) *StoreSQL {
	return &StoreSQL{
		DB: db,
		Queries: New(db),
	}
}



// perform db transaction, database operations written in fn anonymous function
// 1. begin transaction --> 2. perform db operation --> 3. Rollback if error, else commit
func (store *StoreSQL) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResults struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	FromEntry Entry `json:"from_entry"`
	ToAccount Account `json:"to_account"`
	ToEntry Entry `json:"to_entry"`
}

func (store *StoreSQL) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResults, error) {
	var result TransferTxResults

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// 1. Create transfer
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg))
		if err != nil {
			return err
		}
		// 2. Create From Entry
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}
		// 3. Create To Entry
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}
		// fmt.Println("He")
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, store.Queries, -arg.Amount, arg.FromAccountID, arg.Amount, arg.ToAccountID)

		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, store.Queries, arg.Amount, arg.ToAccountID, -arg.Amount, arg.FromAccountID)

		}
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return TransferTxResults{}, err
	}


	return result, nil
}

func addMoney(
	ctx context.Context,
	q *Queries,
	Amount1 int64,
	AccountID1 int64,
	Amount2 int64,
	AccountID2 int64,
) (account1 Account, account2 Account,err error) {
	account1, err = q.AddBalance(ctx, AddBalanceParams{
		Amount: Amount1,
		ID: AccountID1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddBalance(ctx, AddBalanceParams{
		Amount: Amount2,
		ID: AccountID2,
	})
	return
}