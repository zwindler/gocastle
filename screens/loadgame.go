// screens/loadscreen.go

package screens

import (
	"encoding/json"
	"fmt"
	"os"

	"gocastle/maps"
	"gocastle/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// ShowLoadGameScreen is the main function of the load game screen
func ShowLoadGameScreen(window fyne.Window) {
	titleLabel := widget.NewLabel("Load Game Screen")
	backButton := widget.NewButton("Back", func() {
		ShowMenuScreen(window)
	})

	loadButton := widget.NewButton("Load Game", func() {
		fileName, err := showLoadFileDialog(window)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		data, err := loadGameFromFile(fileName)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		// Update player and currentMap with the loaded data
		if err := updateLoadedGameData(data); err != nil {
			dialog.ShowError(err, window)
			return
		}

		dialog.ShowInformation("Game Loaded", "Game loaded successfully!", window)
		ShowGameScreen(window)
	})

	content := container.NewVBox(
		titleLabel,
		backButton,
		loadButton,
	)

	window.SetContent(content)
}

// showLoadFileDialog displays a file dialog to select the file to load.
func showLoadFileDialog(window fyne.Window) (string, error) {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		if reader == nil {
			return
		}

		defer reader.Close()
	}, window)
	// only show .json files
	fd.SetFilter(storageFilter())
	fd.Show()

	return fd.FilePath(), nil
}

// storageFilter filters files to show only JSON files for loading.
func storageFilter() storage.FileFilter {
	return storage.NewExtensionFileFilter([]string{".json"})
}

// loadGameFromFile loads the game data from the specified JSON file.
func loadGameFromFile(fileName string) (map[string]interface{}, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data map[string]interface{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
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
	mapData, ok := data["village"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid village data")
	}
	if err := updateMapData(mapData); err != nil {
		return fmt.Errorf("failed to update village data: %w", err)
	}

	return nil
}

// updatePlayerData updates the player data with the loaded data.
func updatePlayerData(data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, &model.Player); err != nil {
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
	if err := json.Unmarshal(jsonData, &maps.Village); err != nil {
		return err
	}
	return nil
}
