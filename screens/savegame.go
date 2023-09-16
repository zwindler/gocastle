// screens/savescreen.go

package screens

import (
	"encoding/json"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"github.com/zwindler/gocastle/maps"
	"github.com/zwindler/gocastle/model"
	"github.com/zwindler/gocastle/pkg/timespent"
	"github.com/zwindler/gocastle/utils"
)

type savedGameData struct {
	Player         model.CharacterStats
	AllTheMaps     []maps.Map
	TimeSinceBegin int
}

// ShowSaveGameScreen is the main function of the save game screen.
func ShowSaveGameScreen(window fyne.Window) {
	// Remove Images from character & inventory before saving
	playerSaveData := *player
	playerSaveData.Avatar.CanvasImage.Image = nil
	playerSaveData.Avatar.ObjectInMapContainer = nil

	for index := range playerSaveData.Inventory {
		playerSaveData.Inventory[index].CanvasImage = nil
	}

	// Remove Images from Maps, NPCs & Objects before saving
	// Also remove ObjectInMapContainer variable which seem to come from fyne?
	mapSaveData := maps.AllTheMaps
	for indexMap := range mapSaveData {
		mapSaveData[indexMap].MapImage = nil
		for index := range mapSaveData[indexMap].NPCList {
			mapSaveData[indexMap].NPCList[index].Avatar.CanvasImage.Image = nil
			mapSaveData[indexMap].NPCList[index].Avatar.ObjectInMapContainer = nil
		}
		for index := range mapSaveData[indexMap].ObjectList {
			mapSaveData[indexMap].ObjectList[index].CanvasImage.Image = nil
		}
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
	location, err := utils.GetBaseDirectory()
	if err != nil {
		dialog.ShowError(err, window)
	}
	fd.SetLocation(location)
	fd.Show()
}
