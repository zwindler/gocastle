package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func ShowMapScreen(window fyne.Window) {
	h, v := 30, 30
	imageMatrix := createMapMatrix(h, v)
	firstLine := container.NewHBox()

	player := canvas.NewImageFromFile("./static/warrior.png")
	player.FillMode = canvas.ImageFillOriginal
	player.Resize(fyne.NewSize(32, 32))

	horizontalBorder := canvas.NewImageFromFile("static/black_hline.png")
	horizontalBorder.FillMode = canvas.ImageFillOriginal
	horizontalBorder.Resize(fyne.NewSize(float32(h-1)*32, 1))
	firstLine.Add(horizontalBorder)
	verticalBorder := canvas.NewImageFromFile("static/black_vline.png")
	verticalBorder.FillMode = canvas.ImageFillOriginal
	verticalBorder.Resize(fyne.NewSize(1, float32(v-1)*32))
	secondLine := container.NewHBox(verticalBorder)

	mapContainer := container.NewWithoutLayout()
	for i := 0; i < v; i++ {
		currentLine := float32(i) * 32
		for j := 0; j < h; j++ {
			tile := imageMatrix[i][j]
			tile.Resize(fyne.NewSize(32, 32))
			currentPos := fyne.NewPos(currentLine, float32(j)*32)
			tile.Move(currentPos)
			mapContainer.Add(tile)
			if i == 2 && j == 4 {
				player.Move(currentPos)
				mapContainer.Add(player)
			}
		}
	}
	secondLine.Add(mapContainer)

	scrollableMapContainer := container.NewVBox(firstLine, secondLine)
	content := container.NewScroll(scrollableMapContainer)

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
