package query

import (
	"errors"

	"gitlab.com/patchwell/ledger"
)

// WalletBalance returns the current balance of a wallet based on its transactions
// returns an error if wallet has no transactions
func WalletBalance(book ledger.Book, wallet string) (int, error) {
	ts, err := book.WalletTransactions(wallet)

	if err != nil {
		return 0, err
	}

	b := 0

	for _, t := range ts {
		switch t.Type {
		case ledger.TransactionCredit:
			b += t.Amount
		case ledger.TransactionDebit:
			b -= t.Amount
		case ledger.TransactionCashIn:
			b += t.Amount
		case ledger.TransactionCashOut:
			b -= t.Amount
		default:
			return 0, errors.New("invalid transaction type: " + t.Type)
		}
	}

	return b, nil
}

func WalletTransactions(book ledger.Book, wallet string) ([]*ledger.Transaction, error) {
	t, err := book.WalletTransactions(wallet)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func AggregateTransactions(book ledger.Book, aggregate string) ([]*ledger.Transaction, error) {
	t, err := book.AggregateTransactions(aggregate)

	if err != nil {
		return nil, err
	}

	return t, nil
}
