package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockResponse struct {
	Status string
}

func TestZinSearchClientSearchMethod(t *testing.T) {
	searchQuery := `{
        "search_type": "match",
        "query":
        {
            "term": "DEMTSCHENKO",
            "start_time": "2021-06-02T14:28:31.894Z",
            "end_time": "2021-12-02T15:28:31.894Z"
        },
        "from": 0,
        "max_results": 20,
        "_source": []
    }`

	testServer := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("{\"status\":\"ok\"}"))
	}))

	defer testServer.Close()

	zinSearchClient := NewZinSearchClient(testServer.URL, "password")

	searchResult, _ := zinSearchClient.Search(searchQuery, "emails")

	var response MockResponse

	json.Unmarshal(searchResult, &response)

	if response.Status != "ok" {
		t.Errorf("expected response to be %s got %s", "ok", response.Status)
	}

}
