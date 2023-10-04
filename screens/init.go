package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"

	"github.com/zwindler/gocastle/model"
	"github.com/zwindler/gocastle/pkg/avatar"
	"github.com/zwindler/gocastle/pkg/coord"
	"github.com/zwindler/gocastle/pkg/embed"
	"github.com/zwindler/gocastle/pkg/maps"
	"github.com/zwindler/gocastle/pkg/npc"
)

var (
	player     = &model.Player
	currentMap = &maps.AllTheMaps[0]
)

// initGame will initialise all needed variables before start game (start=true) or load game (start=false).
func initGame(window fyne.Window, start bool) {
	// refresh player stats (heal or not depending on "start")
	player.RefreshStats(start)

	// init categories
	model.InitializeCategories()

	// create player Avatar
	if start {
		// load all pregenerated maps from json
		for i := 0; i < len(maps.AllTheMaps); i++ {
			thisMapMatrix, err := embed.GetMapMatrixFromEmbed(fmt.Sprintf("maps/%d.json", i))
			if err != nil {
				dialog.ShowError(err, window)
			}
			maps.AllTheMaps[i].MapMatrix = thisMapMatrix
		}

		player.ChangeGold(10)

		// TODO rework this
		// Map0 Village
		knife, _ := model.CreateObject(model.HuntingKnife, coord.Coord{X: 10, Y: 10, Map: 0})
		maps.AllTheMaps[0].ObjectList = append(maps.AllTheMaps[0].ObjectList, &knife)
		farmer := npc.Spawn(model.FemaleFarmer, coord.Coord{X: 10, Y: 15, Map: 0})
		ant1 := npc.Spawn(model.GiantAnt, coord.Coord{X: 5, Y: 34, Map: 0})
		ant2 := npc.Spawn(model.GiantAnt, coord.Coord{X: 6, Y: 32, Map: 0})
		ant3 := npc.Spawn(model.GiantAnt, coord.Coord{X: 7, Y: 33, Map: 0})
		maps.AllTheMaps[0].NPCList = append(maps.AllTheMaps[0].NPCList, farmer, ant1, ant2, ant3)

		// Map1 To The Old Mine
		sword, _ := model.CreateObject(model.BluntSword, coord.Coord{X: 9, Y: 4, Map: 1})
		maps.AllTheMaps[1].ObjectList = append(maps.AllTheMaps[1].ObjectList, &sword)
		wolf1 := npc.Spawn(model.Wolf, coord.Coord{X: 70, Y: 24, Map: 1})
		wolf2 := npc.Spawn(model.Wolf, coord.Coord{X: 69, Y: 23, Map: 1})
		maps.AllTheMaps[1].NPCList = append(maps.AllTheMaps[1].NPCList, wolf1, wolf2)

		player.Avatar.Coord = coord.Coord{X: 15, Y: 15, Map: 0}
	}

	// pregenerate the map image to save time in game screen
	currentMap.GenerateMapImage()

	currentMap = &maps.AllTheMaps[player.Avatar.Coord.Map]
	player.Avatar = avatar.Spawn(player.Avatar, player.Avatar.Coord)

	ShowGameScreen(window)
}
