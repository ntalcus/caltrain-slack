package transit

// Message is the array of ServiceDelivery.
type Message struct {
	ServiceDelivery ServiceDelivery `json:"ServiceDelivery"`
}

// ServiceDelivery is datatype representing the CalTrain StopMonitoring API response.
type ServiceDelivery struct {
	ResponseTimestamp      string                 `json:"ResponseTimestamp"`
	ProducerRef            string                 `json:"ProducerRef"`
	Status                 bool                   `json:"Status"`
	StopMonitoringDelivery StopMonitoringDelivery `json:"StopMonitoringDelivery"`
}

// StopMonitoringDelivery represents subtype of ServiceDelivery.
type StopMonitoringDelivery struct {
	Version             string               `json:"version"`
	ResponseTimestamp   string               `json:"ResponseTimestamp"`
	Status              bool                 `json:"Status"`
	MonitoredStopVisits []MonitoredStopVisit `json:"MonitoredStopVisit"`
}

// MonitoredStopVisit represents subtype of StopMonitoringDelivery.
type MonitoredStopVisit struct {
	RecordedAtTime          string                  `json:"RecordedAtTime"`
	MonitoringRef           string                  `json:"MonitoringRef"`
	MonitoredVehicleJourney MonitoredVehicleJourney `json:"MonitoredVehicleJourney"`
}

// MonitoredVehicleJourney another subtype lel.
type MonitoredVehicleJourney struct {
	LineRef                 string                  `json:"LineRef"`
	DirectionRef            string                  `json:"DirectionRef"`
	FramedVehicleJourneyRef FramedVehicleJourneyRef `json:"FramedVehicleJourneyRef"`
	PublishedLineName       string                  `json:"PublishedLineName"`
	OperatorRef             string                  `json:"OperatorRef"`
	OriginRef               string                  `json:"OriginRef"`
	OriginName              string                  `json:"OriginName"`
	DestinationRef          string                  `json:"DestinationRef"`
	DestinationName         string                  `json:"DestinationName"`
	Monitored               bool                    `json:"Monitored"`
	InCongestion            bool                    `json:"InCongestion"`
	VehicleLocation         VehicleLocation         `json:"VehicleLocation"`
	Bearing                 string                  `json:"Bearing"`
	Occupancy               string                  `json:"Occupancy"`
	VehicleRef              string                  `json:"VehicleRef"`
	MonitoredCall           MonitoredCall           `json:"MonitoredCall"`
}

// FramedVehicleJourneyRef lel.
type FramedVehicleJourneyRef struct {
	DataFrameRef           string `json:"DataFrameRef"`
	DatedVehicleJourneyRef string `json:"DatedVehicleJourneyRef"`
}

// VehicleLocation lel.
type VehicleLocation struct {
	Longitude string `json:"Longitude"`
	Latitude  string `json:"Latitude"`
}

// MonitoredCall doc to beat that style error.
type MonitoredCall struct {
	StopPointRef          string `json:"StopPointRef"`
	StopPointName         string `json:"StopPointName"`
	VehicleLocationAtStop string `json:"VehicleLocationAtStop"`
	VehicleAtStop         string `json:"VehicleAtStop"`
	AimedArrivalTime      string `json:"AimedArrivalTime"`
	ExpectedArrivalTime   string `json:"ExpectedArrivalTime"`
	AimedDepartureTime    string `json:"AimedDepartureTime"`
	ExpectedDepartureTime string `json:"ExpectedDepartureTime"`
	Distances             string `json:"Distances"`
}
