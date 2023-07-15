package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func ShowMapScreen(window fyne.Window) {
	h, v := 20, 20

	imageMatrix := createMapMatrix(h, v)
	player := canvas.NewImageFromFile("./static/warrior.png")
	player.FillMode = canvas.ImageFillOriginal
	player.Resize(fyne.NewSize(32, 32))

	mapContainer := container.NewWithoutLayout()

	for i := 0; i < v; i++ {
		for j := 0; j < h; j++ {
			tile := container.NewWithoutLayout(
				imageMatrix[i][j],
			)
			if i == 1 && j == 1 {
				tile.Add(player)
			}
			tile.Resize(fyne.NewSize(32, 32))
			tile.Move(fyne.NewPos(float32(i)*32, float32(j)*32))
			mapContainer.Add(tile)
		}
	}

	content := container.NewScroll(mapContainer)

	window.SetContent(content)
}

func createMapMatrix(h, v int) [][]*canvas.Image {
	matrix := make([][]*canvas.Image, v)

	for i := 0; i < v; i++ {
		matrix[i] = make([]*canvas.Image, h)
		for j := 0; j < h; j++ {
			image := canvas.NewImageFromFile("./static/grass.png")
			image.FillMode = canvas.ImageFillOriginal
			image.Resize(fyne.NewSize(32, 32))
			matrix[i][j] = image
		}
	}

	return matrix
}
