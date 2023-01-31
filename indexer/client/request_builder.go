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

	// create the slice into a json
	// pretty json format with json.MarshalIndent
	payloadStringified, _ := json.MarshalIndent(requestPayload, "", "  ")

	return payloadStringified
}
