package client

import (
	"log"

	"github.com/OliverosA/indexer/files"
	"github.com/OliverosA/zincsearch/client"
)

// create documents bulk
func DocumentsBulkCreation(docs []files.EmailJsonData) []byte {

	// set zincsearch localhost url and password
	zsUrl := "http://localhost:4080"
	zsSecret := "Complexpass#123"

	zsClient := client.NewZinSearchClient(zsUrl, zsSecret)
	requestPayload := BuildJsonBulkCreationPayload(docs)

	bulkCreationRes, bulkCreationErr := zsClient.CreateDocumentsBulk(requestPayload)

	if bulkCreationErr != nil {
		log.Fatalf("Error in ZincSearch data indexing process: %s", bulkCreationErr)
	}

	return bulkCreationRes

}
