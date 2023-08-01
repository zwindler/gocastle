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
		player.AddObjectToInventory(model.BareHands)
		knife, err := model.CreateObject(model.HuntingKnife)
		if err != nil {
			err = fmt.Errorf("unable to add knife to inventory: %w", err)
			log.Fatalf("NewGame error: %s", err)
		}
		knifeIndex := player.AddObjectToInventory(knife)
		player.EquipItem(knifeIndex)
		player.ChangeGold(10)

		// set coordinates to "Village" map starting coordinates
		X, Y = currentMap.PlayerStart.X, currentMap.PlayerStart.Y
	} else {
		X, Y = player.Avatar.PosX, player.Avatar.PosY
	}
	player.Avatar = model.CreateAvatar(player.Avatar, X, Y)

	// create NPCs avatars
	currentMap.AddNPCs()

	ShowGameScreen(window)
}
