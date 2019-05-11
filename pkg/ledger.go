package ledger

import "errors"

const (
	transactionCredit  = "credit"
	transactionDebit   = "debit"
	transactionCashIn  = "cash in"
	transactionCashOut = "cash out"
)

// Transaction represents a entry in a LogBook
type Transaction struct {
	Type      string `json:"type"`      // the type of transaction, can be "credit", "debit", "cash in", or "cash out"
	Wallet    string `json:"wallet"`    // the ID of the wallet associated with this transaction
	Amount    int    `json:"amount"`    // the amount involved in the transaction
	Aggregate string `json:"aggregate"` // the ID used to identify a group of transactions
}

// LogBook represents a collection of transactions showing credit/debit actions against a given wallet
type LogBook struct {
	transactions []Transaction             // the collection of transactions in the book
	walletMap    map[string][]*Transaction // bookmarks for each wallet pointing to all transactions for that wallet
	aggregateMap map[string][]*Transaction // bookmarks for each aggregate pointing to all transactions for that aggregate
}

// AddCreditTransaction adds a new credit transaction in the given LogBook
// returns a event representing the newly added transaction called CreditTransactionAdded
func AddCreditTransaction(book LogBook, wallet string, credit int, aggregate string) *CreditTransactionAdded {
	book.addTransaction(transactionCredit, wallet, credit, aggregate)

	return &CreditTransactionAdded{wallet, credit, aggregate}
}

// AddDebitTransaction adds a new debit type transaction in the given LogBook
// returns an event representing the newly added transaction called DebitTransactionAdded
func AddDebitTransaction(book LogBook, wallet string, debit int, aggregate string) *DebitTransactionAdded {
	book.addTransaction(transactionDebit, wallet, debit, aggregate)

	return &DebitTransactionAdded{wallet, debit, aggregate}
}

// AddCashInTransaction adds a new cash in type transaction in the given LogBook
// returns an event representing the newly added transaction called CashInTransactionAdded
func AddCashInTransaction(book LogBook, wallet string, credit int, aggregate string) *CashInTransactionAdded {
	book.addTransaction(transactionCashIn, wallet, credit, aggregate)

	return &CashInTransactionAdded{wallet, credit, aggregate}
}

// AddCashOutTransaction adds a new cash out type transaction in the given LogBook
// returns an event representing the newly added transaction called CashOutTransactionAdded
func AddCashOutTransaction(book LogBook, wallet string, debit int, aggregate string) *CashOutTransactionAdded {
	book.addTransaction(transactionCashOut, wallet, debit, aggregate)

	return &CashOutTransactionAdded{wallet, debit, aggregate}
}

// WalletBalance returns the current balance of a wallet based on its transactions
// returns an error if wallet has no transactions
func WalletBalance(book LogBook, wallet string) (int, error) {
	ts, err := book.walletTransactions(wallet)

	if err != nil {
		return 0, err
	}

	b := 0

	for _, t := range ts {
		switch t.Type {
		case transactionCredit:
			b += t.Amount
		case transactionDebit:
			b -= t.Amount
		case transactionCashIn:
			b += t.Amount
		case transactionCashOut:
			b -= t.Amount
		default:
			return 0, errors.New("invalid transaction type: " + t.Type)
		}
	}

	return b, nil
}

// NewLogBook returns a new LogBook with no existing transactions
func NewLogBook() *LogBook {
	return &LogBook{
		transactions: []Transaction{},
		walletMap:    make(map[string][]*Transaction),
		aggregateMap: make(map[string][]*Transaction),
	}
}

func (l *LogBook) addTransaction(transactionType string, wallet string, amount int, aggregate string) {
	// Create transaction
	t := Transaction{transactionType, wallet, amount, aggregate}

	// Append to master slice of transactions
	l.transactions = append(l.transactions, t)

	l.addWalletMapEntry(t)
	l.addAggregateMapEntry(t)
}

func (l *LogBook) addWalletMapEntry(transaction Transaction) {
	// create wallet map entry if necessary
	if _, ok := l.walletMap[transaction.Wallet]; !ok {
		l.walletMap[transaction.Wallet] = []*Transaction{}
	}

	// append transaction pointer to wallet's slice of transactions
	l.walletMap[transaction.Wallet] = append(l.walletMap[transaction.Wallet], &transaction)
}

func (l *LogBook) addAggregateMapEntry(transaction Transaction) {
	// create aggregate map entry if necessary
	if _, ok := l.aggregateMap[transaction.Aggregate]; !ok {
		l.aggregateMap[transaction.Aggregate] = []*Transaction{}
	}

	// append transaction pointer to aggregate's slice of transactions
	l.aggregateMap[transaction.Aggregate] = append(l.aggregateMap[transaction.Aggregate], &transaction)
}

func (l *LogBook) walletTransactions(wallet string) ([]*Transaction, error) {
	if t, ok := l.walletMap[wallet]; ok {
		return t, nil
	} else {
		return nil, errors.New("no transactions for wallet (" + wallet + ")")
	}
}

func (l *LogBook) aggregateTransactions(aggregate string) ([]*Transaction, error) {
	if t, ok := l.aggregateMap[aggregate]; ok {
		return t, nil
	} else {
		return nil, errors.New("no transactions for aggregate (" + aggregate + ")")
	}
}
