package memory

import (
	"gitlab.com/patchwell/ledger/pkg/test"
	"testing"

	"gitlab.com/patchwell/ledger"
)

func TestNewInMemoryBook(t *testing.T) {
	t.Run("should return a new log book with no transactions or map entries", func(t *testing.T) {
		book := NewInMemoryBook()

		if len(book.transactions) > 0 {
			t.Error("transaction slice not empty")
		}

		if len(book.walletMap) > 0 {
			t.Error("wallet map not empty")
		}

		if len(book.aggregateMap) > 0 {
			t.Error("aggregate map not empty")
		}
	})
}

func TestBook_AddTransaction(t *testing.T) {
	book := NewMockInMemoryBook()
	transactionType := ledger.TransactionDebit
	walletID := "4"
	aggID := "1115"
	debit := int32(10000)
	count := len(book.transactions)

	book.AddTransaction(transactionType, walletID, debit, aggID)

	newCount := len(book.transactions)

	if newCount != count+1 {
		t.Errorf("log book transaction count before addTransaction was %d, should be %d afterwards, got %d", count, count+1, newCount)
	}

	transaction := book.transactions[len(book.transactions)-1]

	test.AssertTransaction(t, *book.walletMap[walletID][0], transaction)
	test.AssertTransaction(t, *book.aggregateMap[aggID][0], transaction)

	if transaction.Type != transactionType {
		t.Errorf("new transaction has type of %s, should be %s", transaction.Type, transactionType)
	}

	if transaction.Wallet != walletID {
		t.Errorf("new transaction has wallet reference of %s, should be %s", transaction.Wallet, walletID)
	}

	if transaction.Amount != debit {
		t.Errorf("new transaction has amount of %d, should be %d", transaction.Amount, debit)
	}

	if transaction.Aggregate != aggID {
		t.Errorf("new transaction has aggregate of %s, should be %s", transaction.Aggregate, aggID)
	}
}

func TestBook_WalletTransactions(t *testing.T) {
	t.Run("should return all transactions for a given wallet ID", func(t *testing.T) {
		book := NewMockInMemoryBook()

		wallet := "1"

		transactions, err := book.WalletTransactions(wallet)

		if err != nil {
			t.Error("walletTransactions returned error: ", err)
		}

		count := len(transactions)

		if count != 7 {
			t.Errorf("mock log book contains 7 entries for wallet (%s), walletTransactions returned %d", wallet, count)
		}
	})
	t.Run("should return an error if the given wallet ID has no transactions", func(t *testing.T) {
		book := NewMockInMemoryBook()

		wallet := "11"

		transactions, err := book.WalletTransactions(wallet)

		if transactions != nil {
			t.Error("non-nil value returned for transactions on error")
		}

		if err == nil {
			t.Error("no error returned")
		}
	})
}

func TestBook_AggregateTransactions(t *testing.T) {
	t.Run("should return all transactions for a given aggregate ID", func(t *testing.T) {
		book := NewMockInMemoryBook()

		aggregate := "1112"

		transactions, err := book.AggregateTransactions(aggregate)

		if err != nil {
			t.Error("aggregateTransactions returned error: ", err)
		}

		count := len(transactions)

		if count != 3 {
			t.Errorf("mock log book contains 3 entries for aggregate (%s), aggregateTransactions returned %d", aggregate, count)
		}
	})
	t.Run("should return an error if no transactions are present for the given aggregate ID", func(t *testing.T) {
		book := NewMockInMemoryBook()

		aggregate := "9999"

		transactions, err := book.AggregateTransactions(aggregate)

		if transactions != nil {
			t.Error("non-nil value returned for transactions")
		}

		if err == nil {
			t.Error("no error returned")
		}
	})
}

func TestBook_Transactions(t *testing.T) {
	t.Run("returns all transactions currently in the book", func(t *testing.T) {
		book := NewMockInMemoryBook()

		transactions := book.Transactions()

		if len(transactions) != len(book.transactions) {
			t.Errorf("did not get the correct number of transactions, got %d, wanted %d", len(transactions), len(book.transactions))
		}
	})
}
