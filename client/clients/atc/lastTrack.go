package atc

import "time"

type LastTrack struct {
	Distance  int64     `json:"distance"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Time      int64     `json:"time"`
	Timestamp time.Time `json:"timestamp"`
}
