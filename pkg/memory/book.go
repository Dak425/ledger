package memory

import (
	"errors"

	"gitlab.com/patchwell/ledger"
)

// Book represents a collection of transactions showing credit/debit actions against a given wallet
type Book struct {
	transactions []ledger.Transaction             // the collection of transactions in the book
	walletMap    map[string][]*ledger.Transaction // bookmarks for each wallet pointing to all transactions for that wallet
	aggregateMap map[string][]*ledger.Transaction // bookmarks for each aggregate pointing to all transactions for that aggregate
}

// NewInMemoryBook returns a new Book with no existing transactions
func NewInMemoryBook() *Book {
	return &Book{
		transactions: []ledger.Transaction{},
		walletMap:    make(map[string][]*ledger.Transaction),
		aggregateMap: make(map[string][]*ledger.Transaction),
	}
}

// NewMockInMemoryBook returns a new Book pre-populated with transactions
// used for testing
func NewMockInMemoryBook() *Book {
	t1 := ledger.Transaction{Type: ledger.TransactionCashIn, Wallet: "1", Amount: 100000, Aggregate: "1111"}
	t2 := ledger.Transaction{Type: ledger.TransactionCashIn, Wallet: "1", Amount: 10000, Aggregate: "1112"}
	t3 := ledger.Transaction{Type: ledger.TransactionDebit, Wallet: "1", Amount: 10000, Aggregate: "1112"}
	t4 := ledger.Transaction{Type: ledger.TransactionCredit, Wallet: "2", Amount: 10000, Aggregate: "1112"}
	t5 := ledger.Transaction{Type: ledger.TransactionDebit, Wallet: "2", Amount: 1000, Aggregate: "1113"}
	t6 := ledger.Transaction{Type: ledger.TransactionCredit, Wallet: "1", Amount: 1000, Aggregate: "1113"}
	t7 := ledger.Transaction{Type: ledger.TransactionDebit, Wallet: "1", Amount: 1000, Aggregate: "1113"}
	t8 := ledger.Transaction{Type: ledger.TransactionCredit, Wallet: "3", Amount: 1000, Aggregate: "1113"}
	t9 := ledger.Transaction{Type: ledger.TransactionDebit, Wallet: "2", Amount: 1000, Aggregate: "1114"}
	t10 := ledger.Transaction{Type: ledger.TransactionCredit, Wallet: "1", Amount: 1000, Aggregate: "1114"}
	t11 := ledger.Transaction{Type: ledger.TransactionCashOut, Wallet: "1", Amount: 1000, Aggregate: "1114"}
	t12 := ledger.Transaction{Type: "invalid", Wallet: "-999", Amount: 1000, Aggregate: "2222"}

	transactions := []ledger.Transaction{t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12}

	l := &Book{transactions, make(map[string][]*ledger.Transaction), make(map[string][]*ledger.Transaction)}

	for _, t := range l.transactions {
		l.addWalletMapEntry(t)
		l.addAggregateMapEntry(t)
	}

	return l
}

func (b *Book) AddTransaction(transactionType string, wallet string, amount int, aggregate string) error {
	// Create transaction
	t := ledger.Transaction{Type: transactionType, Wallet: wallet, Amount: amount, Aggregate: aggregate}

	// Append to master slice of transactions
	b.transactions = append(b.transactions, t)

	b.addWalletMapEntry(t)
	b.addAggregateMapEntry(t)

	return nil
}

func (b *Book) WalletTransactions(wallet string) ([]*ledger.Transaction, error) {
	if t, ok := b.walletMap[wallet]; ok {
		return t, nil
	} else {
		return nil, errors.New("no transactions for wallet (" + wallet + ")")
	}
}

func (b *Book) AggregateTransactions(aggregate string) ([]*ledger.Transaction, error) {
	if t, ok := b.aggregateMap[aggregate]; ok {
		return t, nil
	} else {
		return nil, errors.New("no transactions for aggregate (" + aggregate + ")")
	}
}

func (b *Book) Transactions() []ledger.Transaction {
	return b.transactions
}

func (b *Book) addWalletMapEntry(transaction ledger.Transaction) {
	// create wallet map entry if necessary
	if _, ok := b.walletMap[transaction.Wallet]; !ok {
		b.walletMap[transaction.Wallet] = []*ledger.Transaction{}
	}

	// append transaction pointer to wallet's slice of transactions
	b.walletMap[transaction.Wallet] = append(b.walletMap[transaction.Wallet], &transaction)
}

func (b *Book) addAggregateMapEntry(transaction ledger.Transaction) {
	// create aggregate map entry if necessary
	if _, ok := b.aggregateMap[transaction.Aggregate]; !ok {
		b.aggregateMap[transaction.Aggregate] = []*ledger.Transaction{}
	}

	// append transaction pointer to aggregate's slice of transactions
	b.aggregateMap[transaction.Aggregate] = append(b.aggregateMap[transaction.Aggregate], &transaction)
}
