package embedimages

import (
	"embed"
	"fmt"
	"image"
	"image/png"
)

//go:embed static/*
var EmbeddedImages embed.FS

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
