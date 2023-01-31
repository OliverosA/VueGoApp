package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/OliverosA/zincsearch/client"
)

const (
	// set zincsearch localhost url and password
	zsUrl    = "http://localhost:4080"
	zsSecret = "Complexpass#123"
)

type Email struct {
	Term string
}

func GetEmails(w http.ResponseWriter, r *http.Request) {
	zsClient := client.NewZinSearchClient(zsUrl, zsSecret)

	emailsRequest, emailsRequestErr := zsClient.Search(`
		{
			"search_type": "matchall",
			"from": 0,
			"max_results": 40,
			"_source": []
		}`, "emails")

	if emailsRequestErr != nil {
		w.Write([]byte(fmt.Sprintf("Error: %v", emailsRequestErr)))
		return
	}

	w.Write(emailsRequest)

}

func SearchTerm(w http.ResponseWriter, r *http.Request) {
	zsClient := client.NewZinSearchClient(zsUrl, zsSecret)
	email := Email{}

	defer r.Body.Close()

	body, bodyErr := io.ReadAll(r.Body)

	if bodyErr != nil {
		w.Write([]byte(fmt.Sprintf("Error reading body: %v", bodyErr)))
	}

	unMarshalErr := json.Unmarshal(body, &email)

	if unMarshalErr != nil {
		w.Write([]byte(fmt.Sprintf("Error trying to unmarshal body: %v", unMarshalErr)))
	}

	emailsRequest, emailsRequestErr := zsClient.Search(fmt.Sprintf(`
		{
			"search_type": "match",
			"query": {
				"term": "%v"
			},
			"sort_fields": ["-@timestamp"],
			"from": 0,
			"max_results": 40,
			"_source": []
	    }`, email.Term), "emails")

	if emailsRequestErr != nil {
		w.Write([]byte(fmt.Sprintf("ZincSearch Error: %v", emailsRequestErr)))
		return
	}

	w.Write(emailsRequest)

}
