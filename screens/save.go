// screens/savescreen.go

package screens

import (
	"encoding/json"
	"fmt"
	"os"

	"gocastle/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// ShowSaveGameScreen is the main function of the save game screen
func ShowSaveGameScreen(window fyne.Window) {
	var fileNameEntry *widget.Entry

	fileNameLabel := widget.NewLabel("Enter file name:")
	fileNameEntry = widget.NewEntry()
	fileNameEntry.SetPlaceHolder("game_save.sav")

	backButton := widget.NewButton("Back", func() {
		ShowGameScreen(window)
	})

	validateButton := widget.NewButton("Validate", func() {
		if fileNameEntry.Text == "" {
			content := widget.NewLabel("Please enter a valid file name.")
			dialog.ShowCustom("Invalid File Name", "Close", content, window)
		} else {
			saveData := make(map[string]interface{})
			saveData["player"] = model.Player
			saveData["map"] = currentMap

			err := saveGameToFile(fileNameEntry.Text, saveData)
			if err != nil {
				content := widget.NewLabel(fmt.Sprintf("Error saving game: %s", err))
				dialog.ShowCustom("Error Saving Game", "Close", content, window)
			} else {
				content := widget.NewLabel("Game saved successfully.")
				dialog.ShowCustom("Game Saved", "Close", content, window)
			}
		}
	})

	content := container.NewVBox(
		fileNameLabel,
		fileNameEntry,
		layout.NewSpacer(),
		container.NewHBox(layout.NewSpacer(), backButton, validateButton, layout.NewSpacer()),
	)

	window.SetContent(content)
}

// saveGameToFile saves the game data to a JSON file.
func saveGameToFile(fileName string, data map[string]interface{}) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
