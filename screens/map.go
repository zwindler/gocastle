package screens

import (
	"gocastle/maps"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var (
	playerPosX   int = 2
	playerPosY   int = 4
	mapMaxX      int
	mapMaxY      int
	mapContainer *fyne.Container
	playerAvatar *canvas.Image
	map1         = maps.Map1
)

func ShowMapScreen(window fyne.Window) {
	mapMaxX = len(map1)
	if mapMaxX > 0 {
		mapMaxY = len(map1[0])
	}
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
			image := canvas.NewImageFromFile(maps.TilesTypes[map1[i][j]])
			image.FillMode = canvas.ImageFillOriginal
			image.Resize(fyne.NewSize(32, 32))
			matrix[i][j] = image
		}
	}

	return matrix
}

func mapKeyListener(event *fyne.KeyEvent) {
	if event.Name == fyne.KeyUp || event.Name == fyne.KeyZ {
		if playerPosY > 0 {
			playerPosY = playerPosY - 1
		}
	} else if event.Name == fyne.KeyE {
		if playerPosY > 0 && playerPosX < mapMaxX-1 {
			playerPosX = playerPosX + 1
			playerPosY = playerPosY - 1
		}
	} else if event.Name == fyne.KeyRight || event.Name == fyne.KeyD {
		if playerPosX < mapMaxX-1 {
			playerPosX = playerPosX + 1
		}
	} else if event.Name == fyne.KeyC {
		if playerPosX < mapMaxX-1 && playerPosY < mapMaxY-1 {
			playerPosX = playerPosX + 1
			playerPosY = playerPosY + 1
		}
	} else if event.Name == fyne.KeyDown || event.Name == fyne.KeyS || event.Name == fyne.KeyX {
		if playerPosY < mapMaxY-1 {
			playerPosY = playerPosY + 1
		}
	} else if event.Name == fyne.KeyW {
		if playerPosY < mapMaxY-1 && playerPosX > 0 {
			playerPosX = playerPosX - 1
			playerPosY = playerPosY + 1
		}
	} else if event.Name == fyne.KeyLeft || event.Name == fyne.KeyQ {
		if playerPosX > 0 {
			playerPosX = playerPosX - 1
		}
	} else if event.Name == fyne.KeyA {
		if playerPosX > 0 && playerPosY > 0 {
			playerPosX = playerPosX - 1
			playerPosY = playerPosY - 1
		}
	}

	movePlayer()

}

func movePlayer() {
	playerAvatar.Move(fyne.NewPos(float32(playerPosX*32), float32(playerPosY*32)))
}
