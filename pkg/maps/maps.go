package maps

import (
	"fmt"
	"image"
	"image/draw"
	"log"

	"github.com/zwindler/gocastle/pkg/coord"
	"github.com/zwindler/gocastle/pkg/npc"
	"github.com/zwindler/gocastle/pkg/object"
	"github.com/zwindler/gocastle/pkg/tiles"
)

type Map struct {
	Name           string
	NPCList        []*npc.Stats
	ObjectList     []*object.Object
	MapMatrix      [][]uint16
	MapTransitions []SpecialTile
	MapImage       image.Image
}

// This structure is used to specify tiles that have special meaning, like
// map transitions or traps.
type SpecialTile struct {
	Type        string
	Pos         coord.Coord
	Destination coord.Coord
}

var (
	NotSpecialTile = SpecialTile{"NA", coord.Coord{}, coord.Coord{}}
	Village        = Map{
		Name: "Village",
		MapTransitions: []SpecialTile{
			{"MapTransition", coord.Coord{X: 49, Y: 16, Map: 0}, coord.Coord{X: 5, Y: 6, Map: 1}},
			{"MapTransition", coord.Coord{X: 49, Y: 17, Map: 0}, coord.Coord{X: 5, Y: 6, Map: 1}},
			{"MapTransition", coord.Coord{X: 49, Y: 18, Map: 0}, coord.Coord{X: 5, Y: 6, Map: 1}},
			{"MapTransition", coord.Coord{X: 49, Y: 19, Map: 0}, coord.Coord{X: 5, Y: 6, Map: 1}},
		},
	}
	ToTheOldMine = Map{
		Name: "To the Old Mine",
		MapTransitions: []SpecialTile{
			{"MapTransition", coord.Coord{X: 0, Y: 7, Map: 1}, coord.Coord{X: 47, Y: 17, Map: 0}},
			{"MapTransition", coord.Coord{X: 0, Y: 8, Map: 1}, coord.Coord{X: 47, Y: 17, Map: 0}},
			{"MapTransition", coord.Coord{X: 0, Y: 9, Map: 1}, coord.Coord{X: 47, Y: 17, Map: 0}},
			{"MapTransition", coord.Coord{X: 0, Y: 10, Map: 1}, coord.Coord{X: 47, Y: 17, Map: 0}},
			{"MapTransition", coord.Coord{X: 63, Y: 0, Map: 1}, coord.Coord{X: 1, Y: 1, Map: 2}},
			{"MapTransition", coord.Coord{X: 63, Y: 1, Map: 1}, coord.Coord{X: 1, Y: 1, Map: 2}},
		},
	}
	TheOldMine = Map{
		Name: "The Old Mine",
		MapTransitions: []SpecialTile{
			{"MapTransition", coord.Coord{X: 0, Y: 0, Map: 2}, coord.Coord{X: 63, Y: 2, Map: 1}},
		},
	}

	// Slice containing all the maps of the game.
	AllTheMaps = []Map{Village, ToTheOldMine, TheOldMine}
)

// GetMapSize return number of rows and number of columns of a given map.
func (currentMap *Map) GetMapSize() (mapRows, mapColumns int) {
	return currentMap.getRows(), currentMap.getColumns()
}

// getRows returns the number of rows of a given map.
func (currentMap *Map) getRows() int {
	return len(currentMap.MapMatrix)
}

// getColumns returns the number of columns of a given map.
func (currentMap *Map) getColumns() int {
	if currentMap.getRows() == 0 {
		return 0
	}
	return len(currentMap.MapMatrix[0])
}

// GetMapImageSize returns the image x and y size.
func (currentMap *Map) GetMapImageSize() (float32, float32) {
	return currentMap.getMapImageSizeX(), currentMap.getMapImageSizeY()
}

// getMapImageSizeX return x as Map Image size.
func (currentMap *Map) getMapImageSizeX() float32 {
	return float32(tiles.TileSize * currentMap.getColumns())
}

// getMapImageSizeY return y as Map Image size.
func (currentMap *Map) getMapImageSizeY() float32 {
	return float32(tiles.TileSize * currentMap.getRows())
}

// CheckOutOfBounds checks if x, y coordinates are out of map bounds.
func (currentMap *Map) CheckOutOfBounds(futurePosX, futurePosY int) bool {
	mapRows, mapColumns := currentMap.GetMapSize()
	if futurePosX >= 0 && futurePosX < mapColumns &&
		futurePosY >= 0 && futurePosY < mapRows {
		return false
	}
	return true
}

// CheckTileIsWalkable checks if, for a given map, x,y coordinates are considered walkable.
func (currentMap *Map) CheckTileIsWalkable(futurePosX, futurePosY int) bool {
	return tiles.TilesTypes[currentMap.MapMatrix[futurePosY][futurePosX]].IsWalkable
}

// CheckTileIsSpecial checks if, for a given map, x,y coordinates are special
// If so, return the SpecialTile do deal with effect.
func (currentMap *Map) CheckTileIsSpecial(posX, posY int) SpecialTile {
	// for now, only deal with map transitions
	for _, tile := range currentMap.MapTransitions {
		if tile.Pos.X == posX && tile.Pos.Y == posY {
			return tile
		}
	}
	return NotSpecialTile
}

// FindObjectToRemove loops through the currentMap ObjectList and removes object *model.Object.
func (currentMap *Map) FindObjectToRemove(object *object.Object) error {
	indexToRemove := -1
	for i, obj := range currentMap.ObjectList {
		if obj == object {
			indexToRemove = i
			break
		}
	}

	// If the object was found, remove it from the slice
	if indexToRemove >= 0 {
		currentMap.ObjectList = append(currentMap.ObjectList[:indexToRemove], currentMap.ObjectList[indexToRemove+1:]...)
		return nil
	}

	return fmt.Errorf("unable to find object to remove")
}

// For a given map, remove NPC by list id and hide CanvasImage.
func (currentMap *Map) RemoveNPC(npcToRemove *npc.Stats) error {
	var indexToRemove int = -1
	for i, npc := range currentMap.NPCList {
		if npc == npcToRemove {
			indexToRemove = i
			break
		}
	}

	// If the npc was found, remove it from the slice
	if indexToRemove >= 0 {
		// remove NPC image from fyne map
		currentMap.NPCList[indexToRemove].Avatar.CanvasImage.Hidden = true
		// remove NPC from NPCList slice
		currentMap.NPCList = append(currentMap.NPCList[:indexToRemove], currentMap.NPCList[indexToRemove+1:]...)
		return nil
	}

	return fmt.Errorf("unable to find NPC to remove")
}

// For a given NPCsOnCurrentMap, check if NPCs are located on x,y
// return nil if none or pointer to npc.
func (currentMap *Map) GetNPCAtPosition(x, y int) *npc.Stats {
	// find if a NPC matches our destination
	for _, npc := range currentMap.NPCList {
		if npc.Avatar.Coord.X == x && npc.Avatar.Coord.Y == y {
			return npc
		}
	}
	return nil
}

// GenerateMapImage generates or regenerates the whole image from map tiles.
func (currentMap *Map) GenerateMapImage() {
	numRows, numColumns := currentMap.GetMapSize()
	xSize, ySize := currentMap.GetMapImageSize()

	// extract the needed tiles from the Tileset
	// create a table of images (image.Image type)
	loadedTiles, err := tiles.LoadTilesFromTileset(tiles.TilesTypes)
	if err != nil {
		err = fmt.Errorf("unable to load tile from Tileset: %w", err)
		log.Fatalf("MapMatrix error: %s", err)
	}

	// now, reconstruct the whole map image with tiles images
	fullImage := image.NewRGBA(image.Rect(0, 0, int(xSize), int(ySize)))
	for row := 0; row < numRows; row++ {
		currentRowImage := image.NewRGBA(image.Rect(0, 0, int(xSize), coord.TileSize))
		for column := 0; column < numColumns; column++ {
			currentImage := loadedTiles[currentMap.MapMatrix[row][column]]
			startingPosition := image.Point{column * coord.TileSize, 0}
			currentTileRectangle := image.Rectangle{startingPosition, startingPosition.Add(image.Point{coord.TileSize, coord.TileSize})}
			draw.Draw(currentRowImage, currentTileRectangle.Bounds(), currentImage, image.Point{0, 0}, draw.Src)
		}
		// we have reconstructed the whole row with all the tiles
		// now, we can add the row to the full image
		startingRowPosition := image.Point{0, row * coord.TileSize}
		currentRowRectangle := image.Rectangle{startingRowPosition, startingRowPosition.Add(image.Point{coord.TileSize * numColumns, coord.TileSize})}
		draw.Draw(fullImage, currentRowRectangle.Bounds(), currentRowImage, image.Point{0, 0}, draw.Src)
	}

	currentMap.MapImage = fullImage
}

func (m *Map) Copy() *Map {
	copyMap := &Map{
		Name: m.Name,
	}
	copyMap.NPCList = make([]*npc.Stats, len(m.NPCList))
	for i, npc := range m.NPCList {
		copyMap.NPCList[i] = npc.Copy()
	}
	copyMap.ObjectList = make([]*object.Object, len(m.ObjectList))
	for i, obj := range m.ObjectList {
		copyMap.ObjectList[i] = obj.Copy()
	}
	copyMap.MapMatrix = make([][]uint16, len(m.MapMatrix))
	for i, row := range m.MapMatrix {
		copyMap.MapMatrix[i] = make([]uint16, len(row))
		copy(copyMap.MapMatrix[i], row)
	}
	copyMap.MapTransitions = make([]SpecialTile, len(m.MapTransitions))
	copy(copyMap.MapTransitions, m.MapTransitions)

	return copyMap
}
