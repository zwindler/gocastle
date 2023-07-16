package screens

import (
	"fmt"
	"gocastle/maps"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var (
	playerPosX   int = 2
	playerPosY   int = 4
	playerAvatar *canvas.Image
	PNJ1PosX     int = 10
	PNJ1PosY     int = 15
	PNJ1         *canvas.Image
	mapColumns   int
	mapRows      int
	mapContainer *fyne.Container
	currentMap   = maps.Map1
)

func ShowMapScreen(window fyne.Window) {
	mapRows = len(currentMap)
	if mapRows > 0 {
		mapColumns = len(currentMap[0])
	}
	imageMatrix := createMapMatrix(mapRows, mapColumns)

	firstLine := container.NewHBox()
	horizontalBorder := canvas.NewImageFromFile("static/black_hline.png")
	horizontalBorder.FillMode = canvas.ImageFillOriginal

	verticalLine := canvas.NewImageFromFile("static/black_vline.png")
	verticalLine.FillMode = canvas.ImageFillOriginal
	verticalBorder := container.NewVBox()

	mapContainer = container.NewWithoutLayout()
	for row := 0; row < mapRows; row++ {
		verticalBorder.Add(verticalLine)
		currentLine := float32(row) * 32
		for column := 0; column < mapColumns; column++ {
			if row == 0 {
				firstLine.Add(horizontalBorder)
			}
			tile := imageMatrix[row][column]
			tile.Resize(fyne.NewSize(32, 32))
			currentPos := fyne.NewPos(float32(column)*32, currentLine)
			tile.Move(currentPos)
			mapContainer.Add(tile)
		}
	}
	playerAvatar = canvas.NewImageFromFile("./static/warrior.png")
	playerAvatar.FillMode = canvas.ImageFillOriginal
	playerAvatar.Resize(fyne.NewSize(32, 32))
	playerAvatar.Move(fyne.NewPos(float32(playerPosX*32), float32(playerPosY*32)))
	mapContainer.Add(playerAvatar)

	PNJ1 = canvas.NewImageFromFile("./static/farmer.png")
	PNJ1.FillMode = canvas.ImageFillOriginal
	PNJ1.Resize(fyne.NewSize(32, 32))
	PNJ1.Move(fyne.NewPos(float32(PNJ1PosX*32), float32(PNJ1PosY*32)))
	mapContainer.Add(PNJ1)

	secondLine := container.NewHBox(verticalBorder, mapContainer)

	scrollableMapContainer := container.NewVBox(firstLine, secondLine)
	scrollableMapContainer.Resize(fyne.NewSize(float32(mapColumns)*32, float32(mapRows)*32))
	content := container.NewScroll(scrollableMapContainer)
	window.Canvas().SetOnTypedKey(mapKeyListener)

	window.SetContent(content)
}

func createMapMatrix(numRows, numColumns int) [][]*canvas.Image {
	matrix := make([][]*canvas.Image, numRows)

	// extract the needed tiles from the Tileset
	// create a table of subimages (image.Image type)
	loadedTiles, err := maps.LoadTilesFromTileset(maps.TilesTypes)
	if err != nil {
		fmt.Errorf("unable to load tile from Tileset: %w", err)
		// TODO error handling
	}

	// create the full matrix first to avoid out of bounds exception
	for row := 0; row < mapRows; row++ {
		matrix[row] = make([]*canvas.Image, numColumns)
	}
	for row := 0; row < mapRows; row++ {
		for column := 0; column < numColumns; column++ {
			image := loadedTiles[currentMap[row][column]]
			imageCanvas := canvas.NewImageFromImage(image)
			imageCanvas.FillMode = canvas.ImageFillOriginal
			imageCanvas.Resize(fyne.NewSize(32, 32))
			matrix[row][column] = imageCanvas
		}
	}

	return matrix
}

func mapKeyListener(event *fyne.KeyEvent) {
	directions := map[fyne.KeyName]struct{ dx, dy int }{
		fyne.KeyUp:    {0, -1},
		fyne.KeyZ:     {0, -1},
		fyne.KeyE:     {1, -1},
		fyne.KeyRight: {1, 0},
		fyne.KeyD:     {1, 0},
		fyne.KeyC:     {1, 1},
		fyne.KeyDown:  {0, 1},
		fyne.KeyS:     {0, 1},
		fyne.KeyX:     {0, 1},
		fyne.KeyW:     {-1, 1},
		fyne.KeyLeft:  {-1, 0},
		fyne.KeyQ:     {-1, 0},
		fyne.KeyA:     {-1, -1},
	}

	direction, ok := directions[event.Name]
	if !ok {
		return // Ignore keys that are not part of the directions map
	}

	newX := playerPosX + direction.dx
	newY := playerPosY + direction.dy

	if checkWalkable(newX, newY) {
		movePlayer(newX, newY)
	} else {
		fmt.Println("You are blocked!")
	}

	newTurnForPNJs()
}

func checkWalkable(futurePosX int, futurePosY int) bool {
	if futurePosX >= 0 && futurePosX < mapColumns &&
		futurePosY >= 0 && futurePosY < mapRows &&
		maps.TilesTypes[currentMap[futurePosY][futurePosX]].IsWalkable &&
		(playerPosX != futurePosX || playerPosY != futurePosY) &&
		(PNJ1PosX != futurePosX || PNJ1PosY != futurePosY) {
		//TODO make a function to check for other characters presence
		return true
	}
	return false
}

func newTurnForPNJs() {
	// Generate random numbers between -1 and 1
	randDeltaX := rand.Intn(3) - 1
	randDeltaY := rand.Intn(3) - 1

	newPNJ1PosX := PNJ1PosX + randDeltaX
	newPNJ1PosY := PNJ1PosY + randDeltaY

	if checkWalkable(newPNJ1PosX, newPNJ1PosY) {
		movePNJ1(newPNJ1PosX, newPNJ1PosY)
	}
}

//TODO unify movePlayer and movePNJ
func movePlayer(futurePosX int, futurePosY int) {
	// assign new values for player position
	playerPosX = futurePosX
	playerPosY = futurePosY
	playerAvatar.Move(fyne.NewPos(float32(playerPosX*32), float32(playerPosY*32)))
}

func movePNJ1(futurePosX int, futurePosY int) {
	// Assign new values for PNJ1 position
	PNJ1PosX = futurePosX
	PNJ1PosY = futurePosY
	PNJ1.Move(fyne.NewPos(float32(PNJ1PosX*32), float32(PNJ1PosY*32)))
}
