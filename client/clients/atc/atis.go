package atc

import (
	"github.com/xetarusx/ivao-whazzup-parser/misc"
	"strings"
	"time"
)

type ATIS struct {
	Lines     []string  `json:"lines"`
	Revision  string    `json:"revision"`
	Timestamp time.Time `json:"timestamp"`
}

// GetText Get ATIS as text
func (a ATIS) GetText() string {
	return strings.Join(a.Lines, " ")
}

// GetStationName Extract station name from ATIS
func (a ATIS) GetStationName() string {
	// Check if atis text is available with more than 2 lines
	if len(a.Lines) <= 1 {
		return ""
	}

	// Currently, the station name is contained in line 2 (only guesswork, may be not 100% accurate)
	atisLine2 := a.Lines[1]

	// Get substring before string "Information"
	pos := strings.Index(atisLine2, "Information")
	if pos == -1 {
		// Substring not found, search for lowercase substring "information"
		pos = strings.Index(atisLine2, "information")
		if pos == -1 {
			// Return atis line 2
			return atisLine2
		}
	}

	// Get text before the found substring ("Information"/"information")
	textBeforeInformation := atisLine2[0:pos]

	return textBeforeInformation
}

// GetRevisionLong Get current ATIS revision word (e.g. Bravo, X-ray, etc.) by letter
func (a ATIS) GetRevisionLong() string {
	return misc.GetIcaoCodeWordByLetter(a.Revision)
}
