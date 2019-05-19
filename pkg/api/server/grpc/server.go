package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	if ctx.Err() == context.Canceled {
		return nil, status.Error(codes.Canceled, "client cancelled, aborting")
	}

	t := req.GetTransaction()
	err := s.book.AddTransaction(ledger.TransactionCredit, t.GetWallet(), t.GetCredit(), t.GetAggregate())

	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "problem when adding credit transaction: %v", err)
	}

	return &ledgerpb.AddCreditTransactionResponse{
		Response: "credit transaction added successfully",
	}, nil
}

func (s *Server) AddDebitTransaction(ctx context.Context, req *ledgerpb.AddDebitTransactionRequest) (*ledgerpb.AddDebitTransactionResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Error(codes.Canceled, "client cancelled, aborting")
	}

	t := req.GetTransaction()
	err := s.book.AddTransaction(ledger.TransactionDebit, t.GetWallet(), t.GetDebit(), t.GetAggregate())

	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "problem when adding debit transaction: %v", err)
	}

	return &ledgerpb.AddDebitTransactionResponse{
		Response: "debit transaction added successfully",
	}, nil
}

func (s *Server) AddCashInTransaction(ctx context.Context, req *ledgerpb.AddCashInTransactionRequest) (*ledgerpb.AddCashInTransactionResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Error(codes.Canceled, "client cancelled, aborting")
	}

	t := req.GetTransaction()
	err := s.book.AddTransaction(ledger.TransactionCashIn, t.GetWallet(), t.GetCredit(), t.GetAggregate())

	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "problem when adding cash-in transaction: %v", err)
	}

	return &ledgerpb.AddCashInTransactionResponse{
		Response: "cash in transaction added successfully",
	}, nil
}

func (s *Server) AddCashOutTransaction(ctx context.Context, req *ledgerpb.AddCashOutTransactionRequest) (*ledgerpb.AddCashOutTransactionResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Error(codes.Canceled, "client cancelled, aborting")
	}

	t := req.GetTransaction()
	err := s.book.AddTransaction(ledger.TransactionCashOut, t.GetWallet(), t.GetDebit(), t.GetAggregate())

	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "problem when adding cash-out transaction: %v", err)
	}

	return &ledgerpb.AddCashOutTransactionResponse{
		Response: "cash out transaction added successfully",
	}, nil
}

func (s *Server) WalletTransactions(ctx context.Context, req *ledgerpb.WalletTransactionsRequest) (*ledgerpb.WalletTransactionsResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Error(codes.Canceled, "client cancelled, aborting")
	}

	w := req.GetWallet()

	ts, err := s.book.WalletTransactions(w)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "wallet '%s' has no recorded transactions", w)
	}

	return &ledgerpb.WalletTransactionsResponse{
		Transactions: ts,
	}, nil
}

func (s *Server) WalletBalance(ctx context.Context, req *ledgerpb.WalletBalanceRequest) (*ledgerpb.WalletBalanceResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Error(codes.Canceled, "client cancelled, aborting")
	}

	w := req.GetWallet()

	b, err := s.book.WalletBalance(w)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "wallet '%s' has no recorded transactions", w)
	}

	return &ledgerpb.WalletBalanceResponse{
		Balance: b,
	}, nil
}

func (s *Server) AggregateTransactions(ctx context.Context, req *ledgerpb.AggregateTransactionsRequest) (*ledgerpb.AggregateTransactionsResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Error(codes.Canceled, "client cancelled, aborting")
	}

	ts, err := s.book.AggregateTransactions(req.GetAggregate())

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no transactions found with aggregate ID '%s'", req.GetAggregate())
	}

	return &ledgerpb.AggregateTransactionsResponse{
		Transactions: ts,
	}, nil
}
