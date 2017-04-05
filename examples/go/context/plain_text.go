package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Response []struct {
	UID    string
	Weight float64
	Stated string
	Edges  []string
}

func main() {
	url := "http://u01.unigraph.rocks/api/context/text"
	payload := strings.NewReader("Larry Page and Sergey Brin")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("x-unigraph-api-max", "20") // maximum 50
	req.Header.Add("x-unigraph-api-key", "WebDemos")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	entities := Response{}
	if err := json.Unmarshal(body, &entities); err != nil {
		log.Fatal(err)
	}
	for _, entity := range entities {
		fmt.Printf("Entity: %+v\n", entity)
	}
}
