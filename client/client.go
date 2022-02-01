package client

import "github.com/xetarusx/ivao-whazzup-parser/client/clients"

type Clients struct {
	Pilots       []clients.Pilot       `json:"pilots"`
	ATCs         []clients.ATC         `json:"atcs"`
	FollowMeCars []clients.FollowMeCar `json:"followMe"`
	Observers    []clients.Observer    `json:"oberservers"`
}
