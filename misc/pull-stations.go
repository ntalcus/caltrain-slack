package misc

import (
	"github.com/ntalcus/caltrainslack/secrets"
	"github.com/ntalcus/caltrainslack/transit"
)

func pullStationList() {
	config := transit.Config{}
	sec := secrets.GetSecrets()
	config.APIKey = sec.APIKey
	config.StopCode = "70172"
	config.URL = "http://api.511.org/transit/StopMonitoring"
	config.Method = "GET"

	method := config.Method
	url := config.URL
	req, err := http.NewRequest(method, url, nil)

}

type Station int

const {
	SAN_FRANCISCO Station = iota
	22ND_ST
	BAYSHORE
	SOUTH_SAN_FRAN
	SAN_BRUNO
	MILLBRAE
	BROADWAY
	BURLINGAME
	SAN_MATEO
	HAYWARD_PARK
	HILLSDALE
	BELMONT
	SAN_CARLOS
	REDWOOD_CITY
	ATHERTON
	MENLO_PARK
	PALO_ALTO
	CALIFORNIA_AVE
	SAN_ANTONIO
	MOUNTAIN_VIEW
	SUNNYVALE
	LAWRENCE
	SANTA_CLARA
	COLLEGE_PARK
	SAN_JOSE_DIRIDON
	TAMIEN
	CAPITOL
	BLOSSOM_HILL
	MORGAN_HILL
	SAN_MARTIN
	GILROY
}

var stations = []string {
	"SAN FRANCISCO", 
	"22ND ST",
	"BAYSHORE",
	"SOUTH SAN FRAN",
	"SAN BRUNO",
	"MILLBRAE",
	"BROADWAY",
	"BURLINGAME",
	"SAN MATEO",
	"HAYWARD PARK",
	"HILLSDALE",
	"BELMONT",
	"SAN CARLOS",
	"REDWOOD CITY",
	"ATHERTON",
	"MENLO PARK",
	"PALO ALTO",
	"CALIFORNIA AVE",
	"SAN ANTONIO",
	"MOUNTAIN VIEW",
	"SUNNYVALE",
	"LAWRENCE",
	"SANTA CLARA",
	"COLLEGE PARK",
	"SAN JOSE DIRIDON",
	"TAMIEN",
	"CAPITOL",
	"BLOSSOM HILL",
	"MORGAN HILL",
	"SAN MARTIN",
	"GILROY",
}

func getName(s Station) { 
	return stations[s]
}
