package file

import (
	"encoding/json"
	"fmt"
	"io"

	ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"

	"gitlab.com/patchwell/ledger"
)

type Book struct {
	database     io.ReadWriteSeeker
	transactions []ledgerpb.Transaction
	walletMap    map[string][]*ledgerpb.Transaction
	aggregateMap map[string][]*ledgerpb.Transaction
}

func NewFileSystemBook(database io.ReadWriteSeeker) *Book {
	return &Book{
		database: database,
	}
}

func (b *Book) WalletBalance(wallet string) (int32, error) {
	var balance int32

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

func (b *Book) WalletTransactions(wallet string) ([]*ledgerpb.Transaction, error) {
	var walletTransactions []*ledgerpb.Transaction

	t, err := b.loadTransactions()

	if err != nil {
		return nil, err
	}

	for _, v := range t {
		if v.Wallet == wallet {
			walletTransactions = append(walletTransactions, &v)
		}
	}

	return walletTransactions, err
}

func (b *Book) loadTransactions() ([]ledgerpb.Transaction, error) {
	var transactions []ledgerpb.Transaction

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
