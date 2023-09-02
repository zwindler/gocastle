package maps

import (
	"github.com/zwindler/gocastle/model"
)

type Coord struct {
	X, Y int
	Map  int
}

type Map struct {
	Name           string
	PlayerStart    Coord
	NPCList        []*model.NPCStats
	ObjectList     []*model.Object
	MapMatrix      [][]int
	MapTransitions []SpecialTile
}

// This structure is used to specify tiles that have special meaning, like
// map transitions or traps.
type SpecialTile struct {
	Type        string
	Pos         Coord
	Destination Coord
}

var NotSpecialTile = SpecialTile{"NA", Coord{}, Coord{}}

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
	return TilesTypes[currentMap.MapMatrix[futurePosY][futurePosX]].IsWalkable
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
func (currentMap *Map) FindObjectToRemove(object *model.Object) {
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
	}
}

// For a given map, remove NPC by list id and hide CanvasImage.
func (currentMap *Map) RemoveNPC(npcToRemove *model.NPCStats) {
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
	}
}

// For a given NPCsOnCurrentMap, check if NPCs are located on x,y
// return nil if none or pointer to npc.
func (currentMap *Map) GetNPCAtPosition(x, y int) *model.NPCStats {
	// find if a NPC matches our destination
	for _, npc := range currentMap.NPCList {
		if npc.Avatar.PosX == x && npc.Avatar.PosY == y {
			return npc
		}
	}
	return nil
}
