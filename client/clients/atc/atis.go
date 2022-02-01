package atc

type ATIS struct {
	Lines     []string `json:"lines"`
	Revision  string   `json:"revision"`
	Timestamp string   `json:"timestamp"`
}
