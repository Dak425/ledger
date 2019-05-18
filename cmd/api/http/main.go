package main

import (
	"gitlab.com/patchwell/ledger/pkg/book/file"
	"log"
	"net/http"
	"os"

	ledgerhttp "gitlab.com/patchwell/ledger/pkg/api/server/http"
)

const dbFileName = "transactions.db.json"

func main() {
	f, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("unable to open file %s, %v", dbFileName, err)
	}

	book, err := file.NewFileSystemBook(f)
	if err != nil {
		log.Fatalf("problem when creating file system book, %v", err)
	}

	server := ledgerhttp.NewServer(book)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
