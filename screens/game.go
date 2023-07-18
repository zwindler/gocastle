package screens

import (
	"fmt"
	"gocastle/maps"
	"gocastle/model"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

const (
	tileSize = 32
)

var (
	playerPosX   int = 2
	playerPosY   int = 4
	playerAvatar *canvas.Image
	PNJ1PosX     int = 10
	PNJ1PosY     int = 15
	PNJ1         *canvas.Image

	mapColumns int
	mapRows    int

	currentMap   = maps.Town
	fyneTileSize = fyne.NewSize(tileSize, tileSize)

	//mapContainer           = container.NewWithoutLayout()

	logsArea               = container.NewVBox()
	logsScrollableTextArea = container.NewVScroll(logsArea)

	healthPointsValueLabel = canvas.NewText("10/10", model.TextColor)
	manaPointsValueLabel   = canvas.NewText("10/10", model.TextColor)
	timeSpentValueLabel    = canvas.NewText("0d0:0:0", model.TextColor)
)

func ShowGameScreen(window fyne.Window) {
	mapContainer := container.NewWithoutLayout()
	// generate a scrollable container which contains the map container
	scrollableMapContainer := container.NewScroll(createMapArea(mapContainer))
	scrollableMapContainer.Resize(fyne.NewSize(800, 500))

	// draw player and PNJ
	// TODO create a Subject struct that contains both image and coordinates
	playerAvatar = canvas.NewImageFromFile("./static/warrior.png")
	PNJ1 = canvas.NewImageFromFile("./static/farmer.png")
	drawSubject(mapContainer, playerAvatar, playerPosX, playerPosY)
	drawSubject(mapContainer, PNJ1, PNJ1PosX, PNJ1PosY)

	// already declared in var so has to manipulate it elsewhere
	// TODO improve this?
	logsScrollableTextArea.Resize(fyne.NewSize(600, 100))
	logsScrollableTextArea.Move(fyne.NewPos(0, 501))

	// bottom right corner is the stats box area
	statsTextArea := createStatsArea()
	statsTextArea.Resize(fyne.NewSize(200, 100))
	statsTextArea.Move(fyne.NewPos(601, 501))

	// merge log area and stats area
	bottom := container.NewBorder(nil, nil, nil, statsTextArea, logsScrollableTextArea)

	// merge map and bottom
	content := container.NewBorder(nil, bottom, nil, nil, scrollableMapContainer)

	window.Canvas().SetOnTypedKey(mapKeyListener)
	window.SetContent(content)
}

func createMapArea(mapContainer *fyne.Container) fyne.CanvasObject {
	mapRows = len(currentMap.MapMatrix)
	if mapRows > 0 {
		mapColumns = len(currentMap.MapMatrix[0])
	}
	imageMatrix := createMapMatrix(mapRows, mapColumns)

	horizontalLine := canvas.NewImageFromFile("static/black_hline.png")
	horizontalLine.FillMode = canvas.ImageFillOriginal
	verticalLine := canvas.NewImageFromFile("static/black_vline.png")
	verticalLine.FillMode = canvas.ImageFillOriginal

	// horizontalBorder is composed of images of 1x32px (horizontalLine)
	// to force the minSize of the container
	// TODO improve this
	horizontalBorder := container.NewHBox()
	// vertical border is the same thing but vertical
	verticalBorder := container.NewVBox()

	for row := 0; row < mapRows; row++ {
		verticalBorder.Add(verticalLine)
		currentLine := float32(row) * tileSize
		for column := 0; column < mapColumns; column++ {
			if row == 0 {
				horizontalBorder.Add(horizontalLine)
			}
			tile := imageMatrix[row][column]
			tile.Resize(fyneTileSize)
			currentPos := fyne.NewPos(float32(column)*tileSize, currentLine)
			tile.Move(currentPos)
			mapContainer.Add(tile)
		}
	}
	mapHBox := container.NewHBox(verticalBorder, mapContainer)

	return container.NewVBox(horizontalBorder, mapHBox)
}

func drawSubject(mapContainer *fyne.Container, subject *canvas.Image, posX int, posY int) {
	subject.FillMode = canvas.ImageFillOriginal
	subject.Resize(fyneTileSize)
	subject.Move(fyne.NewPos(float32(posX*tileSize), float32(posY*tileSize)))
	mapContainer.Add(subject)
}

// Create the stats area containing health points, mana points, time spent, and location info.
func createStatsArea() fyne.CanvasObject {
	// Create an array to store all the canvas.NewText objects
	statsTextObjects := []*canvas.Text{
		canvas.NewText("Health Points:", model.TextColor),
		healthPointsValueLabel,
		canvas.NewText("Mana Points:", model.TextColor),
		manaPointsValueLabel,
		canvas.NewText("Time spent:", model.TextColor),
		timeSpentValueLabel,
		canvas.NewText("Location:", model.TextColor),
		canvas.NewText(currentMap.Name, model.TextColor),
	}

	// update HP, MP, time
	updateStats()

	for _, textObj := range statsTextObjects {
		textObj.TextSize = 14
	}

	// Add all the canvas.NewText objects to the statsTextArea
	statsTextArea := container.New(layout.NewGridLayout(2))
	for _, textObj := range statsTextObjects {
		statsTextArea.Add(textObj)
	}

	return statsTextArea
}

func updateStats() {
	healthPointsValueLabel.Text = fmt.Sprintf("%d/%d", model.Player.CurrentHP, model.Player.MaxHP)
	healthPointsValueLabel.Refresh()

	manaPointsValueLabel.Text = fmt.Sprintf("%d/%d", model.Player.CurrentMP, model.Player.MaxMP)
	manaPointsValueLabel.Refresh()

	timeSpentValueLabel.Text = model.FormatDuration(model.TimeSinceBegin, "short")
	timeSpentValueLabel.Refresh()
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
	for row := 0; row < numRows; row++ {
		matrix[row] = make([]*canvas.Image, numColumns)
	}
	for row := 0; row < numRows; row++ {
		for column := 0; column < numColumns; column++ {
			image := loadedTiles[currentMap.MapMatrix[row][column]]
			imageCanvas := canvas.NewImageFromImage(image)
			imageCanvas.FillMode = canvas.ImageFillOriginal
			imageCanvas.Resize(fyneTileSize)
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

	// moving costs 3 seconds
	model.TimeSinceBegin = model.TimeSinceBegin + 3
	updateStats()

	if checkWalkable(newX, newY) {
		movePlayer(newX, newY)
	} else {
		fmt.Println("You are blocked!")
		logsEntry := canvas.NewText(model.FormatDuration(model.TimeSinceBegin, "long")+": you are blocked!", model.TextColor)
		logsEntry.TextSize = 12
		logsArea.Add(logsEntry)
		logsScrollableTextArea.ScrollToBottom()
	}

	newTurnForPNJs()
}

func checkWalkable(futurePosX int, futurePosY int) bool {
	if futurePosX >= 0 && futurePosX < mapColumns &&
		futurePosY >= 0 && futurePosY < mapRows &&
		maps.TilesTypes[currentMap.MapMatrix[futurePosY][futurePosX]].IsWalkable &&
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

	playerAvatar.Move(fyne.NewPos(float32(playerPosX*tileSize), float32(playerPosY*tileSize)))
}

func movePNJ1(futurePosX int, futurePosY int) {
	// Assign new values for PNJ1 position
	PNJ1PosX = futurePosX
	PNJ1PosY = futurePosY
	PNJ1.Move(fyne.NewPos(float32(PNJ1PosX*tileSize), float32(PNJ1PosY*tileSize)))
}
