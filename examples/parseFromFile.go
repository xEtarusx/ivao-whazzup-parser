package examples

import (
	"fmt"
	"github.com/xetarusx/ivao-whazzup-parser/whazzup"
)

func ParseFromFile(path string) {
	// Read file in path and parse content into the whazzup struct
	w, err := whazzup.NewWhazzupByFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// You can now use all functions on the whazzup struct
	// e.g. Print total network connections
	fmt.Println(w.GetConnectionsTotal())
}
