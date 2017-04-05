package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response []struct {
	UID    string   `json:"uid"`
	Weight int      `json:"weight"`
	Stated string   `json:"stated,omitempty"`
	Edges  []string `json:"edges,omitempty"`
}

func main() {

	url := "http://u01.unigraph.rocks/api/context/text"

	payload := strings.NewReader("Larry Page and Sergey Brin")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("max", "20") // maximum 50
	req.Header.Add("x-unigraph-api-key", "WebDemos")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
