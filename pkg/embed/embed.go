package embed

import (
	"embed"
	"fmt"
	"image"
	"image/png"
	"io/fs"
)

//go:embed static/*
var EmbeddedImages embed.FS

func GetImageFromEmbed(path string) (img image.Image, err error) {
	// Read the embedded image file
	closeFunc := func(f fs.File) (err error) {
		if f != nil {
			if err := f.Close(); err != nil {
				return err
			}
		}
		return nil
	}

	file, err := EmbeddedImages.Open(path)
	if err != nil {
		// log.Fatal will cause the program to exit with an error code
		closeFunc(file)
		return nil, fmt.Errorf("error opening embedded image: %w", err)
	}

	// Decode the image using the png package
	img, err = png.Decode(file)
	if err != nil {
		// log.Fatal will cause the program to exit with an error code
		closeFunc(file)
		return nil, fmt.Errorf("error decoding embedded image: %w", err)
	}

	closeFunc(file)

	return img, nil
}
