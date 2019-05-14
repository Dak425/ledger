package query

import (
	"testing"

	"gitlab.com/patchwell/ledger/pkg/memory"
)

const (
	invalidWallet            = "-99"
	invalidWalletTransaction = "-999"
	invalidAggregate         = "-99"
)

func TestWalletBalance(t *testing.T) {
	t.Run("should return the balance for a given wallet ID", func(t *testing.T) {
		book := memory.NewMockInMemoryBook()
		wallet := "1"
		expected := int32(100000)

		balance, err := WalletBalance(book, wallet)

		if err != nil {
			t.Error("WalletBalance returned error: ", err)
		}

		if balance != expected {
			t.Errorf("WalletBalance returned %d, should have returned %d", balance, expected)
		}
	})
	t.Run("should return an error if given wallet ID has no transactions", func(t *testing.T) {
		book := memory.NewMockInMemoryBook()
		wallet := invalidWallet

		balance, err := WalletBalance(book, wallet)

		if balance != 0 {
			t.Error("balance returned was not a zero balance")
		}

		if err == nil {
			t.Error("no error returned")
		}
	})
	t.Run("should return an error if wallet contains an invalid transaction type", func(t *testing.T) {
		book := memory.NewMockInMemoryBook()
		wallet := invalidWalletTransaction

		balance, err := WalletBalance(book, wallet)

		if balance != 0 {
			t.Error("balance returned was not a zero balance")
		}

		if err == nil {
			t.Error("no error returned")
		}
	})
}

func TestWalletTransactions(t *testing.T) {
	t.Run("returns all transactions for a given wallet", func(t *testing.T) {
		book := memory.NewMockInMemoryBook()
		wallet := "1"

		transactions, err := WalletTransactions(book, wallet)

		if err != nil {
			t.Errorf("returned an err %v", err)
		}

		assertTransactionCount(t, len(transactions), 7)
	})
	t.Run("returns an error if wallet has no transactions", func(t *testing.T) {
		book := memory.NewMockInMemoryBook()
		wallet := invalidWallet

		transactions, err := WalletTransactions(book, wallet)

		if transactions != nil {
			t.Error("transactions was not nil")
		}

		if err == nil {
			t.Error("no error returned")
		}
	})
}

func TestAggregateTransactions(t *testing.T) {
	t.Run("returns all transactions for a given aggregate", func(t *testing.T) {
		book := memory.NewMockInMemoryBook()
		aggregate := "1111"

		transactions, err := AggregateTransactions(book, aggregate)

		if err != nil {
			t.Errorf("returned an err %v", err)
		}

		assertTransactionCount(t, len(transactions), 1)
	})
	t.Run("returns an error if aggregate has no transactions", func(t *testing.T) {
		book := memory.NewMockInMemoryBook()
		aggregate := invalidAggregate

		transactions, err := AggregateTransactions(book, aggregate)

		if transactions != nil {
			t.Error("transactions was not nil")
		}

		if err == nil {
			t.Error("no error returned")
		}
	})
}

func assertTransactionCount(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got wrong transaction count, got %d, wanted %d", got, want)
	}
}
