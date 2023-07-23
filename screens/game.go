package screens

import (
	"fmt"
	"gocastle/maps"
	"gocastle/model"
	"log"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
)

const (
	tileSize = 32
)

var (
	mapColumns int
	mapRows    int

	fyneTileSize = fyne.NewSize(tileSize, tileSize)

	//mapContainer           = container.NewWithoutLayout()

	logsArea               = container.NewVBox()
	logsScrollableTextArea = container.NewVScroll(logsArea)

	healthPointsValueLabel = canvas.NewText("10/10", model.TextColor)
	manaPointsValueLabel   = canvas.NewText("10/10", model.TextColor)
	timeSpentValueLabel    = canvas.NewText("0d0:0:0", model.TextColor)

	currentWindow fyne.Window
)

func ShowGameScreen(window fyne.Window) {
	currentWindow = window
	mapContainer := container.NewWithoutLayout()
	// generate a scrollable container which contains the map container
	scrollableMapContainer := container.NewScroll(createMapArea(mapContainer))
	scrollableMapContainer.Resize(fyne.NewSize(800, 500))

	// TODO create a separate function for this
	// set player on map and draw it
	player.Avatar.DrawAvatar(mapContainer)
	drawNPCList(mapContainer)

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
	mapRows, mapColumns = currentMap.GetMapSize()
	imageMatrix := createMapMatrix(mapRows, mapColumns)

	horizontalLine := canvas.NewImageFromFile("static/transparent_hline.png")
	horizontalLine.FillMode = canvas.ImageFillOriginal
	verticalLine := canvas.NewImageFromFile("static/transparent_tile.png")
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
	mapHBox := container.NewHBox(mapContainer, verticalBorder)

	return container.NewVBox(mapHBox, horizontalBorder)
}

func drawNPCList(mapContainer *fyne.Container) {
	// Loop through the NPC data slice and create/draw each NPC
	for _, npc := range NPCList.List {
		npc.Avatar.DrawAvatar(mapContainer)
	}
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
	updateStatsBox()

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

func updateStatsBox() {
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

	// before doing anything, check if we aren't out of bounds
	if currentMap.CheckOutOfBounds(newX, newY) {
		// Player tries to escape map, prevent this, lose 2 seconds
		addLogEntry("you are blocked!")
		model.IncrementTimeSinceBegin(2)
	} else {
		// let's check if we find a NPC on our path
		if npcId := NPCList.GetNPCAtPosition(newX, newY); npcId != -1 {
			// get the real NPC from the list, not a copy
			// TODO improve this
			npc := &NPCList.List[npcId]
			// if yes, is the NPC hostile?
			if npc.Hostile {
				// let's attack!
				// TODO make this depending on gear
				addLogEntry(npc.HandleNPCDamage(player.PhysicalDamage))
				npc.CurrentHP = npc.CurrentHP - player.PhysicalDamage
				if npc.IsNPCDead() {
					if player.ChangeXP(npc.LootXP) {
						levelUpEntry := fmt.Sprintf("Level up! You are now level %d", player.Level)
						addLogEntry(levelUpEntry)
						levelUpPopup := showLevelUpScreen()
						dialog.ShowCustomConfirm("Level up!", "Validate", "Close", levelUpPopup, func(validate bool) {
							player.RefreshStats(true)
							updateStatsBox()
						}, currentWindow)
					}
					player.ChangeGold(npc.LootGold)
					NPCList.RemoveNPCByIndex(npcId)
				}
				// attacking costs 5 seconds
				model.IncrementTimeSinceBegin(5)

			} else {
				// NPC is not hostile, we don't want to hurt them, but lost 2s
				addLogEntry("you are blocked!")
				model.IncrementTimeSinceBegin(2)
			}
		} else {
			// no NPC found on our path, let's check if we can move
			if currentMap.CheckTileIsWalkable(newX, newY) {
				// path is free, let's move (3sec cost)
				player.Avatar.MoveAvatar(newX, newY)
				model.IncrementTimeSinceBegin(3)
			}
		}
	}

	updateStatsBox()

	newTurnForNPCs()
}

func addLogEntry(logString string) {
	fullLogString := model.FormatDuration(model.TimeSinceBegin, "long") + ": " + logString
	logsEntry := canvas.NewText(fullLogString, model.TextColor)
	logsEntry.TextSize = 12
	logsArea.Add(logsEntry)
	logsScrollableTextArea.ScrollToBottom()
}

func newTurnForNPCs() {
	// for all NPCs, move on a random adjacent tile
	for index := range NPCList.List {
		npc := &NPCList.List[index]

		var newX, newY int
		if npc.Hostile && npc.Avatar.DistanceFromAvatar(&player.Avatar) <= 10 {
			// player is near, move toward him/her
			newX, newY = npc.Avatar.MoveAvatarTowardsAvatar(&player.Avatar)
		} else {
			// move randomly
			newX = npc.Avatar.PosX + rand.Intn(3) - 1
			newY = npc.Avatar.PosY + rand.Intn(3) - 1

		}

		// don't check / try to move if coordinates stay the same
		if newX != npc.Avatar.PosX || newY != npc.Avatar.PosY {
			// before doing anything, check if we aren't out of bounds
			if !currentMap.CheckOutOfBounds(newX, newY) {
				// let's check if we find another NPC on our NPC's path
				if npcId := NPCList.GetNPCAtPosition(newX, newY); npcId != -1 {
					otherNPC := &NPCList.List[npcId]
					if (npc.Hostile && !otherNPC.Hostile) ||
						(!npc.Hostile && otherNPC.Hostile) {
						// TODO hostile NPC should attack friendly NPC
						// and vice versa
						addLogEntry(fmt.Sprintf("%s tries to attack %s", npc.Name, otherNPC.Name))
					}
					// let's then check we don't collide with player
				} else if player.Avatar.CollideWithPlayer(newX, newY) {
					if npc.Hostile {
						// TODO hostile NPC should attack player
						addLogEntry(fmt.Sprintf("%s tries to attack you", npc.Name))
					}
					// no ones in our NPC's way
				} else if currentMap.CheckTileIsWalkable(newX, newY) {
					npc.Avatar.MoveAvatar(newX, newY)
				}
			}
		}
	}
}
