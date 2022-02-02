package client

import (
	"errors"
	"github.com/xetarusx/ivao-whazzup-parser/client/clients"
	"strings"
)

type Clients struct {
	Pilots       []clients.Pilot       `json:"pilots"`
	ATCs         []clients.ATC         `json:"atcs"`
	FollowMeCars []clients.FollowMeCar `json:"followMe"`
	Observers    []clients.Observer    `json:"oberservers"`
}

// GetPilotCallsigns Get list of all pilot callsigns
func (c Clients) GetPilotCallsigns() []string {
	var list []string

	for _, p := range c.Pilots {
		list = append(list, p.Callsign)
	}

	return list
}

// GetPilotByCallsign Return pilot struct for a specific callsign
func (c Clients) GetPilotByCallsign(callsign string) (clients.Pilot, error) {
	for _, p := range c.Pilots {
		if p.Callsign == callsign {
			return p, nil
		}
	}

	return clients.Pilot{}, errors.New("no pilot found")
}

// GetPilotsByCallsignPrefix Return all pilots where callsigns begin with prefix
func (c Clients) GetPilotsByCallsignPrefix(prefix string) []clients.Pilot {
	var list []clients.Pilot

	for _, p := range c.Pilots {
		if strings.HasPrefix(p.Callsign, prefix) {
			list = append(list, p)
		}
	}

	return list
}
