package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"gitlab.com/patchwell/ledger"
	"gitlab.com/patchwell/ledger/pkg/memory"
)

func TestGETWalletBalance(t *testing.T) {
	book := memory.NewMockInMemoryBook()
	server := NewServer(book)

	t.Run("returns the current balance of the wallet '1'", func(t *testing.T) {
		request := newGetWalletBalanceRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "100000")
	})
	t.Run("returns the current balance of wallet '2'", func(t *testing.T) {
		request := newGetWalletBalanceRequest("2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "8000")
	})
	t.Run("returns 404 when wallet is not found", func(t *testing.T) {
		request := newGetWalletBalanceRequest("-99")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestGETWalletTransactions(t *testing.T) {
	book := memory.NewMockInMemoryBook()
	server := NewServer(book)
	wantedTransactions := []ledger.Transaction{
		{Type: ledger.TransactionCredit, Wallet: "2", Amount: 10000, Aggregate: "1112"},
		{Type: ledger.TransactionDebit, Wallet: "2", Amount: 1000, Aggregate: "1113"},
		{Type: ledger.TransactionDebit, Wallet: "2", Amount: 1000, Aggregate: "1114"},
	}

	t.Run("returns all transactions for the given wallet", func(t *testing.T) {
		request := newGetWalletTransactionsRequest("2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var transactions []ledger.Transaction

		err := json.NewDecoder(response.Body).Decode(&transactions)

		if err != nil {
			t.Errorf("unable to parse response from server '%s' into slice of Transaction, '%v'", response.Body, err)
		}

		if !reflect.DeepEqual(transactions, wantedTransactions) {
			t.Error("response did not contain the expected transactions")
		}

		assertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("returns 404 if wallet has no transactions", func(t *testing.T) {
		request := newGetWalletTransactionsRequest("99")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestGETAggregateTransactions(t *testing.T) {
	book := memory.NewMockInMemoryBook()
	server := NewServer(book)
	wantedTransactions := []ledger.Transaction{
		{Type: ledger.TransactionCashIn, Wallet: "1", Amount: 100000, Aggregate: "1111"},
	}

	t.Run("returns all transactions for the given aggregate", func(t *testing.T) {
		request := newGetAggregateTransactionsRequest("1111")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		transactions := getTransactionsFromResponse(t, response.Body)

		if !reflect.DeepEqual(transactions, wantedTransactions) {
			t.Error("response did not contain the expected transactions")
		}

		assertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("returns 404 if aggregate has no transactions", func(t *testing.T) {
		request := newGetAggregateTransactionsRequest("9999")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestPOSTCreditTransaction(t *testing.T) {
	book := memory.NewMockInMemoryBook()
	server := NewServer(book)

	t.Run("it should add an additional transaction to the ledger book", func(t *testing.T) {
		request := newPostCreditTransactionRequest("1", 100, "2222")
		response := httptest.NewRecorder()

		want := len(book.Transactions()) + 1

		server.ServeHTTP(response, request)

		got := len(book.Transactions())

		assertStatus(t, response.Code, http.StatusAccepted)

		if got != want {
			t.Errorf("ledger book has invalid transaction count, got %d, wanted %d", got, want)
		}

		transactions := book.Transactions()

		lastTransaction := transactions[len(transactions)-1]

		assertTransactionType(t, lastTransaction.Type, ledger.TransactionCredit)
		assertTransactionWallet(t, lastTransaction.Wallet, "1")
		assertTransactionAmount(t, lastTransaction.Amount, 100)
		assertTransactionAggregate(t, lastTransaction.Aggregate, "2222")
	})
}

func newGetWalletBalanceRequest(wallet string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/balance/wallet/%s", wallet), nil)
	return req
}

func newGetWalletTransactionsRequest(wallet string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/transactions/wallet/%s", wallet), nil)
	return req
}

func newGetAggregateTransactionsRequest(aggregate string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/transactions/aggregate/%s", aggregate), nil)
	return req
}

func newPostCreditTransactionRequest(wallet string, credit int, aggregate string) *http.Request {
	payload := addCreditTransactionDTO{
		Wallet:    wallet,
		Credit:    credit,
		Aggregate: aggregate,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/transaction/credit", bytes.NewBuffer(body))
	return req
}

func getTransactionsFromResponse(t *testing.T, body io.Reader) (transactions []ledger.Transaction) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&transactions)

	if err != nil {
		t.Errorf("unable to parse response from server '%s' into slice of Transaction, '%v'", body, err)
	}

	return
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got '%s', wanted '%s'", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, wanted %d", got, want)
	}
}

func assertTransactionType(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("transaction type is incorrect, got '%s', wanted '%s'", got, want)
	}
}

func assertTransactionWallet(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("transaction has incorrect wallet ID, got '%s', wanted '%s'", got, want)
	}
}

func assertTransactionAmount(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("transaction has incorrect amount, got %d, wanted %d", got, want)
	}
}

func assertTransactionAggregate(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("transaction has incorrect aggregate ID, got '%s', wanted '%s", got, want)
	}
}
