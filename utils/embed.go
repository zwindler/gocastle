package utils

import (
	"embed"
	"image"
	"image/png"
	"io/fs"
	"log"
)

//go:embed static/*
var EmbeddedImages embed.FS

func GetImageFromEmbed(path string) image.Image {
	// Read the embedded image file
	closeFunc := func(f fs.File) {
		if f != nil {
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}

	file, err := EmbeddedImages.Open(path)
	if err != nil {
		// log.Fatal will cause the program to exit with an error code
		closeFunc(file)
		log.Fatal("Error opening embedded image:", err)
	}

	// Decode the image using the png package
	img, err := png.Decode(file)
	if err != nil {
		// log.Fatal will cause the program to exit with an error code
		closeFunc(file)
		log.Fatal("Error decoding embedded image:", err)
	}

	closeFunc(file)

	return img
}
