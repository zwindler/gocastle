package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func ShowMapScreen(window fyne.Window) {
	mapContainer := container.New(layout.NewGridLayout(50))

	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			image := canvas.NewImageFromFile("./static/grass.png")
			image.FillMode = canvas.ImageFillOriginal
			image.Resize(fyne.NewSize(32, 32))
			mapContainer.Add(image)
		}
	}

	content := container.NewMax(container.NewScroll(mapContainer))

	window.SetContent(content)
}
