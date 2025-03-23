package bbformat

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadBBModel(file string) *Model {
	contents, readErr := os.ReadFile(file)
	if readErr != nil {
		fmt.Printf("Could not access file %s:\n", file)
		panic(readErr)
	}
	var result = &Model{}
	parseErr := json.Unmarshal(contents, result)

	if parseErr != nil {
		fmt.Println("Could not parse BlockBench model:")
		panic(parseErr)
	}
	return result
}
