package main

import (
	"context"
	"fmt"
	"log"

	ledgerpb "gitlab.com/patchwell/ledger/gen/api/protobuf"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error when dialing server, %v", err)
	}

	c := ledgerpb.NewLedgerServiceClient(cc)

	wr := &ledgerpb.WalletBalanceRequest{
		Wallet: "1",
	}
	
	res, err := c.WalletBalance(context.Background(), wr)

	if err != nil {
		log.Fatalf("error when retreiving wallet balance, %v", err)
	}

	fmt.Printf("got response for wallet balance: %v", res)
}
