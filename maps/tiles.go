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
	X, Y     int
	filePath string
}

type TileInfo struct {
	Coordinates Coord
	IsWalkable  bool
}

var (
	TilesTypes = []TileInfo{
		{Coordinates: Coord{X: 0, Y: 0, filePath: "static/tilea2_MACK.png"}, IsWalkable: true},      //0, grass
		{Coordinates: Coord{X: 192, Y: 0, filePath: "static/tilea3_MACK.png"}, IsWalkable: false},   //1, upper left straw roof
		{Coordinates: Coord{X: 192, Y: 32, filePath: "static/tilea3_MACK.png"}, IsWalkable: false},  //2, lower left straw roof
		{Coordinates: Coord{X: 192, Y: 64, filePath: "static/tilea3_MACK.png"}, IsWalkable: false},  //3, upper left wooden wall
		{Coordinates: Coord{X: 192, Y: 96, filePath: "static/tilea3_MACK.png"}, IsWalkable: false},  //4, lower left wooden wall
		{Coordinates: Coord{X: 224, Y: 0, filePath: "static/tilea3_MACK.png"}, IsWalkable: false},   //5, upper right straw roof
		{Coordinates: Coord{X: 224, Y: 32, filePath: "static/tilea3_MACK.png"}, IsWalkable: false},  //6, lower right straw roof
		{Coordinates: Coord{X: 224, Y: 64, filePath: "static/tilea3_MACK.png"}, IsWalkable: false},  //7, upper right wooden wall
		{Coordinates: Coord{X: 224, Y: 96, filePath: "static/tilea3_MACK.png"}, IsWalkable: false},  //8, lower right wooden wall
		{Coordinates: Coord{X: 0, Y: 0, filePath: "static/transparent_tile.png"}, IsWalkable: true}, //9, transparent
		{Coordinates: Coord{X: 0, Y: 96, filePath: "static/tilea2_MACK.png"}, IsWalkable: true},     //10, some flowers
		{Coordinates: Coord{X: 32, Y: 96, filePath: "static/tilea2_MACK.png"}, IsWalkable: true},    //11, more flowers
		{Coordinates: Coord{X: 64, Y: 32, filePath: "static/tilea2_MACK.png"}, IsWalkable: false},   //12, upper left wooden border
		{Coordinates: Coord{X: 64, Y: 0, filePath: "static/tilea2_MACK.png"}, IsWalkable: false},    //13, wooden border
		{Coordinates: Coord{X: 64, Y: 64, filePath: "static/tilea2_MACK.png"}, IsWalkable: false},   //14, down left wooden border
		{Coordinates: Coord{X: 96, Y: 0, filePath: "static/tilea2_MACK.png"}, IsWalkable: false},    //15, wooden border
		{Coordinates: Coord{X: 96, Y: 32, filePath: "static/tilea2_MACK.png"}, IsWalkable: false},   //16, upper right wooden border
		{Coordinates: Coord{X: 96, Y: 64, filePath: "static/tilea2_MACK.png"}, IsWalkable: false},   //17, down right wooden border
		{Coordinates: Coord{X: 64, Y: 128, filePath: "static/tilea2_MACK.png"}, IsWalkable: true},   //18 upper left high grass
		{Coordinates: Coord{X: 64, Y: 160, filePath: "static/tilea2_MACK.png"}, IsWalkable: true},   //19 down left high grass
		{Coordinates: Coord{X: 96, Y: 128, filePath: "static/tilea2_MACK.png"}, IsWalkable: true},   //20 upper right high grass
		{Coordinates: Coord{X: 96, Y: 160, filePath: "static/tilea2_MACK.png"}, IsWalkable: true},   //21 down right high grass
		{Coordinates: Coord{X: 80, Y: 144, filePath: "static/tilea2_MACK.png"}, IsWalkable: true},   //22 center high grass
	}
)

// extractTileFromTileset extracts a subimage from coordinates on a tileset
func extractTileFromTileset(coord Coord) (image.Image, error) {
	x, y := coord.X, coord.Y
	filePath := coord.filePath

	file, err := os.Open(filePath)
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

// LoadTilesFromTileset load all tiles from []TileInfo and store them in a []image.Image
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
