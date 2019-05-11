package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/patchwell/ledger"
	"gitlab.com/patchwell/ledger/pkg/command"
)

type Server struct {
	book ledger.Book
	http.Handler
}

func NewServer(book ledger.Book) *Server {
	s := new(Server)

	s.book = book

	router := http.NewServeMux()
	router.HandleFunc("/transaction/credit", s.runAddCreditTransactionCommand)
	router.HandleFunc("/balance/wallet/", s.runWalletBalanceQuery)
	router.HandleFunc("/transactions/aggregate/", s.runAggregateTransactionsQuery)
	router.HandleFunc("/transactions/wallet/", s.runWalletTransactionsQuery)

	s.Handler = router

	return s
}

func (s *Server) runAddCreditTransactionCommand(w http.ResponseWriter, r *http.Request) {
	_, err := command.AddCreditTransaction(s.book, "1", 100, "2222")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) runWalletBalanceQuery(w http.ResponseWriter, r *http.Request) {
	wallet := r.URL.Path[len("/balance/wallet/"):]

	balance, err := command.WalletBalance(s.book, wallet)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, balance)
}

func (s *Server) runWalletTransactionsQuery(w http.ResponseWriter, r *http.Request) {
	wallet := r.URL.Path[len("/transactions/wallet/"):]

	transactions, err := command.WalletTransactions(s.book, wallet)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(transactions)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) runAggregateTransactionsQuery(w http.ResponseWriter, r *http.Request) {
	aggregate := r.URL.Path[len("/transactions/aggregate/"):]

	transactions, err := command.AggregateTransactions(s.book, aggregate)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(transactions)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}