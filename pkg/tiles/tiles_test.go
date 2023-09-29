package tiles

import (
	"fmt"
	"testing"
)

func TestLoadTilesFromTileset(t *testing.T) {
	testTiles := []TileInfo{
		{X: 0, Y: 0, filePath: "static/tilea2_MACK.png", IsWalkable: true},
		{X: 192, Y: 0, filePath: "static/tilea3_MACK.png", IsWalkable: false},
	}

	// Call the LoadTilesFromTileset function
	images, err := LoadTilesFromTileset(testTiles)
	if err != nil {
		t.Errorf("LoadTilesFromTileset returned an error: %v", err)
	}

	// Verify that the number of loaded images matches the number of testTiles
	if len(images) != len(testTiles) {
		t.Errorf("Expected %d images, but got %d", len(testTiles), len(images))
	}
}

func TestExtractTileFromTileset(t *testing.T) {
	testCases := []struct {
		x        int
		y        int
		filePath string
	}{
		{x: 0, y: 0, filePath: "static/tilea2_MACK.png"},
		{x: 192, y: 0, filePath: "static/tilea3_MACK.png"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("x=%d_y=%d", tc.x, tc.y), func(t *testing.T) {
			image, err := extractTileFromTileset(tc.x, tc.y, tc.filePath)
			if err != nil {
				t.Errorf("extractTileFromTileset returned an error: %v", err)
			}

			// Verify that the loaded image is not nil
			if image == nil {
				t.Error("extractTileFromTileset returned a nil image")
			}
		})
	}
}
