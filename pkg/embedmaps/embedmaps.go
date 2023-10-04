package embedmaps

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed maps/*
var EmbeddedMaps embed.FS

// GetMapMatrixFromEmbed returns a MapMatrix for a Map, from an embedded json file.
func GetMapMatrixFromEmbed(path string) (matrix [][]uint16, err error) {
	// Read the embedded json file
	file, err := EmbeddedMaps.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening embedded map: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("error closing embedded map: %w", cerr)
		}
	}()

	// Decode the JSON from the file into the matrix variable
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&matrix); err != nil {
		return nil, fmt.Errorf("error decoding embedded JSON: %w", err)
	}

	return matrix, nil
}
