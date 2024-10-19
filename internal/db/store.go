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