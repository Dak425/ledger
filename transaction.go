package ledger

const (
	TransactionCredit  = "credit"
	TransactionDebit   = "debit"
	TransactionCashIn  = "cash in"
	TransactionCashOut = "cash out"
)

// Transaction represents a entry in a InMemoryBook
type Transaction struct {
	Type      string `json:"type"`      // the type of transaction, can be "credit", "debit", "cash in", or "cash out"
	Wallet    string `json:"wallet"`    // the ID of the wallet associated with this transaction
	Amount    int    `json:"amount"`    // the amount involved in the transaction
	Aggregate string `json:"aggregate"` // the ID used to identify a group of transactions
}
