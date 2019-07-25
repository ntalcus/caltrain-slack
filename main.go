package main

import (
	// "github.com/ntalcus/caltrainslack/cmd"
	"fmt"
	"github.com/ntalcus/caltrainslack/secrets"
	// conf "github.com/ntalcus/caltrainslack/config"
	transit "github.com/ntalcus/caltrainslack/transit"
	"time"
	// "github.com/spf13/cobra"
)

func main() {
	// cmd.Execute()
	config := transit.Config{}
	sec := secrets.GetSecrets()
	config.APIKey = sec.APIKey
	config.StopCode = "70172"
	config.URL = "http://api.511.org/transit/StopMonitoring"
	config.Method = "GET"
	err := make(chan string)
	go func(e chan string) {

		for x := 0; x < 1000; x++ {
			transit.Call(config)
			time.Sleep(time.Duration(10) * time.Second)
		}
		err <- "lel we done"
	}(err)

	msg := <-err
	fmt.Println(msg)

}
