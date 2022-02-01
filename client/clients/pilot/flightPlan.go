package pilot

import (
	"github.com/xetarusx/ivao-whazzup-parser/client/clients/pilot/flightplan"
	"time"
)

type FlightPlan struct {
	Id                       int64               `json:"id"`
	Revision                 int64               `json:"revision"`
	AircraftId               string              `json:"aircraftId"`
	AircraftNumber           int64               `json:"aircraftNumber"`
	DepartureId              string              `json:"departureId"`
	ArrivalId                string              `json:"arrivalId"`
	AlternativeId            string              `json:"alternativeId"`
	AlternativeId2           string              `json:"alternativeId2"`
	Route                    string              `json:"route"`
	Remarks                  string              `json:"remarks"`
	Speed                    string              `json:"speed"`
	Level                    string              `json:"level"`
	FlightRules              string              `json:"flightRules"`
	FlightType               string              `json:"flightType"`
	EET                      int64               `json:"eet"`
	Endurance                int64               `json:"endurance"`
	DepartureTime            int64               `json:"departureTime"`
	ActualDepartureTime      int64               `json:"actualDepartureTime"`
	PeopleOnBoard            int64               `json:"peopleOnBoard"`
	CreatedAt                time.Time           `json:"createdAt"`
	UpdatedAt                time.Time           `json:"updatedAt"`
	AircraftEquipments       string              `json:"aircraftEquipments"`
	AircraftTransponderTypes string              `json:"aircraftTransponderTypes"`
	Aircraft                 flightplan.Aircraft `json:"aircraft"`
}
