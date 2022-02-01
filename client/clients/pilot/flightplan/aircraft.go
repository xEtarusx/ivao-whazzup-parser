package flightplan

type Aircraft struct {
	IcaoCode       string `json:"icaoCode"`
	Model          string `json:"model"`
	WakeTurbulence string `json:"wakeTurbulence"`
	IsMilitary     bool   `json:"isMilitary"`
	Description    string `json:"description"`
}
