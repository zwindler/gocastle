// screens/loadscreen.go

package screens

import (
	"encoding/json"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"github.com/zwindler/gocastle/maps"
	"github.com/zwindler/gocastle/pkg/timespent"
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
	location, err := getBaseDirectory()
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

	player = &loadedData.Player

	// Assign the unmarshaled maps to the maps.AllTheMaps variable
	maps.AllTheMaps = loadedData.AllTheMaps
	// NPCs and Objects were saved without their Image, refresh it
	for indexMap := range maps.AllTheMaps {
		for _, npc := range maps.AllTheMaps[indexMap].NPCList {
			npc.Avatar.RefreshAvatar()
		}
		for _, object := range maps.AllTheMaps[indexMap].ObjectList {
			object.RefreshObject()
		}
	}

	timespent.Set(loadedData.TimeSinceBegin)

	return nil
}
