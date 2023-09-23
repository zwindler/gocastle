package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"github.com/zwindler/gocastle/pkg/newtheme"
	"github.com/zwindler/gocastle/pkg/timespent"
)

const (
	tileSize = 32
)

var (
	fyneTileSize = fyne.NewSize(tileSize, tileSize)

	mainContent   *fyne.Container
	mapContainer  *fyne.Container
	statsTextArea fyne.CanvasObject

	// I have to declare these here because I can't use *widget.Scroll as type :-(
	// It's a shame because I redeclare them later.
	logsArea               = container.NewVBox()
	logsScrollableTextArea = container.NewVScroll(logsArea)
	scrollableMapContainer = container.NewScroll(container.NewWithoutLayout())

	// TODO clean this.
	healthPointsValueLabel *canvas.Text
	manaPointsValueLabel   *canvas.Text
	timeSpentValueLabel    *canvas.Text

	currentWindow fyne.Window
)

// ShowGameScreen is the main function of the game screen.
func ShowGameScreen(window fyne.Window) {
	currentWindow = window
	mapContainer = container.NewWithoutLayout()

	// generate a scrollable container which contains the map container
	mapImage := createMapCanvasImage()
	mapContainer.Add(mapImage)
	scrollableMapContainer = container.NewScroll(mapContainer)

	// TODO create a separate function for this
	// set player on map and draw it
	player.Avatar.Draw(mapContainer)
	drawNPCList(mapContainer)
	drawObjectList(mapContainer)

	// bottom right corner is the stats box area
	statsTextArea = createStatsArea()
	// merge log area and stats area
	bottom := container.NewBorder(nil, nil, nil, statsTextArea, logsScrollableTextArea)
	// merge map and bottom
	mainContent = container.NewBorder(nil, bottom, nil, nil, scrollableMapContainer)

	window.SetContent(mainContent)
	window.Canvas().SetOnTypedKey(mapKeyListener)

	centerMapOnPlayer()
}

// drawNPCList draws the NPC's Avatars images on the mapContainer.
func drawNPCList(mapContainer *fyne.Container) {
	// Loop through the NPC data slice and create/draw each NPC
	for _, npc := range currentMap.NPCList {
		npc.Avatar.Draw(mapContainer)
	}
}

// drawObjectList draws the "Objects on map" images on the mapContainer.
func drawObjectList(mapContainer *fyne.Container) {
	// Loop through the ObjectList slice and create/draw each Object
	for _, object := range currentMap.ObjectList {
		object.DrawObject(mapContainer)
	}
}

// createStatsArea creates the stats area containing health points, mana points,
// time spent, and location info.
func createStatsArea() fyne.CanvasObject {
	healthPointsValueLabel = canvas.NewText("", newtheme.TextColor)
	manaPointsValueLabel = canvas.NewText("", newtheme.TextColor)
	timeSpentValueLabel = canvas.NewText("", newtheme.TextColor)

	// Create an array to store all the canvas.NewText objects
	statsTextObjects := []*canvas.Text{
		canvas.NewText("Health Points:", newtheme.TextColor),
		healthPointsValueLabel,
		canvas.NewText("Mana Points:", newtheme.TextColor),
		manaPointsValueLabel,
		canvas.NewText("Time spent:", newtheme.TextColor),
		timeSpentValueLabel,
		canvas.NewText("Location:", newtheme.TextColor),
		canvas.NewText(currentMap.Name, newtheme.TextColor),
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

// updateStatsArea refreshes the values in StatsArea.
func updateStatsArea() {
	healthPointsValueLabel.Text = player.HP.String()
	healthPointsValueLabel.Refresh()

	manaPointsValueLabel.Text = player.MP.String()
	manaPointsValueLabel.Refresh()

	timeSpentValueLabel.Text = timespent.FormatDuration(timespent.ShortFormat)
	timeSpentValueLabel.Refresh()
}

// createMapImage creates an image based on the tiles stored in currentMap.
func createMapCanvasImage() *canvas.Image {
	if currentMap.MapImage == nil {
		currentMap.GenerateMapImage()
	}
	fullCanvasImage := canvas.NewImageFromImage(currentMap.MapImage)
	fullCanvasImage.FillMode = canvas.ImageFillOriginal
	fullCanvasImage.Resize(fyne.NewSize(currentMap.GetMapImageSize()))

	return fullCanvasImage
}

// mapKeyListener is the main loop function in this screen.
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
		switch event.Name {
		case fyne.KeyI:
			ShowInventoryScreen(currentWindow)
		case fyne.KeyS:
			ShowSaveGameScreen(currentWindow)
		case fyne.KeyL:
			ShowLoadGameScreen(currentWindow)
		}
		return // Ignore keys that are not part of the directions map
	}
	newX := player.Avatar.Coord.X + direction.dx
	newY := player.Avatar.Coord.Y + direction.dy

	actOnDirectionKey(newX, newY)

	centerMapOnPlayer()
	updateStatsArea()
	newTurnForNPCs()
}

// centerMapOnPlayer will center scrollable map focus on player as best it can.
func centerMapOnPlayer() {
	// the idea is to focus on the player position
	// but we need various informations to compute this

	// Let's start by getting the player real coordinates in pixels
	x := float32(tileSize * player.Avatar.Coord.X)
	y := float32(tileSize * player.Avatar.Coord.Y)

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

	if x < containerX/2 {
		x = containerX / 2
	}
	if y < containerY/2 {
		y = containerY / 2
	}
	scrollableMapContainer.Offset = fyne.NewPos(x-containerX/2, y-containerY/2)
	scrollableMapContainer.Refresh()
}

// addLogEntry adds entries in the Log scrollable screen.
func addLogEntry(logString string) {
	fullLogString := timespent.FormatDuration(timespent.LongFormat) + ": " + logString
	logsEntry := canvas.NewText(fullLogString, newtheme.TextColor)
	logsEntry.TextSize = 12
	logsArea.Add(logsEntry)
	logsScrollableTextArea.ScrollToBottom()
}
