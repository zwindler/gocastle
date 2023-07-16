package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var (
	playerPosX   int = 2
	playerPosY   int = 4
	mapMaxX      int = 30
	mapMaxY      int = 30
	mapContainer *fyne.Container
	playerAvatar *canvas.Image
)

func ShowMapScreen(window fyne.Window) {
	imageMatrix := createMapMatrix(mapMaxX, mapMaxY)

	firstLine := container.NewHBox()
	horizontalBorder := canvas.NewImageFromFile("static/black_hline.png")
	horizontalBorder.FillMode = canvas.ImageFillOriginal

	verticalLine := canvas.NewImageFromFile("static/black_vline.png")
	verticalLine.FillMode = canvas.ImageFillOriginal
	verticalBorder := container.NewVBox()

	mapContainer = container.NewWithoutLayout()
	for i := 0; i < mapMaxY; i++ {
		firstLine.Add(horizontalBorder)
		currentLine := float32(i) * 32
		for j := 0; j < mapMaxX; j++ {
			if j == 0 {
				verticalBorder.Add(verticalLine)
			}
			tile := imageMatrix[i][j]
			tile.Resize(fyne.NewSize(32, 32))
			currentPos := fyne.NewPos(currentLine, float32(j)*32)
			tile.Move(currentPos)
			mapContainer.Add(tile)
		}
	}
	playerAvatar = canvas.NewImageFromFile("./static/warrior.png")
	playerAvatar.FillMode = canvas.ImageFillOriginal
	playerAvatar.Resize(fyne.NewSize(32, 32))
	playerAvatar.Move(fyne.NewPos(float32(playerPosX*32), float32(playerPosY*32)))
	mapContainer.Add(playerAvatar)

	secondLine := container.NewHBox(verticalBorder, mapContainer)

	scrollableMapContainer := container.NewVBox(firstLine, secondLine)
	scrollableMapContainer.Resize(fyne.NewSize(float32(mapMaxX)*32, float32(mapMaxY)*32))
	content := container.NewScroll(scrollableMapContainer)
	window.Canvas().SetOnTypedKey(mapKeyListener)

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

func mapKeyListener(event *fyne.KeyEvent) {
	if event.Name == fyne.KeyUp {
		if playerPosY > 0 {
			playerPosY = playerPosY - 1
		}
	} else if event.Name == fyne.KeyDown {
		if playerPosY < mapMaxY-1 {
			playerPosY = playerPosY + 1
		}
	} else if event.Name == fyne.KeyLeft {
		if playerPosX > 0 {
			playerPosX = playerPosX - 1
		}
	} else if event.Name == fyne.KeyRight {
		if playerPosX < mapMaxX-1 {
			playerPosX = playerPosX + 1
		}
	}

	movePlayer()

}

func movePlayer() {
	playerAvatar.Move(fyne.NewPos(float32(playerPosX*32), float32(playerPosY*32)))
}
