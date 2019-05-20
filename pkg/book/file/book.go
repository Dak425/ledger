package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"

	"gitlab.com/patchwell/ledger"
)

func initializeFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("error when info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

type Book struct {
	database     *json.Encoder
	transactions []ledgerpb.Transaction
	walletMap    map[string][]*ledgerpb.Transaction
	aggregateMap map[string][]*ledgerpb.Transaction
}

func (b *Book) TransferWalletFunds(source string, destination string, amount int32) (string, error) {
	panic("implement me")
}

func (b *Book) DepositWalletFunds(wallet string, deposit int32) (string, error) {
	panic("implement me")
}

func (b *Book) WithdrawWalletFunds(wallet string, withdraw int32) (string, error) {
	panic("implement me")
}

func NewFileSystemBook(file *os.File) (*Book, error) {
	err := initializeFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initializing file for book, %v", err)
	}

	b := &Book{
		database:     json.NewEncoder(&tape{file}),
		transactions: []ledgerpb.Transaction{},
		walletMap:    make(map[string][]*ledgerpb.Transaction),
		aggregateMap: make(map[string][]*ledgerpb.Transaction),
	}

	err = b.loadTransactions(file)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Book) AddTransaction(transactionType string, wallet string, amount int32, aggregate string) error {
	t := ledgerpb.Transaction{
		Type:      transactionType,
		Wallet:    wallet,
		Amount:    amount,
		Aggregate: aggregate,
	}

	b.transactions = append(b.transactions, t)
	b.addWalletMapEntry(t)
	b.addAggregateMapEntry(t)

	b.database.Encode(b.transactions)

	return nil
}

func (b *Book) WalletBalance(wallet string) (int32, error) {
	var balance int32

	ts, err := b.WalletTransactions(wallet)

	if err != nil {
		return balance, err
	}

	for _, v := range ts {
		if v.GetWallet() == wallet {
			switch v.GetType() {
			case ledger.TransactionCredit:
				balance += v.GetAmount()
			case ledger.TransactionDebit:
				balance -= v.GetAmount()
			case ledger.TransactionCashIn:
				balance += v.GetAmount()
			case ledger.TransactionCashOut:
				balance -= v.GetAmount()
			}
		}
	}

	return balance, nil
}

func (b *Book) WalletTransactions(wallet string) ([]*ledgerpb.Transaction, error) {
	if t, ok := b.walletMap[wallet]; ok {
		return t, nil
	} else {
		return nil, errors.New("no transactions for wallet (" + wallet + ")")
	}
}

func (b *Book) AggregateTransactions(aggregate string) ([]*ledgerpb.Transaction, error) {
	if t, ok := b.aggregateMap[aggregate]; ok {
		return t, nil
	} else {
		return nil, errors.New("no transactions for aggregate (" + aggregate + ")")
	}
}

func (b *Book) Transactions() []ledgerpb.Transaction {
	return b.transactions
}

func (b *Book) loadTransactions(file *os.File) error {
	err := json.NewDecoder(file).Decode(&b.transactions)
	if err != nil {
		err = fmt.Errorf("problem parsing transactions, %v", err)
		return err
	}

	for _, t := range b.transactions {
		b.addWalletMapEntry(t)
		b.addAggregateMapEntry(t)
	}

	return nil
}

func (b *Book) addWalletMapEntry(transaction ledgerpb.Transaction) {
	// create wallet map entry if necessary
	if _, ok := b.walletMap[transaction.GetWallet()]; !ok {
		b.walletMap[transaction.GetWallet()] = []*ledgerpb.Transaction{}
	}

	// append transaction pointer to wallet's slice of transactions
	b.walletMap[transaction.GetWallet()] = append(b.walletMap[transaction.GetWallet()], &transaction)
}

func (b *Book) addAggregateMapEntry(transaction ledgerpb.Transaction) {
	// create aggregate map entry if necessary
	if _, ok := b.aggregateMap[transaction.GetAggregate()]; !ok {
		b.aggregateMap[transaction.GetAggregate()] = []*ledgerpb.Transaction{}
	}

	// append transaction pointer to aggregate's slice of transactions
	b.aggregateMap[transaction.GetAggregate()] = append(b.aggregateMap[transaction.GetAggregate()], &transaction)
}
