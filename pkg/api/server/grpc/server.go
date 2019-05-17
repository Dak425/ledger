package grpc

import (
	"context"

	"gitlab.com/patchwell/ledger"
	ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"
)

type Server struct {
	book ledger.Book
}

func NewGRPCServer(book ledger.Book) *Server {
	return &Server{
		book: book,
	}
}

func (s *Server) AddCreditTransaction(ctx context.Context, req *ledgerpb.AddCreditTransactionRequest) (*ledgerpb.AddCreditTransactionResponse, error) {
	t := req.GetTransaction()
	err := s.book.AddTransaction(ledger.TransactionCredit, t.GetWallet(), t.GetCredit(), t.GetAggregate())

	if err != nil {
		return nil, err
	}

	return &ledgerpb.AddCreditTransactionResponse{
		Response: "credit transaction added successfully",
	}, nil
}

func (s *Server) AddDebitTransaction(ctx context.Context, req *ledgerpb.AddDebitTransactionRequest) (*ledgerpb.AddDebitTransactionResponse, error) {
	t := req.GetTransaction()
	err := s.book.AddTransaction(ledger.TransactionDebit, t.GetWallet(), t.GetDebit(), t.GetAggregate())

	if err != nil {
		return nil, err
	}

	return &ledgerpb.AddDebitTransactionResponse{
		Response: "debit transaction added successfully",
	}, nil
}

func (s *Server) AddCashInTransaction(ctx context.Context, req *ledgerpb.AddCashInTransactionRequest) (*ledgerpb.AddCashInTransactionResponse, error) {
	t := req.GetTransaction()
	err := s.book.AddTransaction(ledger.TransactionCashIn, t.GetWallet(), t.GetCredit(), t.GetAggregate())

	if err != nil {
		return nil, err
	}

	return &ledgerpb.AddCashInTransactionResponse{
		Response: "cash in transaction added successfully",
	}, nil
}

func (s *Server) AddCashOutTransaction(ctx context.Context, req *ledgerpb.AddCashOutTransactionRequest) (*ledgerpb.AddCashOutTransactionResponse, error) {
	t := req.GetTransaction()
	err := s.book.AddTransaction(ledger.TransactionCashOut, t.GetWallet(), t.GetDebit(), t.GetAggregate())

	if err != nil {
		return nil, err
	}

	return &ledgerpb.AddCashOutTransactionResponse{
		Response: "cash out transaction added successfully",
	}, nil
}

func (s *Server) WalletTransactions(ctx context.Context, req *ledgerpb.WalletTransactionsRequest) (*ledgerpb.WalletTransactionsResponse, error) {
	w := req.GetWallet()

	t, err := s.book.WalletTransactions(w)

	if err != nil {
		return nil, err
	}

	transactions := toProtoTransaction(t)

	return &ledgerpb.WalletTransactionsResponse{
		Transactions: transactions,
	}, nil
}

func (s *Server) WalletBalance(ctx context.Context, req *ledgerpb.WalletBalanceRequest) (*ledgerpb.WalletBalanceResponse, error) {
	w := req.GetWallet()

	b, err := s.book.WalletBalance(w)

	if err != nil {
		return nil, err
	}

	return &ledgerpb.WalletBalanceResponse{
		Balance: b,
	}, nil
}

func (s *Server) AggregateTransactions(ctx context.Context, req *ledgerpb.AggregateTransactionsRequest) (*ledgerpb.AggregateTransactionsResponse, error) {
	ts, err := s.book.AggregateTransactions(req.GetAggregate())

	if err != nil {
		return nil, err
	}

	return &ledgerpb.AggregateTransactionsResponse{
		Transactions: ts,
	}, nil
}

func toProtoTransaction(transactions []*ledgerpb.Transaction) []*ledgerpb.Transaction {
	t := make([]*ledgerpb.Transaction, len(transactions))

	for i, v := range transactions {
		t[i] = &ledgerpb.Transaction{
			Type:      v.Type,
			Wallet:    v.Wallet,
			Amount:    v.Amount,
			Aggregate: v.Aggregate,
		}
	}

	return t
}
