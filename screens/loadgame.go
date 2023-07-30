// screens/loadscreen.go

package screens

import (
	"encoding/json"
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

// ShowLoadGameScreen displays a file dialog to select the file to load.
func ShowLoadGameScreen(window fyne.Window) error {
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
		updateLoadedGameData(data)
		ShowGameScreen(window)
	}, window)
	// only show .json files
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".sav"}))
	fd.Show()

	return nil
}

// loadGameFromFile loads the game data from the specified JSON file.
func loadGameFromFile(r io.Reader) (map[string]interface{}, error) {
	var data map[string]interface{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// updateLoadedGameData updates the player and currentMap with the loaded data.
func updateLoadedGameData(data map[string]interface{}) error {
	// Update player
	playerData, ok := data["player"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid player data")
	}
	if err := updatePlayerData(playerData); err != nil {
		return fmt.Errorf("failed to update player data: %w", err)
	}

	// Update currentMap
	mapData, ok := data["map"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid map data")
	}
	if err := updateMapData(mapData); err != nil {
		return fmt.Errorf("failed to update map data: %w", err)
	}

	return nil
}

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
func updateMapData(data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, &currentMap); err != nil {
		return err
	}
	return nil
}
