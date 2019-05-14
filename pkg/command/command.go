package command

import (
	"gitlab.com/patchwell/ledger"
)

// AddCreditTransaction adds a new credit transaction in the given Book
// returns a event representing the newly added transaction called CreditTransactionAdded
func AddCreditTransaction(book ledger.Book, wallet string, credit int32, aggregate string) (*ledger.CreditTransactionAdded, error) {
	err := book.AddTransaction(ledger.TransactionCredit, wallet, credit, aggregate)

	if err != nil {
		return nil, err
	}

	return &ledger.CreditTransactionAdded{Wallet: wallet, Credit: credit, Aggregate: aggregate}, nil
}

// AddDebitTransaction adds a new debit type transaction in the given Book
// returns an event representing the newly added transaction called DebitTransactionAdded
func AddDebitTransaction(book ledger.Book, wallet string, debit int32, aggregate string) (*ledger.DebitTransactionAdded, error) {
	err := book.AddTransaction(ledger.TransactionDebit, wallet, debit, aggregate)

	if err != nil {
		return nil, err
	}

	return &ledger.DebitTransactionAdded{Wallet: wallet, Debit: debit, Aggregate: aggregate}, nil
}

// AddCashInTransaction adds a new cash in type transaction in the given Book
// returns an event representing the newly added transaction called CashInTransactionAdded
func AddCashInTransaction(book ledger.Book, wallet string, credit int32, aggregate string) (*ledger.CashInTransactionAdded, error) {
	err := book.AddTransaction(ledger.TransactionCashIn, wallet, credit, aggregate)

	if err != nil {
		return nil, err
	}

	return &ledger.CashInTransactionAdded{Wallet: wallet, Credit: credit, Aggregate: aggregate}, nil
}

// AddCashOutTransaction adds a new cash out type transaction in the given Book
// returns an event representing the newly added transaction called CashOutTransactionAdded
func AddCashOutTransaction(book ledger.Book, wallet string, debit int32, aggregate string) (*ledger.CashOutTransactionAdded, error) {
	err := book.AddTransaction(ledger.TransactionCashOut, wallet, debit, aggregate)

	if err != nil {
		return nil, err
	}

	return &ledger.CashOutTransactionAdded{Wallet: wallet, Debit: debit, Aggregate: aggregate}, nil
}
