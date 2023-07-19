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
	player  = model.Player
	NPCList = model.NPCsOnCurrentMap{}

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

	// TODO create a separate function for this
	// set player on map and draw it
	player.Avatar.PosX, player.Avatar.PosY = 2, 4
	drawSubject(mapContainer, player.Avatar)

	// set farmer on map and draw it
	farmer := model.Farmer
	farmer.Avatar.PosX, farmer.Avatar.PosY = 10, 15
	NPCList.List = append(NPCList.List, farmer)
	drawSubject(mapContainer, farmer.Avatar)

	// set wolf on map and draw it
	wolf := model.Wolf
	wolf.Avatar.PosX, wolf.Avatar.PosY = 22, 22
	NPCList.List = append(NPCList.List, wolf)
	drawSubject(mapContainer, wolf.Avatar)

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

func drawSubject(mapContainer *fyne.Container, subject model.Avatar) {
	subject.CanvasImage.FillMode = canvas.ImageFillOriginal
	subject.CanvasImage.Resize(fyneTileSize)
	subject.CanvasImage.Move(fyne.NewPos(float32(subject.PosX*tileSize), float32(subject.PosY*tileSize)))
	mapContainer.Add(subject.CanvasImage)
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

	newX := player.Avatar.PosX + direction.dx
	newY := player.Avatar.PosY + direction.dy

	// moving costs 3 seconds
	model.TimeSinceBegin = model.TimeSinceBegin + 3
	updateStats()

	fmt.Printf("Trying to move character \n")
	if checkWalkable(newX, newY) {
		fmt.Printf("Moving character \n")
		moveAvatar(newX, newY, &player.Avatar)
	} else {
		fmt.Println("You are blocked!")
		logsEntry := canvas.NewText(model.FormatDuration(model.TimeSinceBegin, "long")+": you are blocked!", model.TextColor)
		logsEntry.TextSize = 12
		logsArea.Add(logsEntry)
		logsScrollableTextArea.ScrollToBottom()
	}

	newTurnForNPCs()
}

func newTurnForNPCs() {
	// for all NPCs, move on a random adjacent tile
	for _, npc := range NPCList.List {
		fmt.Printf("Moving %s from %d %d\n", npc.Name, npc.Avatar.PosX, npc.Avatar.PosY)
		newX := npc.Avatar.PosX + rand.Intn(3) - 1
		newY := npc.Avatar.PosY + rand.Intn(3) - 1
		if checkWalkable(newX, newY) {
			moveAvatar(newX, newY, &npc.Avatar)
		}
	}

}

func checkWalkable(futurePosX int, futurePosY int) bool {
	if !checkOutOfBounds(futurePosX, futurePosY) &&
		checkTileIsWalkable(futurePosX, futurePosY) &&
		dontCollideWithPlayer(futurePosX, futurePosY) &&
		dontCollideWithNPCs(futurePosX, futurePosY) {
		return true
	}
	return false
}

func checkOutOfBounds(futurePosX int, futurePosY int) bool {
	if futurePosX >= 0 && futurePosX < mapColumns &&
		futurePosY >= 0 && futurePosY < mapRows {
		return false
	}
	return true
}

func checkTileIsWalkable(futurePosX int, futurePosY int) bool {
	return maps.TilesTypes[currentMap.MapMatrix[futurePosY][futurePosX]].IsWalkable
}

func dontCollideWithPlayer(futurePosX int, futurePosY int) bool {
	if player.PosX == futurePosX && player.PosY == futurePosY {
		return false
	}
	return true
}

func dontCollideWithNPCs(futurePosX int, futurePosY int) bool {
	// for all NPCs, check if future position would collide
	for _, npc := range NPCList.List {
		if npc.Avatar.PosX == futurePosX && npc.Avatar.PosY == futurePosY {
			return false
		}
	}
	// coordinates are free for movement
	return true
}

func moveAvatar(futurePosX int, futurePosY int, subject *model.Avatar) {
	fmt.Printf("Moving from %d %d to %d %d\n", subject.PosX, subject.PosY, futurePosX, futurePosY)

	// assign new values for subject position
	subject.PosX = futurePosX
	subject.PosY = futurePosY

	subject.CanvasImage.Move(fyne.NewPos(float32(futurePosX*tileSize), float32(futurePosY*tileSize)))
}
