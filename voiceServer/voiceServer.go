package voiceServer

type VoiceServer struct {
	Id                 string `json:"id"`
	Hostname           string `json:"hostname"`
	IP                 string `json:"ip"`
	Description        string `json:"description"`
	CountryId          string `json:"countryId"`
	CurrentConnections int64  `json:"currentConnections"`
	MaximumConnections int64  `json:"MaximumConnections"`
}
