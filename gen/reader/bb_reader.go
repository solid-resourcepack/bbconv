package reader

import (
	"encoding/json"
	"fmt"
	"github.com/solid-resourcepack/bbconv/types"
	"os"
)

func ReadBBModel(file string) *types.BlockBenchModel {
	contents, readErr := os.ReadFile(file)
	if readErr != nil {
		fmt.Printf("Could not access file %s:\n", file)
		panic(readErr)
	}
	var result = &types.BlockBenchModel{}
	parseErr := json.Unmarshal(contents, result)

	if parseErr != nil {
		fmt.Println("Could not parse BlockBench model:")
		panic(parseErr)
	}
	return result
}
