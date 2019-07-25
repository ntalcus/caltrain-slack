package config

import (
	"encoding/json"
	"fmt"
	"github.com/ntalcus/caltrainslack/secrets"
	"io/ioutil"
)

// Config represents the options users need to specify for their specific CalTrain stop such that
// the API can be called to return the proper information.
type Config struct {
	StopCode string `json:"stopcode"`
	APIKey   string `json:"apikey"`
	URL      string `json:"url"`
	Method   string `json:"method"`
}

// GetConfig returns a Config struct that can be passed to the transit.Call function.
func GetConfig() Config {
	sec := Config{}
	file, err := ioutil.ReadFile("./config/config.json")
	// fmt.Println(string(file))
	if err != nil {
		fmt.Println("Could not parse secrets file")
		panic(err)
	}
	// fmt.Println(file)
	err = json.Unmarshal(file, &sec)
	if err != nil {
		panic(err)
	}
	return sec
}
