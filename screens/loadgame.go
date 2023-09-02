// screens/loadscreen.go

package screens

import (
	"encoding/json"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"github.com/zwindler/gocastle/maps"
	"github.com/zwindler/gocastle/model"
	"github.com/zwindler/gocastle/utils"
)

// ShowLoadGameScreen displays a file dialog to select the file to load.
func ShowLoadGameScreen(window fyne.Window) {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		if reader == nil {
			return
		}

		defer reader.Close()

		data, err := loadGameFromFile(reader)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		if err := updateLoadedGameData(data); err != nil {
			dialog.ShowError(err, window)
			return
		}

		// initialise game objects but don't reset to start
		initGame(window, false)
	}, window)
	// only show .sav files
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".sav"}))
	location, err := utils.GetBaseDirectory()
	// TODO: don't change path for iOS, Android, Flatpak
	if err != nil {
		dialog.ShowError(err, window)
	}
	fd.SetLocation(location)
	fd.Show()
}

// loadGameFromFile loads the game data from the specified JSON file.
func loadGameFromFile(r io.Reader) (data map[string]interface{}, err error) {
	return data, json.NewDecoder(r).Decode(&data)
}

// updateLoadedGameData updates the player, currentMap and TimeSinceBegin with the loaded data.
func updateLoadedGameData(data map[string]interface{}) error {
	loadedData := savedGameData{}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &loadedData)
	if err != nil {
		return err
	}

	// Assign the unmarshaled maps to the maps.AllTheMaps variable
	maps.AllTheMaps = loadedData.AllTheMaps
	player = &loadedData.Player
	model.TimeSinceBegin = loadedData.TimeSinceBegin
	/*
		// Update player
		playerData, ok := data["Player"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid player data")
		}
		if err := updatePlayerData(playerData); err != nil {
			return fmt.Errorf("failed to update player data: %w", err)
		}

		// Update AllTheMaps
		mapData, ok := data["AllTheMaps"].([]map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid map data")
		}
		if err := updateMapData(mapData); err != nil {
			return fmt.Errorf("failed to update map data: %w", err)
		}

		// Update currentMap
		timeData, ok := data["TimeSinceBegin"].(float64)
		if !ok {
			// Handle the case when the "TimeSinceBegin" key is not a float64 (or not present)
			// You can choose to show an error or set a default value, as needed.
			return fmt.Errorf("error: TimeSinceBegin is not present or not a valid float64 value")
		}
		// Convert the float64 value to int (assuming model.TimeSinceBegin is of type int)
		model.TimeSinceBegin = int(timeData)
	*/

	return nil
}

/*
// updatePlayerData updates the player data with the loaded data.
func updatePlayerData(data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, &player); err != nil {
		return err
	}
	return nil
}

// updateMapData updates the currentMap data with the loaded data.
func updateMapData(data []map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, &maps.AllTheMaps); err != nil {
		return err
	}
	return nil
}
*/
