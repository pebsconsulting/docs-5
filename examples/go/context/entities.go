package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Identifier struct {
	Property string `json:"property,omitempty"`
	Value    string `json:"value,omitempty"`
}

type Entity struct {
	Identifier *Identifier `json:"identifier,omitempty"`
	Label      string      `json:"label,omitempty"`
	Language   string      `json:"language,omitempty"`
	Types      []string    `json:"types,omitempty"`
}

type EntitiesRequest struct {
	Entities []Entity `json:"entities"`
}

type Response []struct {
	UID    string
	Weight float64
	Tag    string
	Edges  []string
}

func main() {

	url := "http://u01.unigraph.rocks/api/context/entities"

	request := EntitiesRequest{
		Entities: []Entity{
			{Label: "Google", Language: "en"},
			{Label: "Larry Page", Language: "de", Types: []string{"1205"}},
			{Identifier: &Identifier{
				Property: "wikidata_id",
				Value:    "Q92764"},
			},
		},
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

	fmt.Println(string(jsonRequest))

	for _, entity := range entities {
		fmt.Printf("Entity: %+v\n", entity)
	}
}
