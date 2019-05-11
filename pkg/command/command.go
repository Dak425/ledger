package command

import (
	"errors"

	"gitlab.com/patchwell/ledger"
)

// AddCreditTransaction adds a new credit transaction in the given Book
// returns a event representing the newly added transaction called CreditTransactionAdded
func AddCreditTransaction(book ledger.Book, wallet string, credit int, aggregate string) (*ledger.CreditTransactionAdded, error) {
	err := book.AddTransaction(ledger.TransactionCredit, wallet, credit, aggregate)

	if err != nil {
		return nil, err
	}

	return &ledger.CreditTransactionAdded{Wallet: wallet, Credit: credit, Aggregate: aggregate}, nil
}

// AddDebitTransaction adds a new debit type transaction in the given Book
// returns an event representing the newly added transaction called DebitTransactionAdded
func AddDebitTransaction(book ledger.Book, wallet string, debit int, aggregate string) (*ledger.DebitTransactionAdded, error) {
	err := book.AddTransaction(ledger.TransactionDebit, wallet, debit, aggregate)

	if err != nil {
		return nil, err
	}

	return &ledger.DebitTransactionAdded{Wallet: wallet, Debit: debit, Aggregate: aggregate}, nil
}

// AddCashInTransaction adds a new cash in type transaction in the given Book
// returns an event representing the newly added transaction called CashInTransactionAdded
func AddCashInTransaction(book ledger.Book, wallet string, credit int, aggregate string) (*ledger.CashInTransactionAdded, error) {
	err := book.AddTransaction(ledger.TransactionCashIn, wallet, credit, aggregate)

	if err != nil {
		return nil, err
	}

	return &ledger.CashInTransactionAdded{Wallet: wallet, Credit: credit, Aggregate: aggregate}, nil
}

// AddCashOutTransaction adds a new cash out type transaction in the given Book
// returns an event representing the newly added transaction called CashOutTransactionAdded
func AddCashOutTransaction(book ledger.Book, wallet string, debit int, aggregate string) (*ledger.CashOutTransactionAdded, error) {
	err := book.AddTransaction(ledger.TransactionCashOut, wallet, debit, aggregate)

	if err != nil {
		return nil, err
	}

	return &ledger.CashOutTransactionAdded{Wallet: wallet, Debit: debit, Aggregate: aggregate}, nil
}

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
