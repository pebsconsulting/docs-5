package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DisambiguateRequest struct {
	Label    string   `json:"label"`
	Language string   `json:"language"`
	Types    []string `json:"types,omitempty"`
}

type Response []struct {
	UID    string
	Weight float64
}

func main() {
	url := "http://u01.unigraph.rocks/api/disambiguate"

	request := DisambiguateRequest{
		Label:    "Silvio Berlusconi",
		Language: "en",
		Types:    []string{"12fa01440b"},
	}

	jsonRequest, _ := json.Marshal(request)
	payload := bytes.NewReader(jsonRequest)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("x-unigraph-api-key", "WebDemos")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	entities := Response{}
	json.Unmarshal(body, &entities)

	for _, entity := range entities {
		fmt.Printf("Entity: %+v\n", entity)
	}
}
