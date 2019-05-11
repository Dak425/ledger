package http

import (
	"encoding/json"
	"fmt"
	"gitlab.com/patchwell/ledger"
	"gitlab.com/patchwell/ledger/pkg/memory"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TransactionResponse struct {
	Transactions []ledger.Transaction
}

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

	t.Run("returns all transactions for the given wallet", func(t *testing.T) {
		request := newGetWalletTransactionsRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var transactions []ledger.Transaction

		err := json.NewDecoder(response.Body).Decode(&transactions)

		if err != nil {
			t.Errorf("unable to parse response from server '%s' into slice of Transaction, '%v'", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func TestGETAggregateTransactions(t *testing.T) {
	book := memory.NewMockInMemoryBook()
	server := NewServer(book)

	t.Run("returns all transactions for the given aggregate", func(t *testing.T) {
		request := newGetAggregateTransactionsRequest("1112")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var transactions []ledger.Transaction

		err := json.NewDecoder(response.Body).Decode(&transactions)

		if err != nil {
			t.Errorf("unable to parse response from server '%s' into slice of Transaction, '%v'", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
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
	req, _ := http.NewRequest(http.MethodPost, "/transaction/credit", nil)
	return req
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
