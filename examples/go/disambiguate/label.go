package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DisambiguateRequest struct {
	Label    string `json:"label"`
	Language string `json:"language"`
}

type Response []struct {
	UID    string
	Weight float64
}

func main() {
	url := "http://u01.unigraph.rocks/api/disambiguate"

	request := DisambiguateRequest{"Silvio Berlusconi", "en"}

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
