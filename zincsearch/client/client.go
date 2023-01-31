package client

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
)

type ZinSearchClient struct {
	url         string
	contentType string
	userAgent   string
	basicAuth   string
}

// set the ZincSearch Client
// following the next documentation https://docs.zincsearch.com/api/search/search/#golang-example
func NewZinSearchClient(url string, basicAuth string) ZinSearchClient {
	zinSearchClient := ZinSearchClient{}

	zinSearchClient.url = url
	zinSearchClient.basicAuth = basicAuth
	zinSearchClient.contentType = "application/json"
	zinSearchClient.userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36"

	return zinSearchClient
}

// send bulk to zincSearch
func (client *ZinSearchClient) CreateDocumentsBulk(requestPayload []byte) ([]byte, error) {
	finalUrl := client.url + "/api/_bulkv2"

	createRequest, createRequestErr := http.NewRequest("POST", finalUrl, bytes.NewBuffer(requestPayload))

	if createRequestErr != nil {
		log.Printf("[POST] - ZinSearch Search Request Creation Error: %s", createRequestErr)
		return nil, createRequestErr
	}

	createRequest.SetBasicAuth("admin", client.basicAuth)
	createRequest.Header.Set("Content-Type", client.contentType)
	createRequest.Header.Set("User-Agent", client.userAgent)

	createResponse, createResponseErr := http.DefaultClient.Do(createRequest)

	if createRequestErr != nil {
		log.Printf("[POST] - ZinSearch Search Request Error: %s", createResponseErr)
		return nil, createResponseErr
	}

	defer createResponse.Body.Close()

	createResult, resultErr := io.ReadAll(createResponse.Body)

	if resultErr != nil {
		log.Printf("[GET] - ZinSearch Response Body Parsing Error: %s", resultErr)
		return nil, resultErr
	}

	return createResult, nil

}

func (client *ZinSearchClient) Search(searchQuery string, index string) ([]byte, error) {
	finalUrl := client.url + "/api/" + index + "/_search"

	searchRequest, searchRequestErr := http.NewRequest("POST", finalUrl, strings.NewReader(searchQuery))

	if searchRequestErr != nil {
		log.Printf("[GET] - ZinSearch Search Request Creation Error: %s", searchRequestErr)
		return nil, searchRequestErr
	}

	searchRequest.SetBasicAuth("admin", client.basicAuth)
	searchRequest.Header.Set("Content-Type", client.contentType)
	searchRequest.Header.Set("User-Agent", client.userAgent)

	searchResponse, searchResponseErr := http.DefaultClient.Do(searchRequest)

	if searchResponseErr != nil {
		log.Printf("[GET] - ZinSearch Search Request Error: %s", searchResponseErr)
		return nil, searchResponseErr
	}

	defer searchResponse.Body.Close()

	// reading the response
	searchResult, resultErr := io.ReadAll(searchResponse.Body)

	if resultErr != nil {
		log.Printf("[GET] - ZinSearch Response Body Parsing Error: %s", resultErr)
		return nil, resultErr
	}

	return searchResult, nil
}
