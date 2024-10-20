package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
)


func TestTransferTx(t *testing.T) {
	user1 := createRandomUser(t)
	account1 := createRandomAccount(t, user1)

	user2 := createRandomUser(t)
	account2 := createRandomAccount(t, user2)

	// Create 10 concurrent transfers:
	// 5 concurrent transfers from account 1 to account 2
	// 5 concurrent transfers from account 2 to account 1
	results := make(chan TransferTxResults)
	errs := make(chan error)
	for i := 0; i < 10; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID
		if i % 2 == 0 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}
		// Perform transfer money
		go func() {
			result, err := testStore.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID: toAccountID,
				Amount: 10,
			})

			errs <- err
			results <- result

		}()
	}

	// Check result
	for i := 0; i < 10; i++ {
		err := <- errs
		require.NoError(t, err)
		result := <- results
		require.NotEmpty(t, result)
	}

	updatedAccount1, _ := testStore.GetAccount(context.Background(), account1.ID)

	updatedAccount2, _ := testStore.GetAccount(context.Background(), account2.ID)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}