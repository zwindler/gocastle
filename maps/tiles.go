package maps

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/zwindler/gocastle/utils"

	_ "image/jpeg" // Import image/jpeg to support JPEG images
	_ "image/png"  // Import image/png to support PNG images
)

type TileInfo struct {
	X, Y       int
	filePath   string
	IsWalkable bool
}

var TilesTypes = []TileInfo{
	{X: 0, Y: 0, filePath: "static/tilea2_MACK.png", IsWalkable: true},      // 0, grass
	{X: 192, Y: 0, filePath: "static/tilea3_MACK.png", IsWalkable: false},   // 1, upper left straw roof
	{X: 192, Y: 32, filePath: "static/tilea3_MACK.png", IsWalkable: false},  // 2, lower left straw roof
	{X: 192, Y: 64, filePath: "static/tilea3_MACK.png", IsWalkable: false},  // 3, upper left wooden wall
	{X: 192, Y: 96, filePath: "static/tilea3_MACK.png", IsWalkable: false},  // 4, lower left wooden wall
	{X: 224, Y: 0, filePath: "static/tilea3_MACK.png", IsWalkable: false},   // 5, upper right straw roof
	{X: 224, Y: 32, filePath: "static/tilea3_MACK.png", IsWalkable: false},  // 6, lower right straw roof
	{X: 224, Y: 64, filePath: "static/tilea3_MACK.png", IsWalkable: false},  // 7, upper right wooden wall
	{X: 224, Y: 96, filePath: "static/tilea3_MACK.png", IsWalkable: false},  // 8, lower right wooden wall
	{X: 0, Y: 0, filePath: "static/transparent_tile.png", IsWalkable: true}, // 9, transparent
	{X: 0, Y: 96, filePath: "static/tilea2_MACK.png", IsWalkable: true},     // 10, some flowers
	{X: 32, Y: 96, filePath: "static/tilea2_MACK.png", IsWalkable: true},    // 11, more flowers
	{X: 64, Y: 32, filePath: "static/tilea2_MACK.png", IsWalkable: false},   // 12, upper left wooden border
	{X: 64, Y: 0, filePath: "static/tilea2_MACK.png", IsWalkable: false},    // 13, wooden border
	{X: 64, Y: 64, filePath: "static/tilea2_MACK.png", IsWalkable: false},   // 14, down left wooden border
	{X: 96, Y: 0, filePath: "static/tilea2_MACK.png", IsWalkable: false},    // 15, wooden border
	{X: 96, Y: 32, filePath: "static/tilea2_MACK.png", IsWalkable: false},   // 16, upper right wooden border
	{X: 96, Y: 64, filePath: "static/tilea2_MACK.png", IsWalkable: false},   // 17, down right wooden border
	{X: 64, Y: 128, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 18 upper left high grass
	{X: 64, Y: 160, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 19 down left high grass
	{X: 96, Y: 128, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 20 upper right high grass
	{X: 96, Y: 160, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 21 down right high grass
	{X: 80, Y: 144, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 22 center high grass
	{X: 224, Y: 0, filePath: "static/tilea2_MACK.png", IsWalkable: true},    // 23 sand (one tile)
	{X: 192, Y: 32, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 24 sand (with grass upper left)
	{X: 224, Y: 32, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 25 sand (with grass upper right)
	{X: 192, Y: 64, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 26 sand (with grass lower left)
	{X: 224, Y: 64, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 27 sand (with grass lower right)
	{X: 128, Y: 64, filePath: "static/tilea4_MACK.png", IsWalkable: false},  // 28 tree top left
	{X: 128, Y: 96, filePath: "static/tilea4_MACK.png", IsWalkable: false},  // 29 tree middle left
	{X: 128, Y: 128, filePath: "static/tilea4_MACK.png", IsWalkable: false}, // 30 tree bottom left
	{X: 160, Y: 64, filePath: "static/tilea4_MACK.png", IsWalkable: false},  // 31 tree top right
	{X: 160, Y: 96, filePath: "static/tilea4_MACK.png", IsWalkable: false},  // 32 tree middle right
	{X: 160, Y: 128, filePath: "static/tilea4_MACK.png", IsWalkable: false}, // 33 tree bottom right
	{X: 128, Y: 32, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 34 stone upper left
	{X: 128, Y: 64, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 35 stone bottom left
	{X: 160, Y: 32, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 36 stone upper right
	{X: 160, Y: 64, filePath: "static/tilea2_MACK.png", IsWalkable: true},   // 37 stone bottom right
	{X: 32, Y: 256, filePath: "static/tilee_MACK.png", IsWalkable: true},    // 38 upper cave entrance
	{X: 32, Y: 288, filePath: "static/tilee_MACK.png", IsWalkable: true},    // 39 lower cave entrance
	{X: 192, Y: 416, filePath: "static/tilea4_MACK.png", IsWalkable: false}, // 40 montain upper left
	{X: 192, Y: 448, filePath: "static/tilea4_MACK.png", IsWalkable: false}, // 41 montain bottom left
	{X: 224, Y: 416, filePath: "static/tilea4_MACK.png", IsWalkable: false}, // 42 montain upper right
	{X: 224, Y: 448, filePath: "static/tilea4_MACK.png", IsWalkable: false}, // 43 montain bottom right
	{X: 144, Y: 64, filePath: "static/tilea4_MACK.png", IsWalkable: false},  // 44 tree top center
	{X: 144, Y: 96, filePath: "static/tilea4_MACK.png", IsWalkable: false},  // 45 tree middle center
	{X: 144, Y: 128, filePath: "static/tilea4_MACK.png", IsWalkable: false}, // 46 tree bottom center

}

// extractTileFromTileset extracts a subimage from coordinates on a tileset.
func extractTileFromTileset(x, y int, filePath string) (image.Image, error) {
	file, err := utils.EmbeddedImages.Open(filePath)
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

// LoadTilesFromTileset load all tiles from []TileInfo and store them in a []image.Image.
func LoadTilesFromTileset(tiles []TileInfo) ([]image.Image, error) {
	var images []image.Image

	for _, tile := range tiles {
		image, err := extractTileFromTileset(tile.X, tile.Y, tile.filePath)
		if err != nil {
			return nil, fmt.Errorf("unable to load tile from Tileset: %w", err)
		}
		images = append(images, image)
	}

	return images, nil
}
