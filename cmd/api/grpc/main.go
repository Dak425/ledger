package main

import (
	grpc2 "gitlab.com/patchwell/ledger/pkg/grpc"
	"gitlab.com/patchwell/ledger/pkg/memory"
	"log"
	"net"

	"gitlab.com/patchwell/ledger/gen/api/protobuf"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen, %v", err)
	}

	s := grpc.NewServer()
	book := memory.NewInMemoryBook()
	ledgerpb.RegisterLedgerServiceServer(s, grpc2.NewGRPCServer(book))

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve, %v", err)
	}
}
