package coord

import "fyne.io/fyne/v2"

type Coord struct {
	X, Y int
	Map  int
}

const (
	TileSize = 32
)

var FyneTileSize = fyne.NewSize(TileSize, TileSize)
