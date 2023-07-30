// screens/savescreen.go

package screens

import (
	"encoding/json"
	"gocastle/maps"
	"gocastle/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// ShowSaveGameScreen is the main function of the save game screen
func ShowSaveGameScreen(window fyne.Window) {
	// Get the data to save
	gameData := struct {
		Player         model.CharacterStats
		CurrentMap     maps.Map
		TimeSinceBegin int
	}{
		Player:         *player,
		CurrentMap:     currentMap,
		TimeSinceBegin: model.TimeSinceBegin,
	}

	// Show file save dialog
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err == nil && writer != nil {
			defer writer.Close()

			// Create JSON encoder
			encoder := json.NewEncoder(writer)

			// Write game data to JSON file
			if err := encoder.Encode(gameData); err != nil {
				dialog.ShowError(err, window)
			} else {
				dialog.ShowInformation("Game Saved", "Game data has been successfully saved.", window)
			}
		}
	}, window)
}
