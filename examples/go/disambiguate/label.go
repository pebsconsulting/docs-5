package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response []struct {
	UID    string  `json:"uid"`
	Weight float64 `json:"weight"`
}

func main() {

	url := "http://u01.unigraph.rocks/api/disambiguate"

	payload := strings.NewReader("{\"label\": \"Silvio Berlusconi\", \"language\": \"en\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("x-unigraph-api-key", "WebDemos")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
