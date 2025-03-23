package baseformat

import (
	"encoding/json"
	"log"
	"os"
)

func WriteModel(file string, model *Model) error {
	write, err := json.MarshalIndent(&model, "", "  ")
	if err != nil {
		return err
	}
	if _, err := os.Stat(file); err != nil {
		_, err := os.Create(file)
		if err != nil {
			panic(err)
			return err
		}
	}

	log.Printf("Writing model to file: %s\n", file)
	err = os.WriteFile(file, write, 0644)
	if err != nil {
		return err
	}
	return nil
}
