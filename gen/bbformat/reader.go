package bbformat

import (
	"encoding/json"
	"log"
	"os"
)

func ReadModel(file string) *Model {
	contents, readErr := os.ReadFile(file)
	if readErr != nil {
		log.Fatal(readErr)
		return nil
	}
	var result = &Model{}
	parseErr := json.Unmarshal(contents, result)

	if parseErr != nil {
		log.Fatal(parseErr)
		return nil
	}
	return result
}
