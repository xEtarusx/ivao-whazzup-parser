package clients

import "github.com/xetarusx/ivao-whazzup-parser/client/clients/atc"

type ATC struct {
	Time            int64         `json:"time"`
	Id              int64         `json:"id"`
	UserId          int64         `json:"userId"`
	Callsign        string        `json:"callsign"`
	ServerId        string        `json:"serverId"`
	SoftwareTypeId  string        `json:"softwareTypeId"`
	SoftwareVersion string        `json:"softwareVersion"`
	Rating          int64         `json:"rating"`
	CreatedAt       string        `json:"createdAt"`
	ATCSession      atc.Session   `json:"atcSession"`
	ATIS            atc.ATIS      `json:"atis"`
	LastTrack       atc.LastTrack `json:"lastTrack"`
}
