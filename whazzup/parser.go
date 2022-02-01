package whazzup

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// NewWhazzupByFile Read a json file and try to parse its content into the whazzup struct
func NewWhazzupByFile(path string) (Whazzup, error) {
	file, err := os.Open(path)

	if err != nil {
		return Whazzup{}, err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	return NewWhazzupByJson(byteValue)
}

// NewWhazzupByJson Try to unmarshal bytes into the whazzup struct
func NewWhazzupByJson(bytes []byte) (Whazzup, error) {
	var w Whazzup
	err := json.Unmarshal(bytes, &w)
	return w, err
}
