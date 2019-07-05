package main

import (
	"log"
	"net/http"
	"os"
	// "strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	method := "GET"
	url := "http://api.511.org/transit/StopMonitoring"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("api_key", secret)
	q.Add("agency", "CT")
	q.Add("stopCode", "70172")
	q.Add("format", "Json")
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Println(res.Status)
	respText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Printf("%s", respText)
	// fmt.Println(res)

	err = json.Unmarshal(respText, jObj)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

}