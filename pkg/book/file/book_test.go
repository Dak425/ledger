package file

import (
	"testing"

	ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"
	"gitlab.com/patchwell/ledger/pkg/test"
)

func TestBook_WalletTransactions(t *testing.T) {
	data := `[
		{"type": "credit", "wallet": "1", "amount": 100000, "aggregate": "1111"},
		{"type": "debit", "wallet": "1", "amount": 50000, "aggregate": "1112"}]`

	t.Run("should return all transactions for a given wallet ID", func(t *testing.T) {
		database, clean := test.CreateTempFile(t, data, "db")
		defer clean()

		book := Book{database}

		transactions, err := book.WalletTransactions("1")

		if err != nil {
			t.Errorf("returned error %v", err)
		}

		want := []ledgerpb.Transaction{
			{Type: "credit", Wallet: "1", Amount: 100000, Aggregate: "1111"},
			{Type: "debit", Wallet: "1", Amount: 50000, Aggregate: "1112"},
		}

		test.AssertTransactinCount(t, len(transactions), len(want))
	})
}

func TestBook_WalletBalance(t *testing.T) {
	data := `[
		{"type": "credit", "wallet": "1", "amount": 100000, "aggregate": "1111"},
		{"type": "debit", "wallet": "1", "amount": 50000, "aggregate": "1112"}]`

	t.Run("should return the balance of the given wallet", func(t *testing.T) {
		database, clean := test.CreateTempFile(t, data, "db")
		defer clean()

		book := Book{database}

		balance, err := book.WalletBalance("1")

		if err != nil {
			t.Errorf("returned error, %v", err)
		}

		want := int32(50000)

		test.AssertWalletBalance(t, balance, want)
	})
}
