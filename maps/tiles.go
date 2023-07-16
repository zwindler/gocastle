package maps

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg" // Import image/jpeg to support JPEG images
	_ "image/png"  // Import image/png to support PNG images
	"os"
)

type Coord struct {
	X, Y int
}

type TileInfo struct {
	Coordinates Coord
	IsWalkable  bool
}

var (
	TilesTypes = []TileInfo{
		{Coordinates: Coord{X: 0, Y: 64}, IsWalkable: true},     //0
		{Coordinates: Coord{X: 576, Y: 96}, IsWalkable: false},  //1
		{Coordinates: Coord{X: 576, Y: 128}, IsWalkable: false}, //2
		{Coordinates: Coord{X: 608, Y: 128}, IsWalkable: false}, //3
		{Coordinates: Coord{X: 608, Y: 96}, IsWalkable: false},  //4
		{Coordinates: Coord{X: 512, Y: 160}, IsWalkable: false}, //5
		{Coordinates: Coord{X: 480, Y: 128}, IsWalkable: false}, //6
		{Coordinates: Coord{X: 512, Y: 96}, IsWalkable: false},  //7
		{Coordinates: Coord{X: 544, Y: 128}, IsWalkable: false}, //8
		{Coordinates: Coord{X: 128, Y: 64}, IsWalkable: false},  //9
		{Coordinates: Coord{X: 96, Y: 64}, IsWalkable: false},   //10
	}
)

func extractTileFromTileset(coord Coord) (image.Image, error) {
	x, y := coord.X, coord.Y
	file, err := os.Open("static/RPG Nature Tileset.png")
	if err != nil {
		fmt.Println("Error opening image:", err)
		return nil, err
	}
	defer file.Close()

	bigImage, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil, err
	}

	width := 32
	height := 32

	partImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(partImage, partImage.Bounds(), bigImage, image.Point{x, y}, draw.Src)

	return partImage, nil
}

func LoadTilesFromTileset(tiles []TileInfo) ([]image.Image, error) {
	var images []image.Image

	for _, tile := range tiles {
		image, err := extractTileFromTileset(tile.Coordinates)
		if err != nil {
			return nil, fmt.Errorf("unable to load tile from Tileset: %w", err)
		}
		images = append(images, image)
	}

	return images, nil
}
