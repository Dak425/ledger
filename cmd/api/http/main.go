package main

import (
	"log"
	"net/http"

	ledgerhttp "gitlab.com/patchwell/ledger/pkg/api/server/http"
	"gitlab.com/patchwell/ledger/pkg/book/memory"
)

func main() {
	book := memory.NewMockInMemoryBook()
	server := ledgerhttp.NewServer(book)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
