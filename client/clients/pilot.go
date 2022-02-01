package clients

import (
	"github.com/xetarusx/ivao-whazzup-parser/client/clients/pilot"
	"time"
)

type Pilot struct {
	Time            int64            `json:"time"`
	Id              int64            `json:"id"`
	UserId          int64            `json:"userId"`
	Callsign        string           `json:"callsign"`
	ServerId        string           `json:"serverId"`
	SoftwareTypeId  string           `json:"softwareTypeId"`
	SoftwareVersion string           `json:"softwareVersion"`
	Rating          int64            `json:"rating"`
	CreatedAt       time.Time        `json:"createdAt"`
	FlightPlan      pilot.FlightPlan `json:"flightPlan"`
	PilotSession    pilot.Session    `json:"pilotSession"`
	LastTrack       pilot.LastTrack  `json:"lastTrack"`
}
