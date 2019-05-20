package ledger

type WalletFundsTransferred struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Amount      int32  `json:"amount"`
	Aggregate   string `json:"aggregate"`
}

type WalletFundsDeposited struct {
	Wallet    string `json:"wallet"`
	Deposit   int32  `json:"deposit"`
	Aggregate string `json:"aggregate"`
}

type WalletFundsWithdrawn struct {
	Wallet    string `json:"wallet"`
	Withdraw  int32  `json:"withdraw"`
	Aggregate string `json:"aggregate"`
}

type CreditTransactionAdded struct {
	Wallet    string `json:"wallet"`
	Credit    int32  `json:"credit"`
	Aggregate string `json:"aggregate"`
}

type DebitTransactionAdded struct {
	Wallet    string `json:"wallet"`
	Debit     int32  `json:"debit"`
	Aggregate string `json:"aggregate"`
}

type CashInTransactionAdded struct {
	Wallet    string `json:"wallet"`
	Credit    int32  `json:"credit"`
	Aggregate string `json:"aggregate"`
}

type CashOutTransactionAdded struct {
	Wallet    string `json:"wallet"`
	Debit     int32  `json:"debit"`
	Aggregate string `json:"aggregate"`
}
