package db

import (
	"context"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "Saul",
		Balance:  1000000,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
}
