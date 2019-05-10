package ledger

type CreditTransactionAdded struct {
	Wallet    string `json:"wallet"`
	Credit    int    `json:"credit"`
	Aggregate string `json:"aggregate"`
}

type DebitTransactionAdded struct {
	Wallet    string `json:"wallet"`
	Debit     int    `json:"debit"`
	Aggregate string `json:"aggregate"`
}
