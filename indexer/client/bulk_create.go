package client

import (
	"log"

	"github.com/OliverosA/indexer/files"
	"github.com/OliverosA/zincsearch/client"
)

const (
	zsUrl    = "http://localhost:4080"
	zsSecret = "Complexpass#123"
)

func DocumentsBulkCreation(docs []files.EmailJsonData) []byte {

	zsClient := client.NewZinSearchClient(zsUrl, zsSecret)
	requestPayload := BuildJsonBulkCreationPayload(docs)

	bulkCreationRes, bulkCreationErr := zsClient.CreateDocumentsBulk(requestPayload)

	if bulkCreationErr != nil {
		log.Fatalf("Error in ZincSearch data indexing process: %s", bulkCreationErr)
	}

	return bulkCreationRes

}
