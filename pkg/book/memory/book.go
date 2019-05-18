package memory

import (
	"errors"

	ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"

	"gitlab.com/patchwell/ledger"
)

// Book represents a collection of transactions showing credit/debit actions against a given wallet
type Book struct {
	transactions []ledgerpb.Transaction             // the collection of transactions in the book
	walletMap    map[string][]*ledgerpb.Transaction // bookmarks for each wallet pointing to all transactions for that wallet
	aggregateMap map[string][]*ledgerpb.Transaction // bookmarks for each aggregate pointing to all transactions for that aggregate
}

// NewInMemoryBook returns a new Book with no existing transactions
func NewInMemoryBook() *Book {
	return &Book{
		transactions: []ledgerpb.Transaction{},
		walletMap:    make(map[string][]*ledgerpb.Transaction),
		aggregateMap: make(map[string][]*ledgerpb.Transaction),
	}
}

// NewMockInMemoryBook returns a new Book pre-populated with transactions
// used for testing
func NewMockInMemoryBook() *Book {
	t1 := ledgerpb.Transaction{Type: ledger.TransactionCashIn, Wallet: "1", Amount: 100000, Aggregate: "1111"}
	t2 := ledgerpb.Transaction{Type: ledger.TransactionCashIn, Wallet: "1", Amount: 10000, Aggregate: "1112"}
	t3 := ledgerpb.Transaction{Type: ledger.TransactionDebit, Wallet: "1", Amount: 10000, Aggregate: "1112"}
	t4 := ledgerpb.Transaction{Type: ledger.TransactionCredit, Wallet: "2", Amount: 10000, Aggregate: "1112"}
	t5 := ledgerpb.Transaction{Type: ledger.TransactionDebit, Wallet: "2", Amount: 1000, Aggregate: "1113"}
	t6 := ledgerpb.Transaction{Type: ledger.TransactionCredit, Wallet: "1", Amount: 1000, Aggregate: "1113"}
	t7 := ledgerpb.Transaction{Type: ledger.TransactionDebit, Wallet: "1", Amount: 1000, Aggregate: "1113"}
	t8 := ledgerpb.Transaction{Type: ledger.TransactionCredit, Wallet: "3", Amount: 1000, Aggregate: "1113"}
	t9 := ledgerpb.Transaction{Type: ledger.TransactionDebit, Wallet: "2", Amount: 1000, Aggregate: "1114"}
	t10 := ledgerpb.Transaction{Type: ledger.TransactionCredit, Wallet: "1", Amount: 1000, Aggregate: "1114"}
	t11 := ledgerpb.Transaction{Type: ledger.TransactionCashOut, Wallet: "1", Amount: 1000, Aggregate: "1114"}
	t12 := ledgerpb.Transaction{Type: "invalid", Wallet: "-999", Amount: 1000, Aggregate: "2222"}

	transactions := []ledgerpb.Transaction{t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12}

	l := &Book{transactions, make(map[string][]*ledgerpb.Transaction), make(map[string][]*ledgerpb.Transaction)}

	for _, t := range l.transactions {
		l.addWalletMapEntry(t)
		l.addAggregateMapEntry(t)
	}

	return l
}

func (b *Book) AddTransaction(transactionType string, wallet string, amount int32, aggregate string) error {
	// Create transaction
	t := ledgerpb.Transaction{Type: transactionType, Wallet: wallet, Amount: amount, Aggregate: aggregate}

	// Append to master slice of transactions
	b.transactions = append(b.transactions, t)

	b.addWalletMapEntry(t)
	b.addAggregateMapEntry(t)

	return nil
}

func (b *Book) WalletTransactions(wallet string) ([]*ledgerpb.Transaction, error) {
	if t, ok := b.walletMap[wallet]; ok {
		return t, nil
	} else {
		return nil, errors.New("no transactions for wallet (" + wallet + ")")
	}
}

func (b *Book) WalletBalance(wallet string) (int32, error) {
	ts, err := b.WalletTransactions(wallet)

	if err != nil {
		return 0, err
	}

	balance := int32(0)

	for _, t := range ts {
		switch t.Type {
		case ledger.TransactionCredit:
			balance += t.Amount
		case ledger.TransactionDebit:
			balance -= t.Amount
		case ledger.TransactionCashIn:
			balance += t.Amount
		case ledger.TransactionCashOut:
			balance -= t.Amount
		default:
			return 0, errors.New("invalid transaction type: " + t.Type)
		}
	}

	return balance, nil
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

func (b *Book) addWalletMapEntry(transaction ledgerpb.Transaction) {
	// create wallet map entry if necessary
	if _, ok := b.walletMap[transaction.Wallet]; !ok {
		b.walletMap[transaction.Wallet] = []*ledgerpb.Transaction{}
	}

	// append transaction pointer to wallet's slice of transactions
	b.walletMap[transaction.Wallet] = append(b.walletMap[transaction.Wallet], &transaction)
}

func (b *Book) addAggregateMapEntry(transaction ledgerpb.Transaction) {
	// create aggregate map entry if necessary
	if _, ok := b.aggregateMap[transaction.Aggregate]; !ok {
		b.aggregateMap[transaction.Aggregate] = []*ledgerpb.Transaction{}
	}

	// append transaction pointer to aggregate's slice of transactions
	b.aggregateMap[transaction.Aggregate] = append(b.aggregateMap[transaction.Aggregate], &transaction)
}