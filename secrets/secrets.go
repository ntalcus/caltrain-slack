package secrets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "os"
)

// Secrets including apikeys
type Secrets struct {
	APIKey string
}

// GetSecrets lololol
func GetSecrets() Secrets {
	sec := Secrets{}
	file, err := ioutil.ReadFile("./secrets/secrets.json")
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
