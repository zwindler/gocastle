package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/zwindler/gocastle/pkg/embedimages"
	"github.com/zwindler/gocastle/pkg/object"
)

// ShowMenuScreen is the main function of the main screen.
func ShowMenuScreen(window fyne.Window) {
	// Load the background image
	img, _ := embedimages.GetImageFromEmbed("static/castle_back.png")
	backgroundImage := canvas.NewImageFromImage(img)

	backgroundImage.FillMode = canvas.ImageFillStretch
	backgroundImage.SetMinSize(fyne.NewSize(800, 600))

	// Initialize a few things here
	object.InitializeCategories()

	// Create buttons
	newGameButton := widget.NewButton("New Game", func() {
		ShowNewGameScreen(window)
	})
	loadGameButton := widget.NewButton("Load Game", func() {
		ShowLoadGameScreen(window)
	})
	quitButton := widget.NewButton("Quit", func() {
		window.Close()
	})

	// Create container with all elements
	buttons := container.New(layout.NewVBoxLayout(),
		newGameButton,
		loadGameButton,
		quitButton,
	)
	menu := container.New(layout.NewCenterLayout(), buttons)

	window.SetContent(container.New(layout.NewStackLayout(),
		backgroundImage,
		menu))
}
