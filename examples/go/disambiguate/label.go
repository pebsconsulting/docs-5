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

	req.Header.Add("x-unigraph-api-key", "1c42f9c1ca2f65441465b43cd9339d6c")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "e6a951ed-36a6-9705-8ef2-e8a602b0fea0")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
