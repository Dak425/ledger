package http

import (
	"gitlab.com/patchwell/ledger/pkg/book/memory"
	"testing"
)

func TestAddCreditTransaction(t *testing.T) {
	book := memory.NewMockInMemoryBook()
	walletID := "4"
	aggID := "1114"
	credit := int32(10000)

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
	debit := int32(10000)

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
		credit := int32(1000)

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
		debit := int32(1000)

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
