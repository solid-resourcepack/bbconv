package mcformat

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WritePackData(data MCPackData, outDir string, namespace string) error {
	folder := fmt.Sprintf("%sassets/%s/", outDir, namespace)
	err := WriteModels(folder, data.Models)
	if err != nil {
		return err
	}
	err = WriteTextures(folder, data.ModelName, data.Textures)
	if err != nil {
		return err
	}

	return nil
}

func WriteTextures(basePath string, modelName string, textures map[int16]string) error {
	texturesPath := filepath.Join(basePath, "textures", "item")
	err := os.MkdirAll(texturesPath, os.ModePerm)
	if err != nil {
		return err
	}
	for id, texture := range textures {
		out := fmt.Sprintf("%s/%s_%d.png", texturesPath, modelName, id)
		data, err := base64.StdEncoding.DecodeString(strings.Replace(texture, "data:image/png;base64,", "", 1))
		if err != nil {
			return fmt.Errorf("failed to decode base64: %w", err)
		}
		// Write the decoded data to a PNG file
		if err := os.WriteFile(out, data, 0644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}
	return nil
}

func WriteModels(basePath string, models map[string]Model) error {
	for path, model := range models {
		// Replace ":" with the OS-specific path separator
		split := strings.Split(path, ":")
		filePath := filepath.Join(basePath, "models", split[1])

		// Ensure the directory exists
		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, 0644); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}

		// Marshal model data to JSON
		data, err := json.MarshalIndent(model, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal model data: %w", err)
		}

		// Write to file
		if err := os.WriteFile(filePath+".json", data, 0644); err != nil {
			return fmt.Errorf("failed to write model file %s: %w", filePath, err)
		}
	}

	return nil
}
