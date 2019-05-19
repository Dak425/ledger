package ledger

import ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"

const (
	TransactionCredit  = "credit"
	TransactionDebit   = "debit"
	TransactionCashIn  = "cash in"
	TransactionCashOut = "cash out"
)

type Book interface {
	TransferWalletFunds(source string, destination string, amount int32) (string, error)
	DepositWalletFunds(wallet string, deposit int32) (string, error)
	WithdrawWalletFunds(wallet string, withdraw int32) (string, error)
	AddTransaction(transactionType string, wallet string, amount int32, aggregate string) error
	Transactions() []ledgerpb.Transaction
	WalletBalance(wallet string) (int32, error)
	WalletTransactions(wallet string) ([]*ledgerpb.Transaction, error)
	AggregateTransactions(aggregate string) ([]*ledgerpb.Transaction, error)
}
