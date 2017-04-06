package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EntitiesRequest struct {
	Entities []Entity `json:"entities"`
}

type Entity struct {
	Identifier Identifier `json:"identifier"`
}

type Identifier struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

type Response []struct {
	UID    string
	Weight float64
	Edges  []string
}

func main() {
	url := "http://u01.unigraph.rocks/api/context/entities"

	request := EntitiesRequest{
		Entities: []Entity{
			{Identifier{"freebase_id", "/m/0cc883w"}},
			{Identifier{"wikidata_id", "Q57775"}},
			{Identifier{"wikidata_id", "Q133968"}},
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

	for _, entity := range entities {
		fmt.Printf("Entity: %+v\n", entity)
	}
}
