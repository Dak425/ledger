package ledger

import "testing"

func TestAddCreditTransaction(t *testing.T) {
	ledger := setupMockLedger()
	walletID := "4"
	aggID := "1114"
	credit := 10000

	event, err := AddCreditTransaction(ledger, walletID, credit, aggID)

	if err != nil {
		t.Error("AddCreditTransaction returned error: ", err)
	}

	if event.Wallet != walletID {
		t.Error("CreditTransactionAdded event has incorrect destination ID")
	}

	if event.Credit != credit {
		t.Error("CreditTransactionAdded event has incorrect credit amount")
	}

	if event.Aggregate != aggID {
		t.Error("CreditTransactionAdded event has incorrect aggregate ID")
	}
}

func TestAddDebitTransaction(t *testing.T) {
	ledger := setupMockLedger()
	walletID := "4"
	aggID := "1114"
	debit := 10000

	event, err := AddDebitTransaction(ledger, walletID, debit, aggID)

	if err != nil {
		t.Error("AddDebitTransaction returned error: ", err)
	}

	if event.Wallet != walletID {
		t.Error("DebitTransactionAdded event has incorrect wallet ID")
	}

	if event.Debit != debit {
		t.Error("DebitTransactionAdded event has incorrect debit amount")
	}

	if event.Aggregate != aggID {
		t.Error("DebitTransactionAdded event has incorrect aggregate ID")
	}
}

func TestWalletBalance(t *testing.T) {
	ledger := setupMockLedger()
	wallet := "2"
	expected := 9000

	balance, err := WalletBalance(ledger, wallet)

	if err != nil {
		t.Error("WalletBalance returned error: ", err)
	}

	if balance != expected {
		t.Errorf("WalletBalance returned %d, should have returned %d", balance, expected)
	}
}

func TestLedger_AddTransaction(t *testing.T) {
	ledger := setupMockLedger()
	transactionType := transactionDebit
	walletID := "4"
	aggID := "1114"
	debit := 10000
	count := len(ledger.transactions)

	ledger.addTransaction(transactionType, walletID, debit, aggID)

	newCount := len(ledger.transactions)

	if newCount != count + 1 {
		t.Errorf("ledger count before addTransaction was %d, should be %d afterwards, got %d", count, count + 1, newCount)
	}

	transaction := ledger.transactions[len(ledger.transactions) - 1]

	if *ledger.walletMap[walletID][0] != transaction {
		t.Error("ledger wallet transaction map should contain the new transaction")
	}

	if *ledger.aggregateMap[aggID][0] != transaction {
		t.Error("ledger aggregate transaction map should container the new transaction")
	}

	if transaction.Type != transactionType {
		t.Errorf("new transaction has type of %s, should be %s", transaction.Type, transactionType)
	}

	if transaction.Wallet != walletID {
		t.Errorf("new transaction has wallet reference of %s, should be %s", transaction.Wallet, walletID)
	}

	if transaction.Amount != debit {
		t.Errorf("new transaction has amount of %d, should be %d", transaction.Amount, debit)
	}

	if transaction.Aggregate != aggID {
		t.Errorf("new transaction has aggregate of %s, should be %s", transaction.Aggregate, aggID)
	}
}

func TestLedger_WalletTransactions(t *testing.T) {
	ledger := setupMockLedger()

	wallet := "1"

	transactions, err := ledger.walletTransactions(wallet)

	if err != nil {
		t.Error("walletTransactions returned error: ", err)
	}

	count := len(transactions)

	if count != 5 {
		t.Errorf("mock ledger contains 5 entries for wallet (%s), walletTransactions returned %d", wallet, count)
	}
}

func TestLedger_AggregateTransactions(t *testing.T) {
	ledger := setupMockLedger()

	aggregate := "1112"

	transactions, err := ledger.aggregateTransactions(aggregate)

	if err != nil {
		t.Error("aggregateTransactions returned error: ", err)
	}

	count := len(transactions)

	if count != 3 {
		t.Errorf("mock ledger contains 3 entries for aggregate (%s), aggregateTransactions returned %d", aggregate, count)
	}
}

func setupMockLedger() Ledger {
	t1 := Transaction{transactionCashIn, "1", 100000, "1111"}
	t2 := Transaction{transactionCashIn, "1", 10000, "1112"}
	t3 := Transaction{transactionDebit, "1", 10000, "1112"}
	t4 := Transaction{transactionCredit, "2", 10000, "1112"}
	t5 := Transaction{transactionDebit, "2", 1000, "1113"}
	t6 := Transaction{transactionCredit, "1", 1000, "1113"}
	t7 := Transaction{transactionDebit, "1", 1000, "1113"}
	t8 := Transaction{transactionCredit, "3", 1000, "1113"}

	transactions := []Transaction{t1, t2, t3, t4, t5, t6, t7, t8}

	l := Ledger{transactions, make(map[string][]*Transaction), make(map[string][]*Transaction)}

	for _, t := range l.transactions {
		l.addWalletMapEntry(t)
		l.addAggregateMapEntry(t)
	}

	return l
}