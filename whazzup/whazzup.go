package whazzup

import (
	"errors"
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

// GetServerIds Get list of all server ids
func (w Whazzup) GetServerIds() []string {
	var list []string

	for _, s := range w.Servers {
		list = append(list, s.Id)
	}

	return list
}

// GetServerById Return server struct for a specific id
func (w Whazzup) GetServerById(id string) (server.Server, error) {
	for _, s := range w.Servers {
		if s.Id == id {
			return s, nil
		}
	}

	return server.Server{}, errors.New("server not found")
}

// GetVoiceServerIds Get list of all voice server ids
func (w Whazzup) GetVoiceServerIds() []string {
	var list []string

	for _, s := range w.Servers {
		list = append(list, s.Id)
	}

	return list
}

// GetVoiceServerById Return VoiceServer struct for a specific id
func (w Whazzup) GetVoiceServerById(id string) (voiceServer.VoiceServer, error) {
	for _, vs := range w.VoiceServers {
		if vs.Id == id {
			return vs, nil
		}
	}

	return voiceServer.VoiceServer{}, errors.New("server not found")
}
