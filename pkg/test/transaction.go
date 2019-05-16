package test

import (
	"reflect"
	"testing"

	ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"
)

func AssertTransaction(t *testing.T, transaction, want ledgerpb.Transaction) {
	t.Helper()
	if !reflect.DeepEqual(transaction, want) {
		t.Errorf("transactions are different, got '%v', wanted '%v'", transaction, want)
	}
}

func AssertTransactions(t *testing.T, transactions, want []ledgerpb.Transaction) {
	t.Helper()
	if !reflect.DeepEqual(transactions, want) {
		t.Errorf("transaction slices are different, got '%v', wanted '%v'", transactions, want)
	}
}

func AssertTransactionType(t *testing.T, transaction *ledgerpb.Transaction, want string) {
	t.Helper()
	tt := transaction.GetType()
	if tt != want {
		t.Errorf("transaction type is incorrect, got '%s', wanted '%s'", tt, want)
	}
}

func AssertTransactionWallet(t *testing.T, transaction *ledgerpb.Transaction, want string) {
	t.Helper()
	tw := transaction.GetWallet()
	if tw != want {
		t.Errorf("transaction has incorrect wallet ID, transaction '%s', wanted '%s'", tw, want)
	}
}

func AssertTransactionAmount(t *testing.T, transaction *ledgerpb.Transaction, want int32) {
	t.Helper()
	ta := transaction.GetAmount()
	if ta != want {
		t.Errorf("transaction has incorrect amount, got %d, wanted %d", ta, want)
	}
}

func AssertTransactionAggregate(t *testing.T, transaction *ledgerpb.Transaction, want string) {
	t.Helper()
	ta := transaction.GetAggregate()
	if ta != want {
		t.Errorf("transaction has incorrect aggregate ID, got '%s', wanted '%s", ta, want)
	}
}

func AssertTransactionCount(t *testing.T, transactions []*ledgerpb.Transaction, want int) {
	t.Helper()
	tl := len(transactions)
	if tl != want {
		t.Errorf("transaction slices have different lengths, got %d, wanted %d", tl, want)
	}
}
