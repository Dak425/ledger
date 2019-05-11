package command

import (
	"testing"

	"gitlab.com/patchwell/ledger/pkg/memory"
)

const (
	invalidWallet            = "-99"
	invalidWalletTransaction = "-999"
	invalidAggregate         = "-99"
)

func TestAddCreditTransaction(t *testing.T) {
	book := memory.NewMockInMemoryBook()
	walletID := "4"
	aggID := "1114"
	credit := 10000

	event, err := AddCreditTransaction(book, walletID, credit, aggID)

	if err != nil {
		t.Errorf("error returned %v", err)
	}

	if event.Wallet != walletID {
		t.Error("CreditTransactionAdded event has incorrect destination ID")
	}

	if event.Credit != credit {
		t.Error("CreditTransactionAdded event has incorrect credit amount")
	}

	if event.Aggregate != aggID {
		t.Error("CreditTransactionAdded event has incorrect aggregate ID")
	}
}

func TestAddDebitTransaction(t *testing.T) {
	book := memory.NewMockInMemoryBook()
	walletID := "4"
	aggID := "1114"
	debit := 10000

	event, err := AddDebitTransaction(book, walletID, debit, aggID)

	if err != nil {
		t.Errorf("error returned %v", err)
	}

	if event.Wallet != walletID {
		t.Error("DebitTransactionAdded event has incorrect wallet ID")
	}

	if event.Debit != debit {
		t.Error("DebitTransactionAdded event has incorrect debit amount")
	}

	if event.Aggregate != aggID {
		t.Error("DebitTransactionAdded event has incorrect aggregate ID")
	}
}

func TestAddCashInTransaction(t *testing.T) {
	t.Run("should return a valid CashInTransactionAdded event", func(t *testing.T) {
		book := memory.NewMockInMemoryBook()
		wallet := "1"
		aggregate := "3333"
		credit := 1000

		event, err := AddCashInTransaction(book, wallet, credit, aggregate)

		if err != nil {
			t.Errorf("error returned %v", err)
		}

		if event.Wallet != wallet {
			t.Error("DebitTransactionAdded event has incorrect wallet ID")
		}

		if event.Credit != credit {
			t.Error("DebitTransactionAdded event has incorrect debit amount")
		}

		if event.Aggregate != aggregate {
			t.Error("DebitTransactionAdded event has incorrect aggregate ID")
		}
	})
}

func TestAddCashOutTransaction(t *testing.T) {
	t.Run("should return a valid CashOutTransactionAdded event", func(t *testing.T) {
		book := memory.NewMockInMemoryBook()
		wallet := "1"
		aggregate := "3333"
		debit := 1000

		event, err := AddCashOutTransaction(book, wallet, debit, aggregate)

		if err != nil {
			t.Errorf("error returned %v", err)
		}

		if event.Wallet != wallet {
			t.Error("DebitTransactionAdded event has incorrect wallet ID")
		}

		if event.Debit != debit {
			t.Error("DebitTransactionAdded event has incorrect debit amount")
		}

		if event.Aggregate != aggregate {
			t.Error("DebitTransactionAdded event has incorrect aggregate ID")
		}
	})
}

func TestWalletBalance(t *testing.T) {
	t.Run("should return the balance for a given wallet ID", func(t *testing.T) {
		book := memory.NewMockInMemoryBook()
		wallet := "1"
		expected := 100000

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
