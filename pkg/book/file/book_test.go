package file

import (
	"gitlab.com/patchwell/ledger"
	"testing"

	ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"
	"gitlab.com/patchwell/ledger/pkg/test"
)

func TestBook_AddTransaction(t *testing.T) {
	data := `[
		{"type": "credit", "wallet": "1", "amount": 100000, "aggregate": "1111"},
		{"type": "debit", "wallet": "1", "amount": 50000, "aggregate": "1112"}]`

	t.Run("should add new transaction to slice and write to file", func(t *testing.T) {
		database, clean := test.CreateTempFile(t, data, "db.json")
		defer clean()

		book, err := NewFileSystemBook(database)
		if err != nil {
			t.Errorf("error returned when creating file system book, %v", err)
		}

		err = book.AddTransaction(ledger.TransactionCredit, "2", 50000, "1112")
		if err != nil {
			t.Errorf("error returned from adding transaction, %v", err)
		}

		newBook, err := NewFileSystemBook(database)
		if err != nil {
			t.Errorf("error when reloading file, %v", err)
		}

		transactions := newBook.Transactions()

		want := []ledgerpb.Transaction{
			{Type: "credit", Wallet: "1", Amount: 100000, Aggregate: "1111"},
			{Type: "debit", Wallet: "1", Amount: 50000, Aggregate: "1112"},
			{Type: "credit", Wallet: "2", Amount: 50000, Aggregate: "1112"},
		}

		test.AssertTransactions(t, transactions, want)
	})
}

func TestBook_WalletTransactions(t *testing.T) {
	data := `[
		{"type": "credit", "wallet": "1", "amount": 100000, "aggregate": "1111"},
		{"type": "debit", "wallet": "1", "amount": 50000, "aggregate": "1112"}]`

	t.Run("should return all transactions for a given wallet ID", func(t *testing.T) {
		database, clean := test.CreateTempFile(t, data, "db.json")
		defer clean()

		book, err := NewFileSystemBook(database)
		if err != nil {
			t.Errorf("error returned when creating file system book, %v", err)
		}

		transactions, err := book.WalletTransactions("1")
		if err != nil {
			t.Errorf("returned error %v", err)
		}

		want := []ledgerpb.Transaction{
			{Type: "credit", Wallet: "1", Amount: 100000, Aggregate: "1111"},
			{Type: "debit", Wallet: "1", Amount: 50000, Aggregate: "1112"},
		}

		test.AssertTransactionCount(t, transactions, len(want))
	})
}

func TestBook_WalletBalance(t *testing.T) {
	data := `[
		{"type": "credit", "wallet": "1", "amount": 100000, "aggregate": "1111"},
		{"type": "debit", "wallet": "1", "amount": 50000, "aggregate": "1112"}]`

	t.Run("should return the balance of the given wallet", func(t *testing.T) {
		database, clean := test.CreateTempFile(t, data, "db.json")
		defer clean()

		book, err := NewFileSystemBook(database)
		if err != nil {
			t.Errorf("error returned when created file system book, %v", err)
		}

		balance, err := book.WalletBalance("1")
		if err != nil {
			t.Errorf("returned error, %v", err)
		}

		want := int32(50000)

		test.AssertWalletBalance(t, balance, want)
	})
}

func TestBook_AggregateTransactions(t *testing.T) {
	data := `[
		{"type": "credit", "wallet": "1", "amount": 100000, "aggregate": "1111"},
		{"type": "debit", "wallet": "1", "amount": 50000, "aggregate": "1112"},
		{"type": "credit", "wallet": "2", "amount": 50000, "aggregate": "1112"}]`

	t.Run("should return all transactions for the given aggregate ID", func(t *testing.T) {
		database, clean := test.CreateTempFile(t, data, "db.json")
		defer clean()

		book, err := NewFileSystemBook(database)
		if err != nil {
			t.Errorf("error returned when created file system book, %v", err)
		}

		ts, err := book.AggregateTransactions("1112")
		if err != nil {
			t.Errorf("error returned, %v", err)
		}

		want := []*ledgerpb.Transaction{
			{Type: "debit", Wallet: "1", Amount: 50000, Aggregate: "1112"},
			{Type: "credit", Wallet: "2", Amount: 50000, Aggregate: "1112"},
		}

		test.AssertTransactionPointers(t, ts, want)
	})
}
