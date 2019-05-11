package ledger

import "testing"

func TestAddCreditTransaction(t *testing.T) {
	book := setupMockLogBook()
	walletID := "4"
	aggID := "1114"
	credit := 10000

	event := AddCreditTransaction(book, walletID, credit, aggID)

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
	book := setupMockLogBook()
	walletID := "4"
	aggID := "1114"
	debit := 10000

	event := AddDebitTransaction(book, walletID, debit, aggID)

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

func TestAddCashInTransaction(t *testing.T) {
	t.Run("should return a valid CashInTransactionAdded event", func(t *testing.T) {
		book := setupMockLogBook()
		wallet := "1"
		aggregate := "3333"
		credit := 1000

		event := AddCashInTransaction(book, wallet, credit, aggregate)

		if event.Wallet != wallet {
			t.Error("DebitTransactionAdded event has incorrect wallet ID")
		}

		if event.Credit != credit {
			t.Error("DebitTransactionAdded event has incorrect debit amount")
		}

		if event.Aggregate != aggregate {
			t.Error("DebitTransactionAdded event has incorrect aggregate ID")
		}
	})
}

func TestAddCashOutTransaction(t *testing.T) {
	t.Run("should return a valid CashOutTransactionAdded event", func(t *testing.T) {
		book := setupMockLogBook()
		wallet := "1"
		aggregate := "3333"
		debit := 1000

		event := AddCashOutTransaction(book, wallet, debit, aggregate)

		if event.Wallet != wallet {
			t.Error("DebitTransactionAdded event has incorrect wallet ID")
		}

		if event.Debit != debit {
			t.Error("DebitTransactionAdded event has incorrect debit amount")
		}

		if event.Aggregate != aggregate {
			t.Error("DebitTransactionAdded event has incorrect aggregate ID")
		}
	})
}

func TestWalletBalance(t *testing.T) {
	t.Run("should return the balance for a given wallet ID", func(t *testing.T) {
		book := setupMockLogBook()
		wallet := "1"
		expected := 100000

		balance, err := WalletBalance(book, wallet)

		if err != nil {
			t.Error("WalletBalance returned error: ", err)
		}

		if balance != expected {
			t.Errorf("WalletBalance returned %d, should have returned %d", balance, expected)
		}
	})
	t.Run("should return an error if given wallet ID has no transactions", func(t *testing.T) {
		book := setupMockLogBook()
		wallet := "4"

		balance, err := WalletBalance(book, wallet)

		if balance != 0 {
			t.Error("balance returned was not a zero balance")
		}

		if err == nil {
			t.Error("no error returned")
		}
	})
	t.Run("should return an error if wallet contains an invalid transaction type", func(t *testing.T) {
		book := setupMockLogBook()
		wallet := "5"

		balance, err := WalletBalance(book, wallet)

		if balance != 0 {
			t.Error("balance returned was not a zero balance")
		}

		if err == nil {
			t.Error("no error returned")
		}
	})
}

func TestNewLogBook(t *testing.T) {
	t.Run("should return a new log book with no transactions or map entries", func(t *testing.T) {
		book := NewLogBook()

		if len(book.transactions) > 0 {
			t.Error("transaction slice not empty")
		}

		if len(book.walletMap) > 0 {
			t.Error("wallet map not empty")
		}

		if len(book.aggregateMap) > 0 {
			t.Error("aggregate map not empty")
		}
	})
}

func TestLogBook_AddTransaction(t *testing.T) {
	book := setupMockLogBook()
	transactionType := transactionDebit
	walletID := "4"
	aggID := "1115"
	debit := 10000
	count := len(book.transactions)

	book.addTransaction(transactionType, walletID, debit, aggID)

	newCount := len(book.transactions)

	if newCount != count+1 {
		t.Errorf("log book transaction count before addTransaction was %d, should be %d afterwards, got %d", count, count+1, newCount)
	}

	transaction := book.transactions[len(book.transactions)-1]

	if *book.walletMap[walletID][0] != transaction {
		t.Error("log book wallet transaction map should contain the new transaction")
	}

	if *book.aggregateMap[aggID][0] != transaction {
		t.Error("log book aggregate transaction map should container the new transaction")
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

func TestLogBook_WalletTransactions(t *testing.T) {
	t.Run("should return all transactions for a given wallet ID", func(t *testing.T) {
		book := setupMockLogBook()

		wallet := "1"

		transactions, err := book.walletTransactions(wallet)

		if err != nil {
			t.Error("walletTransactions returned error: ", err)
		}

		count := len(transactions)

		if count != 7 {
			t.Errorf("mock log book contains 7 entries for wallet (%s), walletTransactions returned %d", wallet, count)
		}
	})
	t.Run("should return an error if the given wallet ID has no transactions", func(t *testing.T) {
		book := setupMockLogBook()

		wallet := "11"

		transactions, err := book.walletTransactions(wallet)

		if transactions != nil {
			t.Error("non-nil value returned for transactions on error")
		}

		if err == nil {
			t.Error("no error returned")
		}
	})
}

func TestLogBook_AggregateTransactions(t *testing.T) {
	t.Run("should return all transactions for a given aggregate ID", func(t *testing.T) {
		book := setupMockLogBook()

		aggregate := "1112"

		transactions, err := book.aggregateTransactions(aggregate)

		if err != nil {
			t.Error("aggregateTransactions returned error: ", err)
		}

		count := len(transactions)

		if count != 3 {
			t.Errorf("mock log book contains 3 entries for aggregate (%s), aggregateTransactions returned %d", aggregate, count)
		}
	})
	t.Run("should return an error if no transactions are present for the given aggregate ID", func(t *testing.T) {
		book := setupMockLogBook()

		aggregate := "9999"

		transactions, err := book.aggregateTransactions(aggregate)

		if transactions != nil {
			t.Error("non-nil value returned for transactions")
		}

		if err == nil {
			t.Error("no error returned")
		}
	})
}

func setupMockLogBook() LogBook {
	t1 := Transaction{transactionCashIn, "1", 100000, "1111"}
	t2 := Transaction{transactionCashIn, "1", 10000, "1112"}
	t3 := Transaction{transactionDebit, "1", 10000, "1112"}
	t4 := Transaction{transactionCredit, "2", 10000, "1112"}
	t5 := Transaction{transactionDebit, "2", 1000, "1113"}
	t6 := Transaction{transactionCredit, "1", 1000, "1113"}
	t7 := Transaction{transactionDebit, "1", 1000, "1113"}
	t8 := Transaction{transactionCredit, "3", 1000, "1113"}
	t9 := Transaction{transactionDebit, "2", 1000, "1114"}
	t10 := Transaction{transactionCredit, "1", 1000, "1114"}
	t11 := Transaction{transactionCashOut, "1", 1000, "1114"}
	t12 := Transaction{"invalid", "5", 1000, "2222"}

	transactions := []Transaction{t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12}

	l := LogBook{transactions, make(map[string][]*Transaction), make(map[string][]*Transaction)}

	for _, t := range l.transactions {
		l.addWalletMapEntry(t)
		l.addAggregateMapEntry(t)
	}

	return l
}
