package maps

import (
	"testing"

	"github.com/zwindler/gocastle/model"
	"github.com/zwindler/gocastle/pkg/tiles"
)

var testMap = Map{
	MapMatrix: [][]int{
		{0, 0, 0},
		{0, 13, 0}, // 13 is not walkable
	},
	MapTransitions: []SpecialTile{
		{"MapTransition", model.Coord{X: 0, Y: 0, Map: 2}, model.Coord{X: 73, Y: 2, Map: 1}},
	},
}

func TestGetMapSize(t *testing.T) {
	rows, columns := testMap.GetMapSize()

	if rows != 2 || columns != 3 {
		t.Errorf("GetMapSize() = (%d, %d); want (2, 3)", rows, columns)
	}
}

func TestGetMapImageSize(t *testing.T) {
	sizeY, sizeX := testMap.GetMapImageSize()

	if sizeX != float32(2*tiles.TileSize) || sizeY != float32(3*tiles.TileSize) {
		t.Errorf("GetMapSize() = (%f, %f); want (64, 96)", sizeX, sizeY)
	}
}

func TestCheckOutOfBounds(t *testing.T) {
	tcs := []struct {
		x, y     int
		expected bool
	}{
		{0, 0, false},
		{-1, 0, true},
		{0, 2, true},
		{4, 0, true},
		{1, 1, false},
	}

	for _, tc := range tcs {
		result := testMap.CheckOutOfBounds(tc.x, tc.y)
		if result != tc.expected {
			t.Errorf("CheckOutOfBounds(%d, %d) = %v; want %v", tc.x, tc.y, result, tc.expected)
		}
	}
}

func TestCheckTileIsWalkable(t *testing.T) {
	tcs := []struct {
		x, y     int
		expected bool
	}{
		{1, 1, false},
		{0, 0, true},
	}

	for _, tc := range tcs {
		result := testMap.CheckTileIsWalkable(tc.x, tc.y)
		if result != tc.expected {
			t.Errorf("CheckTileIsWalkable(%d, %d) = %v; want %v", tc.x, tc.y, result, tc.expected)
		}
	}
}

func TestCheckTileIsSpecial(t *testing.T) {
	tcs := []struct {
		x, y     int
		expected SpecialTile
	}{
		{1, 1, NotSpecialTile},
		{0, 0, testMap.MapTransitions[0]},
	}

	for _, tc := range tcs {
		result := testMap.CheckTileIsSpecial(tc.x, tc.y)
		if result != tc.expected {
			t.Errorf("CheckTileIsSpecial(%d, %d) = %v; want %v", tc.x, tc.y, result, tc.expected)
		}
	}
}

func TestFindObjectToRemove(t *testing.T) {
	knife, _ := model.CreateObject(model.HuntingKnife, model.Coord{X: 0, Y: 0, Map: 0})
	sword, _ := model.CreateObject(model.BluntSword, model.Coord{X: 1, Y: 1, Map: 0})

	// add a sword and a knife on the map
	testMap.ObjectList = append(testMap.ObjectList, &knife, &sword)

	// remove the knife from the map
	err := testMap.FindObjectToRemove(&knife)
	if err != nil {
		t.Errorf("failed to remove knife: %s", err)
	}

	// remove the knife from the map again
	err = testMap.FindObjectToRemove(&knife)
	if err == nil {
		t.Errorf("function FindObjectToRemove removed a knife and shouldn't have")
	}
}

func TestRemoveNPC(t *testing.T) {
	wolf1 := model.CreateNPC(model.Wolf, model.Coord{X: 0, Y: 1, Map: 0})

	// add a wolf on the map
	testMap.NPCList = append(testMap.NPCList, wolf1)

	// remove the wolf from the map
	err := testMap.RemoveNPC(wolf1)
	if err != nil {
		t.Errorf("failed to remove wolf: %s", err)
	}

	// remove the wolf from the map again
	err = testMap.RemoveNPC(wolf1)
	if err == nil {
		t.Errorf("function RemoveNPC removed a wolf and shouldn't have")
	}
}

func TestGetNPCAtPosition(t *testing.T) {
	wolf1 := model.CreateNPC(model.Wolf, model.Coord{X: 0, Y: 1, Map: 0})

	// add a wolf on the map
	testMap.NPCList = append(testMap.NPCList, wolf1)

	// Test case 1: NPC exists at the specified position
	npc1 := testMap.GetNPCAtPosition(0, 1)
	if npc1 == nil {
		t.Error("Expected NPC to exist at (0, 1), but it was not found.")
	}

	// Test case 2: NPC does not exist at the specified position
	npc2 := testMap.GetNPCAtPosition(5, 5)
	if npc2 != nil {
		t.Error("Expected no NPC at (5, 5), but one was found.")
	}
}

func TestGenerateMapImage(t *testing.T) {
	testMap.GenerateMapImage()

	if testMap.MapImage == nil {
		t.Error("Expected a non-nil map image, but it was nil.")
	}

	expectedWidth := float32(len(testMap.MapMatrix[0]) * tiles.TileSize)
	expectedHeight := float32(len(testMap.MapMatrix) * tiles.TileSize)

	imgWidth := float32(testMap.MapImage.Bounds().Dx())
	imgHeight := float32(testMap.MapImage.Bounds().Dy())

	if imgWidth != expectedWidth || imgHeight != expectedHeight {
		t.Errorf("Expected map image dimensions (%.2f, %.2f), but got (%.2f, %.2f).", expectedWidth, expectedHeight, imgWidth, imgHeight)
	}
}
