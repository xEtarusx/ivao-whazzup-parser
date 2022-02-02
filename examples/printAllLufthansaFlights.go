package examples

import (
	"fmt"
	"github.com/xetarusx/ivao-whazzup-parser/whazzup"
)

func printAllLufthansaFlights(path string) {
	// Read file in path and parse content into the whazzup struct
	w, err := whazzup.NewWhazzupByFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// You can now use all functions on the whazzup struct
	// e.g. Search all pilots starting with "DLH" (Lufthansa) callsign and print them
	for _, p := range w.Clients.GetPilotsByCallsignPrefix("DLH") {
		fmt.Println(p.Callsign)
	}
}
