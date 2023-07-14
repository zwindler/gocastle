package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowMapScreen(window fyne.Window) {
	characterNameLabel := widget.NewLabel("Character's name")
	content := container.NewVBox(characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel, characterNameLabel)
	scrollContainer := container.NewScroll(content)
	scrollContainer.SetMinSize(fyne.NewSize(800, 600))

	window.SetContent(content)
	window.Resize(fyne.NewSize(800, 600))
}
