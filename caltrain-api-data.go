package main

// ServiceDelivery is datatype representing the CalTrain StopMonitoring API response.
type ServiceDelivery struct {
	ResponseTimestamp      string
	ProducerRef            string
	Status                 bool
	StopMonitoringDelivery StopMonitoringDelivery
}

// StopMonitoringDelivery represents subtype of ServiceDelivery.
type StopMonitoringDelivery struct {
	version             string
	ResponseTimestamp   string
	Status              bool
	MonitoredStopVisits []MonitoredStopVisit
}

// MonitoredStopVisit represents subtype of StopMonitoringDelivery.
type MonitoredStopVisit struct {
	RecordedAtTime          string
	MonitoringRef           string
	MonitoredVehicleJourney MonitoredVehicleJourney
}

// MonitoredVehicleJourney another subtype lel.
type MonitoredVehicleJourney struct {
	LineRef                 string
	DirectionRef            string
	FramedVehicleJourneyRef FramedVehicleJourneyRef
	PublishedLineName       string
	OperatorRef             string
	Monitored               bool
	VehicleLocation         VehicleLocation
}

// FramedVehicleJourneyRef lel.
type FramedVehicleJourneyRef struct {
	DataFrameRef           string
	DatedVehicleJourneyRef string
}

// VehicleLocation lel.
type VehicleLocation struct {
	Longitude string
	Latitude  string
}
