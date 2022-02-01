package pilot

import "time"

type LastTrack struct {
	Altitude           int64     `json:"altitude"`
	AltitudeDifference int64     `json:"altitudeDifference"`
	ArrivalDistance    float64   `json:"arrivalDistance"`
	DepartureDistance  float64   `json:"departureDistance"`
	GroundSpeed        int64     `json:"groundSpeed"`
	Heading            int64     `json:"heading"`
	Latitude           float64   `json:"latitude"`
	Longitude          float64   `json:"longitude"`
	OnGround           bool      `json:"onGround"`
	State              string    `json:"state"`
	Time               int64     `json:"time"`
	Timestamp          time.Time `json:"timestamp"`
	Transponder        int64     `json:"transponder"`
	TransponderMode    string    `json:"transponderMode"`
}
