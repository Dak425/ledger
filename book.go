package ledger

import ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"

const (
	TransactionCredit  = "credit"
	TransactionDebit   = "debit"
	TransactionCashIn  = "cash in"
	TransactionCashOut = "cash out"
)

type Book interface {
	AddTransaction(transactionType string, wallet string, amount int32, aggregate string) error
	Transactions() []ledgerpb.Transaction
	WalletBalance(wallet string) (int32, error)
	WalletTransactions(wallet string) ([]*ledgerpb.Transaction, error)
	AggregateTransactions(aggregate string) ([]*ledgerpb.Transaction, error)
}
