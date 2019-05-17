package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/patchwell/ledger"
)

const jsonContentType = "application/json"

type addCreditTransactionDTO struct {
	Wallet    string `json:"wallet"`
	Credit    int32  `json:"credit"`
	Aggregate string `json:"aggregate"`
}

type Server struct {
	book ledger.Book
	http.Handler
}

func NewServer(book ledger.Book) *Server {
	s := new(Server)

	s.book = book

	router := http.NewServeMux()

	// Commands
	router.HandleFunc("/transaction/credit", s.runAddCreditTransactionCommand)

	// Queries
	router.HandleFunc("/balance/wallet/", s.runWalletBalanceQuery)
	router.HandleFunc("/transactions/aggregate/", s.runAggregateTransactionsQuery)
	router.HandleFunc("/transactions/wallet/", s.runWalletTransactionsQuery)

	s.Handler = router

	return s
}

func (s *Server) runAddCreditTransactionCommand(w http.ResponseWriter, r *http.Request) {
	var input addCreditTransactionDTO

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = AddCreditTransaction(s.book, input.Wallet, input.Credit, input.Aggregate)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) runWalletBalanceQuery(w http.ResponseWriter, r *http.Request) {
	wallet := r.URL.Path[len("/balance/wallet/"):]

	balance, err := WalletBalance(s.book, wallet)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, balance)
}

func (s *Server) runWalletTransactionsQuery(w http.ResponseWriter, r *http.Request) {
	wallet := r.URL.Path[len("/transactions/wallet/"):]

	transactions, err := WalletTransactions(s.book, wallet)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s.respondWithJSON(w, transactions)
}

func (s *Server) runAggregateTransactionsQuery(w http.ResponseWriter, r *http.Request) {
	aggregate := r.URL.Path[len("/transactions/aggregate/"):]

	transactions, err := AggregateTransactions(s.book, aggregate)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s.respondWithJSON(w, transactions)
}

func (s *Server) respondWithJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-type", jsonContentType)

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
