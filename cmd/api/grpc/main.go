package main

import (
	"log"
	"net"

	ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"
	ledgergrpc "gitlab.com/patchwell/ledger/pkg/api/server/grpc"
	"gitlab.com/patchwell/ledger/pkg/book/memory"

	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen, %v", err)
	}

	s := grpc.NewServer()
	book := memory.NewInMemoryBook()
	ledgerpb.RegisterLedgerServiceServer(s, ledgergrpc.NewGRPCServer(book))

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve, %v", err)
	}
}
