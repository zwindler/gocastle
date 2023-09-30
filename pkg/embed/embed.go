package embed

import (
	"embed"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
)

var (
	//go:embed static/*
	EmbeddedImages embed.FS
	//go:embed maps/*
	EmbeddedMaps embed.FS
)

// GetImageFromEmbed will return an image.Image from an embedded file.
func GetImageFromEmbed(path string) (img image.Image, err error) {
	// Read the embedded image file
	file, err := EmbeddedImages.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening embedded image: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("error closing embedded image: %w", cerr)
		}
	}()

	// Decode the image using the png package
	img, err = png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("error decoding embedded image: %w", err)
	}

	return img, nil
}

// GetMapMatrixFromEmbed returns a MapMatrix for a Map, from an embedded json file.
func GetMapMatrixFromEmbed(path string) (matrix [][]int, err error) {
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
