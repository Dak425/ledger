package test

import (
	"io"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"gitlab.com/patchwell/ledger"
)

func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got '%s', wanted '%s'", got, want)
	}
}

func AssertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Header().Get("content-type") != want {
		t.Errorf("response content type is wrong, got '%s', wanted '%s'", response.Header().Get("content-type"), want)
	}
}

func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, wanted %d", got, want)
	}
}

func AssertTransactionType(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("transaction type is incorrect, got '%s', wanted '%s'", got, want)
	}
}

func AssertTransactionWallet(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("transaction has incorrect wallet ID, got '%s', wanted '%s'", got, want)
	}
}

func AssertTransactionAmount(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("transaction has incorrect amount, got %d, wanted %d", got, want)
	}
}

func AssertTransactionAggregate(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("transaction has incorrect aggregate ID, got '%s', wanted '%s", got, want)
	}
}

func AssertTransactions(t *testing.T, got, want []ledger.Transaction) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("transaction slices are different, got '%v', wanted '%v'", got, want)
	}
}

func AssertTransactinCount(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got incorrect amount of transactions, got %d, wanted %d", got, want)
	}
}

func AssertWalletBalance(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got incorrect wallet balance, got %d, wanted %d", got, want)
	}
}

func CreateTempFile(t *testing.T, data string, name string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	file, err := ioutil.TempFile("", name)
	if err != nil {
		t.Fatalf("unable to create temp file %v", err)
	}

	file.Write([]byte(data))

	cleanUp := func() {
		os.Remove(file.Name())
	}

	return file, cleanUp
}