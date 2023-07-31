package utils

import (
	"embed"
	"image"
	"image/png"
	"log"
)

//go:embed static/*
var EmbeddedImages embed.FS

func GetImageFromEmbed(path string) image.Image {
	// Read the embedded image file
	file, err := EmbeddedImages.Open(path)
	if err != nil {
		log.Fatal("Error opening embedded image:", err)
	}
	defer file.Close()

	// Decode the image using the png package
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal("Error decoding embedded image:", err)
	}

	return img
}
