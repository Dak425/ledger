package file

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/patchwell/ledger"
)

type Book struct {
	database io.ReadWriteSeeker
}

func (b *Book) WalletBalance(wallet string) (int, error) {
	var balance int

	t, err := b.loadTransactions()

	if err != nil {
		return balance, err
	}

	for _, v := range t {
		if v.Wallet == wallet {
			switch v.Type {
			case ledger.TransactionCredit:
				balance += v.Amount
			case ledger.TransactionDebit:
				balance -= v.Amount
			case ledger.TransactionCashIn:
				balance += v.Amount
			case ledger.TransactionCashOut:
				balance -= v.Amount
			}
		}
	}

	return balance, nil
}

func (b *Book) WalletTransactions(wallet string) ([]ledger.Transaction, error) {
	var walletTransactions []ledger.Transaction

	t, err := b.loadTransactions()

	if err != nil {
		return nil, err
	}

	for _, v := range t {
		if v.Wallet == wallet {
			walletTransactions = append(walletTransactions, v)
		}
	}

	return walletTransactions, err
}

func (b *Book) loadTransactions() ([]ledger.Transaction, error) {
	var transactions []ledger.Transaction

	_, err := b.database.Seek(0, 0)
	if err != nil {
		err = fmt.Errorf("problem seeking to load transactions, %v", err)
		return transactions, err
	}

	err = json.NewDecoder(b.database).Decode(&transactions)
	if err != nil {
		err = fmt.Errorf("problem parsing transactions, %v", err)
	}

	return transactions, err
}
