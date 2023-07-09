package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowLoadGameScreen(window fyne.Window) {
	titleLabel := widget.NewLabel("Load Game Screen")
	backButton := widget.NewButton("Back", func() {
		ShowMenuScreen(window) // Switch back to the menu screen
	})

	// TODO: Add other UI elements and logic specific to the new game screen

	content := container.NewVBox(
		titleLabel,
		backButton,
		// Add other UI elements here
	)

	window.SetContent(content)
}
