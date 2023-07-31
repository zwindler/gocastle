package screens

import (
	"fmt"
	"gocastle/maps"
	"gocastle/model"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

const (
	tileSize = 32
)

var (
	mapColumns int
	mapRows    int

	fyneTileSize = fyne.NewSize(tileSize, tileSize)

	mainContent   *fyne.Container
	mapContainer  *fyne.Container
	statsTextArea fyne.CanvasObject

	// I have to declare these here because I can't use *widget.Scroll as type :-(
	// It's a shame because I redeclare them later
	logsArea               = container.NewVBox()
	logsScrollableTextArea = container.NewVScroll(logsArea)
	scrollableMapContainer = container.NewScroll(container.NewWithoutLayout())

	// TODO clean this
	healthPointsValueLabel *canvas.Text
	manaPointsValueLabel   *canvas.Text
	timeSpentValueLabel    *canvas.Text

	currentWindow fyne.Window
)

// ShowGameScreen is the main function of the game screen
func ShowGameScreen(window fyne.Window) {
	currentWindow = window
	mapContainer = container.NewWithoutLayout()

	// generate a scrollable container which contains the map container
	scrollableMapContainer = container.NewScroll(createMapArea(mapContainer))
	// bottom right corner is the stats box area
	statsTextArea = createStatsArea()
	// merge log area and stats area
	bottom := container.NewBorder(nil, nil, nil, statsTextArea, logsScrollableTextArea)
	// merge map and bottom
	mainContent = container.NewBorder(nil, bottom, nil, nil, scrollableMapContainer)

	window.Canvas().SetOnTypedKey(mapKeyListener)
	window.SetContent(mainContent)

	// TODO create a separate function for this
	// set player on map and draw it
	player.RefreshStats(false)
	player.Avatar.DrawAvatar(mapContainer)
	centerMapOnPlayer()
	drawNPCList(mapContainer)
}

// createMapArea generates a fyne container containing the map tiles
func createMapArea(mapContainer *fyne.Container) fyne.CanvasObject {
	mapRows, mapColumns = currentMap.GetMapSize()
	imageMatrix := createMapMatrix(mapRows, mapColumns)

	for row := 0; row < mapRows; row++ {
		currentLine := float32(row) * tileSize
		for column := 0; column < mapColumns; column++ {
			tile := imageMatrix[row][column]
			tile.Resize(fyneTileSize)
			currentPos := fyne.NewPos(float32(column)*tileSize, currentLine)
			tile.Move(currentPos)
			mapContainer.Add(tile)
		}
	}

	// create a transparent filler to trick scrollable containers
	limits := canvas.NewImageFromFile("static/transparent_tile.png")
	limits.FillMode = canvas.ImageFillStretch
	limits.SetMinSize(fyne.NewSize(float32(mapColumns)*16, float32(mapRows)*16))

	return container.NewGridWithColumns(2,
		mapContainer, layout.NewSpacer(),
		layout.NewSpacer(), limits)
}

// drawNPCList draws the NPC's Avatars images on the mapContainer
func drawNPCList(mapContainer *fyne.Container) {
	// Loop through the NPC data slice and create/draw each NPC
	for _, npc := range currentMap.NPCList.List {
		npc.Avatar.DrawAvatar(mapContainer)
	}
}

// createStatsArea creates the stats area containing health points, mana points,
// time spent, and location info.
func createStatsArea() fyne.CanvasObject {
	healthPointsValueLabel = canvas.NewText("", model.TextColor)
	manaPointsValueLabel = canvas.NewText("", model.TextColor)
	timeSpentValueLabel = canvas.NewText("", model.TextColor)

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
	updateStatsArea()

	for _, textObj := range statsTextObjects {
		textObj.TextSize = 16
	}

	// Add all the canvas.NewText objects to the statsTextArea
	statsTextArea := container.New(layout.NewGridLayout(2))
	for _, textObj := range statsTextObjects {
		statsTextArea.Add(textObj)
	}

	return statsTextArea
}

// updateStatsArea refreshes the values in StatsArea
func updateStatsArea() {
	healthPointsValueLabel.Text = fmt.Sprintf("%d/%d", player.CurrentHP, player.MaxHP)
	healthPointsValueLabel.Refresh()

	manaPointsValueLabel.Text = fmt.Sprintf("%d/%d", player.CurrentMP, player.MaxMP)
	manaPointsValueLabel.Refresh()

	timeSpentValueLabel.Text = model.FormatDuration(model.TimeSinceBegin, "short")
	timeSpentValueLabel.Refresh()
}

// createMapMatrix creates the tiles matrix ([][]*canvas.Image)
func createMapMatrix(numRows, numColumns int) [][]*canvas.Image {
	matrix := make([][]*canvas.Image, numRows)

	// extract the needed tiles from the Tileset
	// create a table of subimages (image.Image type)
	loadedTiles, err := maps.LoadTilesFromTileset(maps.TilesTypes)
	if err != nil {
		err = fmt.Errorf("unable to load tile from Tileset: %w", err)
		log.Fatalf("MapMatrix error: %s", err)
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

// mapKeyListener is the main loop function in this screen
func mapKeyListener(event *fyne.KeyEvent) {
	directions := map[fyne.KeyName]struct{ dx, dy int }{
		fyne.KeyUp:    {0, -1},
		fyne.KeyZ:     {0, -1},
		fyne.KeyE:     {1, -1},
		fyne.KeyRight: {1, 0},
		fyne.KeyD:     {1, 0},
		fyne.KeyC:     {1, 1},
		fyne.KeyDown:  {0, 1},
		fyne.KeyX:     {0, 1},
		fyne.KeyW:     {-1, 1},
		fyne.KeyLeft:  {-1, 0},
		fyne.KeyQ:     {-1, 0},
		fyne.KeyA:     {-1, -1},
	}

	direction, ok := directions[event.Name]
	if !ok {
		if event.Name == fyne.KeyI {
			// Open inventory screen
			ShowInventoryScreen(currentWindow)
		} else if event.Name == fyne.KeyS {
			ShowSaveGameScreen(currentWindow)
		}
		return // Ignore keys that are not part of the directions map
	}
	newX := player.Avatar.PosX + direction.dx
	newY := player.Avatar.PosY + direction.dy

	actOnDirectionKey(newX, newY)

	centerMapOnPlayer()
	updateStatsArea()
	newTurnForNPCs()
}

// centerMapOnPlayer will center scrollable map focus on player as best it can
func centerMapOnPlayer() {
	// the idea is to focus on the player position
	// but we need various informations to compute this

	// Let's start by getting the player real coordinates in pixels
	x := float32(tileSize * player.Avatar.PosX)
	y := float32(tileSize * player.Avatar.PosY)

	// we also need window size (because by default it's 800x600
	// but it can be resized!)
	// I can't use scrollableMapContainer because it's Size() is wrong (always 32x32)
	// window X size is easy to determine, it's the width of the content container
	containerX := mainContent.Size().Width
	// window Y is harder to get. If we take "content" container it will be off
	// because content also includes logs+stats container
	// so I need to remove statsTextArea Height
	containerY := mainContent.Size().Height - statsTextArea.MinSize().Height

	// now, I can focus the scrollable map by adding an offset
	// but the tricky part is that you don't want to move the offset until
	// player is already in the middle of the screen, or else when the player
	// is close to the border, it'll look weird
	// Castle of the Wind had the exact same behavior
	// The easiest way to do this is to remove half of the screen width/height
	// (but make sure before it's always >= 0)

	//fmt.Printf("%f %f", containerX, containerY)
	if x < containerX/2 {
		x = containerX / 2
	}
	if y < containerY/2 {
		y = containerY / 2
	}
	scrollableMapContainer.Offset = fyne.NewPos(x-containerX/2, y-containerY/2)
	scrollableMapContainer.Refresh()
}

// addLogEntry adds entries in the Log scrollable screen
func addLogEntry(logString string) {
	fullLogString := model.FormatDuration(model.TimeSinceBegin, "long") + ": " + logString
	logsEntry := canvas.NewText(fullLogString, model.TextColor)
	logsEntry.TextSize = 12
	logsArea.Add(logsEntry)
	logsScrollableTextArea.ScrollToBottom()
}
