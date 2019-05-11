package main

import (
	"log"
	"net/http"

	http2 "gitlab.com/patchwell/ledger/pkg/http"
	"gitlab.com/patchwell/ledger/pkg/memory"
)

func main() {
	server := &http2.Server{memory.NewInMemoryBook()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
