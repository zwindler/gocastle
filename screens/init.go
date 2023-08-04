package screens

import (
	"fmt"
	"gocastle/model"
	"log"

	"fyne.io/fyne/v2"
)

// initGame will initialise all needed variables before start game (start=true) or load game (start=false)
func initGame(window fyne.Window, start bool) {
	var X, Y int

	// refresh player stats (heal or not depending on "start")
	player.RefreshStats(start)

	// init categories
	model.InitializeCategories()

	// create player Avatar
	if start {
		// create a knife, add it to player's inventory, equip it
		// TODO rework later
		player.AddObjectToInventory(model.BareHands, false)
		knife, err := model.CreateObject(model.HuntingKnife, 0, 0)
		if err != nil {
			err = fmt.Errorf("unable to create knife: %w", err)
			log.Fatalf("NewGame error: %s", err)
		}
		player.AddObjectToInventory(knife, true)
		player.ChangeGold(10)

		// set coordinates to "Village" map starting coordinates
		X, Y = currentMap.PlayerStart.X, currentMap.PlayerStart.Y
	} else {
		// we are loading game, set position to current position
		X, Y = player.Avatar.PosX, player.Avatar.PosY
	}
	player.Avatar = model.CreateAvatar(player.Avatar, X, Y)

	// create NPCs avatars
	currentMap.AddNPCs()

	ShowGameScreen(window)
}
