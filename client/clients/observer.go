package clients

import (
	"github.com/xetarusx/ivao-whazzup-parser/client/clients/atc"
	"time"
)

type Observer struct {
	Time            int64         `json:"time"`
	Id              int64         `json:"id"`
	UserId          int64         `json:"userId"`
	Callsign        string        `json:"callsign"`
	ServerId        string        `json:"serverId"`
	SoftwareTypeId  string        `json:"softwareTypeId"`
	SoftwareVersion string        `json:"softwareVersion"`
	CreatedAt       time.Time     `json:"createdAt"`
	ATCSession      atc.Session   `json:"atcSession"`
	LastTrack       atc.LastTrack `json:"lastTrack"`
}
