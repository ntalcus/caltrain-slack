package transit

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "github.com/ntalcus/caltrainslack/secrets"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Config represents the options users need to specify for their specific CalTrain stop such that
// the API can be called to return the proper information.
type Config struct {
	StopCode string
	APIKey   string
	URL      string
	Method   string
}

func GetLineInformation() {

}

// Call calls the 511.org API to get information about a train.
func Call(config Config) {
	tNow := time.Now()
	fmt.Println(tNow)

	method := config.Method
	url := config.URL
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("api_key", config.APIKey)
	q.Add("agency", "CT")
	q.Add("stopCode", config.StopCode)
	q.Add("format", "Json")
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	respText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	respText = bytes.TrimPrefix(respText, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	var APICallResp = Message{}

	err = json.Unmarshal(respText, &APICallResp)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	// fmt.Printf("%+v\n\n", APICallResp)
	// fmt.Printf()
	printDetails(APICallResp)

}

func printDetails(mes Message) {
	stops := mes.ServiceDelivery.StopMonitoringDelivery.MonitoredStopVisits
	// fmt.Printf("%s: # Stops %d\n\n", time.Now(), len(stops))
	for _, stopDelivery := range stops {
		stop := stopDelivery.MonitoredVehicleJourney
		// fmt.Println(stop.LineRef)
		// fmt.Println(stop.PublishedLineName)
		// fmt.Println(stop.DirectionRef)
		// fmt.Println(stop.Monitored)
		fmt.Printf("TRAIN: %s (%s) headed %s to %s. Stopping at %s expected at %s (scheduled for %s).\n", stop.LineRef, stop.OriginName, stop.DirectionRef, stop.DestinationName, stop.MonitoredCall.StopPointName, stop.MonitoredCall.ExpectedArrivalTime, stop.MonitoredCall.AimedArrivalTime)
	}
}
