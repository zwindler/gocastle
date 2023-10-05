// screens/savescreen.go

package screens

import (
	"encoding/json"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"github.com/zwindler/gocastle/pkg/character"
	"github.com/zwindler/gocastle/pkg/game"
	"github.com/zwindler/gocastle/pkg/maps"
	"github.com/zwindler/gocastle/pkg/timespent"
)

type savedGameData struct {
	Player         character.Stats
	AllTheMaps     []maps.Map
	TimeSinceBegin int
}

// ShowSaveGameScreen is the main function of the save game screen.
func ShowSaveGameScreen(window fyne.Window) {
	playerSaveData := game.Player

	// Create copies of maps.AllTheMaps without images
	mapSaveData := make([]maps.Map, len(maps.AllTheMaps))
	for indexMap := range maps.AllTheMaps {
		mapSaveData[indexMap] = *maps.AllTheMaps[indexMap].Copy()

		for index := range mapSaveData[indexMap].NPCList {
			npc := *maps.AllTheMaps[indexMap].NPCList[index]
			mapSaveData[indexMap].NPCList[index].Avatar = npc.Avatar.Copy()
		}

		for index := range mapSaveData[indexMap].ObjectList {
			mapSaveData[indexMap].ObjectList[index] = mapSaveData[indexMap].ObjectList[index].Copy()
		}
	}
	// Remove Images from character & inventory before saving
	playerSaveData.Avatar.CanvasImage.Image = nil
	playerSaveData.Avatar.ObjectInMapContainer = nil

	for index := range playerSaveData.Inventory {
		playerSaveData.Inventory[index].CanvasImage = nil
	}

	// Get the data to save
	gameData := savedGameData{
		Player:         playerSaveData,
		AllTheMaps:     mapSaveData,
		TimeSinceBegin: timespent.Get(),
	}

	// Show file save dialog
	fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err == nil && writer != nil {
			defer writer.Close()

			// Create JSON encoder
			encoder := json.NewEncoder(writer)

			// Write game data to JSON .sav file
			if err := encoder.Encode(gameData); err != nil {
				dialog.ShowError(err, window)
			} else {
				dialog.ShowInformation("Game Saved", "Game data has been successfully saved.", window)
			}
		}
	}, window)
	// only allow .sav files
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".sav"}))
	fd.SetFileName("backup.sav")
	// TODO: don't change path for iOS, Android, Flatpak
	location, err := getBaseDirectory()
	if err != nil {
		dialog.ShowError(err, window)
	}
	fd.SetLocation(location)
	fd.Show()
}
