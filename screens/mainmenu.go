package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/zwindler/gocastle/model"
	"github.com/zwindler/gocastle/pkg/embed"
)

// ShowMenuScreen is the main function of the main screen.
func ShowMenuScreen(window fyne.Window) {
	// Load the background image

	backgroundImage := canvas.NewImageFromImage(embed.GetImageFromEmbed("static/castle_back.png"))

	backgroundImage.FillMode = canvas.ImageFillStretch
	backgroundImage.SetMinSize(fyne.NewSize(800, 600))

	// Initialize a few things here
	model.InitializeCategories()

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

	window.SetContent(container.New(layout.NewMaxLayout(),
		backgroundImage,
		menu))
}
