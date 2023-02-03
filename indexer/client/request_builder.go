package client

import (
	"encoding/json"

	"github.com/OliverosA/indexer/files"
)

type EmailJsonBulkCreationPayload struct {
	Index   string                `json:"index"`
	Records []files.EmailJsonData `json:"records"`
}

func BuildJsonBulkCreationPayload(docs []files.EmailJsonData) []byte {
	requestPayload := EmailJsonBulkCreationPayload{
		Index:   "emails",
		Records: docs,
	}

	payloadStringified, _ := json.MarshalIndent(requestPayload, "", "  ")

	return payloadStringified
}
