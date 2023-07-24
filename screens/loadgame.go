package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ShowLoadGameScreen is the main function of the load game screen
func ShowLoadGameScreen(window fyne.Window) {
	titleLabel := widget.NewLabel("Load Game Screen")
	backButton := widget.NewButton("Back", func() {
		ShowMenuScreen(window)
	})

	content := container.NewVBox(
		titleLabel,
		backButton,
	)

	window.SetContent(content)
}
