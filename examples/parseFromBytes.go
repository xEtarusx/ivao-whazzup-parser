package examples

import (
	"fmt"
	"github.com/xetarusx/ivao-whazzup-parser/whazzup"
)

func ParseFromBytes(bytes []byte) {
	// Parse bytes into the whazzup struct
	w, err := whazzup.NewWhazzupByJson(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// You can now use all functions on the whazzup struct
	// e.g. Print total network connections
	fmt.Println(w.GetConnectionsTotal())
}
