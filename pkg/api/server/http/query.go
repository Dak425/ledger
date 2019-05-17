package http

import (
	ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"

	"gitlab.com/patchwell/ledger"
)

// WalletBalance returns the current balance of a wallet based on its transactions
// returns an error if wallet has no transactions
func WalletBalance(book ledger.Book, wallet string) (int32, error) {
	return book.WalletBalance(wallet)
}

func WalletTransactions(book ledger.Book, wallet string) ([]*ledgerpb.Transaction, error) {
	t, err := book.WalletTransactions(wallet)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func AggregateTransactions(book ledger.Book, aggregate string) ([]*ledgerpb.Transaction, error) {
	t, err := book.AggregateTransactions(aggregate)

	if err != nil {
		return nil, err
	}

	return t, nil
}
