package ledger

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
