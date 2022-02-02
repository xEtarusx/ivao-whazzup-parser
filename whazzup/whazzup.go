package whazzup

import (
	"github.com/xetarusx/ivao-whazzup-parser/client"
	"github.com/xetarusx/ivao-whazzup-parser/connections"
	"github.com/xetarusx/ivao-whazzup-parser/server"
	"github.com/xetarusx/ivao-whazzup-parser/voiceServer"
	"time"
)

type Whazzup struct {
	UpdatedAt    time.Time                 `json:"updatedAt"`
	Servers      []server.Server           `json:"servers"`
	VoiceServers []voiceServer.VoiceServer `json:"voiceServers"`
	Clients      client.Clients            `json:"clients"`
	Connections  connections.Connections   `json:"connections"`
}

// GetConnectionsTotal Get total network connections
func (w Whazzup) GetConnectionsTotal() int64 {
	return w.Connections.Total
}
