package screens

import (
	"fyne.io/fyne/v2"

	"github.com/zwindler/gocastle/maps"
	"github.com/zwindler/gocastle/model"
)

var (
	player     = &model.Player
	currentMap = &maps.AllTheMaps[0]
)

// initGame will initialise all needed variables before start game (start=true) or load game (start=false).
func initGame(window fyne.Window, start bool) {
	var X, Y int

	// refresh player stats (heal or not depending on "start")
	player.RefreshStats(start)

	// init categories
	model.InitializeCategories()

	// create player Avatar
	if start {
		player.ChangeGold(10)

		// TODO rework this
		knife, _ := model.CreateObject(model.HuntingKnife, 10, 10)
		sword, _ := model.CreateObject(model.BluntSword, 20, 20)
		maps.AllTheMaps[0].ObjectList = append(maps.AllTheMaps[0].ObjectList, &knife, &sword)
		farmer := model.CreateNPC(model.FemaleFarmer, 10, 15)
		wolf1 := model.CreateNPC(model.Wolf, 25, 26)
		wolf2 := model.CreateNPC(model.Wolf, 28, 27)
		ogre := model.CreateNPC(model.Ogre, 30, 25)
		maps.AllTheMaps[0].NPCList = append(maps.AllTheMaps[0].NPCList, farmer, wolf1, wolf2, ogre)

		// set coordinates to "Village" map starting coordinates
		X, Y = currentMap.PlayerStart.X, currentMap.PlayerStart.Y
	} else {
		// we are loading game, set position to current position
		X, Y = player.Avatar.PosX, player.Avatar.PosY
	}
	player.Avatar = model.CreateAvatar(player.Avatar, X, Y)

	ShowGameScreen(window)
}
