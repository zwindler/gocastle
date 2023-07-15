package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func ShowMapScreen(window fyne.Window) {
	h, v := 30, 30
	firstLine := container.NewGridWithColumns(h)
	vBorder := container.NewGridWithRows(v - 1)

	imageMatrix := createMapMatrix(h, v)
	player := canvas.NewImageFromFile("./static/warrior.png")
	player.FillMode = canvas.ImageFillOriginal
	player.Resize(fyne.NewSize(32, 32))

	usableMapContainer := container.NewWithoutLayout()

	for i := 0; i < v; i++ {
		for j := 0; j < h; j++ {
			tile := imageMatrix[i][j]
			tile.Resize(fyne.NewSize(32, 32))
			if i == 0 {
				firstLine.Add(tile)
			} else {
				if j == 0 {
					vBorder.Add(tile)
				} else {
					currentPos := fyne.NewPos(float32(i)*32, float32(j)*32)
					tile.Move(currentPos)
					usableMapContainer.Add(tile)
					if i == 1 && j == 1 {
						player.Move(currentPos)
						usableMapContainer.Add(player)
					}
				}
			}
		}
	}

	secondLine := container.NewHBox(vBorder, usableMapContainer)
	mapContainer := container.NewVBox(firstLine, secondLine)

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
