package ledger

type Book interface {
	AddTransaction(transactionType string, wallet string, amount int, aggregate string) error
	Transactions() []Transaction
	WalletTransactions(wallet string) ([]*Transaction, error)
	AggregateTransactions(aggregate string) ([]*Transaction, error)
}
