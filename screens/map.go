package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func ShowMapScreen(window fyne.Window) {
	content := container.New(layout.NewVBoxLayout())
	scrollContainer := container.NewScroll(content)
	scrollContainer.SetMinSize(fyne.NewSize(800, 600))

	window.SetContent(content)
}
