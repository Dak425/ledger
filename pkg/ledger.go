package ledger

import "errors"

const (
	transactionCredit  = "credit"
	transactionDebit   = "debit"
	transactionCashIn  = "cash in"
	transactionCashOut = "cash out"
)

type Transaction struct {
	Type      string `json:"type"`
	Wallet    string `json:"wallet"`
	Amount    int    `json:"amount"`
	Aggregate string `json:"aggregate"`
}

type Ledger struct {
	transactions []Transaction
	walletMap    map[string][]*Transaction
	aggregateMap map[string][]*Transaction
}

func AddCreditTransaction(ledger Ledger, wallet string, credit int, aggregate string) (*CreditTransactionAdded, error) {
	ledger.addTransaction(transactionCredit, wallet, credit, aggregate)

	return &CreditTransactionAdded{wallet, credit, aggregate}, nil
}

func AddDebitTransaction(ledger Ledger, wallet string, debit int, aggregate string) (*DebitTransactionAdded, error) {
	ledger.addTransaction(transactionDebit, wallet, debit, aggregate)

	return &DebitTransactionAdded{wallet, debit, aggregate}, nil
}

func WalletBalance(ledger Ledger, wallet string) (int, error) {
	ts, err := ledger.walletTransactions(wallet)

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

func newLedger() *Ledger {
	return &Ledger{
		transactions: []Transaction{},
		walletMap:    make(map[string][]*Transaction),
		aggregateMap: make(map[string][]*Transaction),
	}
}

func (l *Ledger) addTransaction(transactionType string, wallet string, amount int, aggregate string) {
	// Create transaction
	t := Transaction{transactionType, wallet, amount, aggregate}

	// Append to master slice of transactions
	l.transactions = append(l.transactions, t)

	l.addWalletMapEntry(t)
	l.addAggregateMapEntry(t)
}

func (l *Ledger) addWalletMapEntry(transaction Transaction) {
	// create wallet map entry if necessary
	if _, ok := l.walletMap[transaction.Wallet]; !ok {
		l.walletMap[transaction.Wallet] = []*Transaction{}
	}

	// append transaction pointer to wallet's slice of transactions
	l.walletMap[transaction.Wallet] = append(l.walletMap[transaction.Wallet], &transaction)
}

func (l *Ledger) addAggregateMapEntry(transaction Transaction) {
	// create aggregate map entry if necessary
	if _, ok := l.aggregateMap[transaction.Aggregate]; !ok {
		l.aggregateMap[transaction.Aggregate] = []*Transaction{}
	}

	// append transaction pointer to aggregate's slice of transactions
	l.aggregateMap[transaction.Aggregate] = append(l.aggregateMap[transaction.Aggregate], &transaction)
}

func (l *Ledger) walletTransactions(wallet string) ([]*Transaction, error) {
	if t, ok := l.walletMap[wallet]; ok {
		return t, nil
	} else {
		return nil, errors.New("no transactions for wallet (" + wallet + ")")
	}
}

func (l *Ledger) aggregateTransactions(aggregate string) ([]*Transaction, error) {
	if t, ok := l.aggregateMap[aggregate]; ok {
		return t, nil
	} else {
		return nil, errors.New("no transactions for aggregate (" + aggregate + ")")
	}
}
