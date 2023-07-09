package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowNewGameScreen(window fyne.Window) {
	characterNameLabel := widget.NewLabel("Character name")

	characterNameEntry := widget.NewEntry()

	firstLine := container.NewHBox(characterNameLabel, characterNameEntry)

	backButton := widget.NewButton("Back", func() {
		ShowMenuScreen(window) // Switch back to the menu screen
	})

	// TODO: Add other UI elements and logic specific to the new game screen

	content := container.NewVBox(
		firstLine,
		backButton,
		// Add other UI elements here
	)

	window.SetContent(content)
}
