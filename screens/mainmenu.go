package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowMenuScreen(window fyne.Window) {
	// Load the background image
	backgroundImage := canvas.NewImageFromFile("castle_back.png")
	backgroundImage.FillMode = canvas.ImageFillContain
	backgroundImage.SetMinSize(fyne.NewSize(800, 600))

	// Create buttons
	newGameButton := widget.NewButton("New Game", func() {
	})
	loadGameButton := widget.NewButton("Load Game", func() {
	})
	quitButton := widget.NewButton("Quit", func() {
		window.Close()
	})

	// Resize buttons
	defaultButtonSize := fyne.NewSize(120, 40)
	newGameButton.Resize(defaultButtonSize)
	loadGameButton.Resize(defaultButtonSize)
	quitButton.Resize(defaultButtonSize)

	// Move buttons
	newGameButton.Move(fyne.NewPos(340, 220))
	loadGameButton.Move(fyne.NewPos(340, 275))
	quitButton.Move(fyne.NewPos(340, 330))

	// Create container with all elements
	menu := container.NewWithoutLayout(
		newGameButton,
		loadGameButton,
		quitButton,
	)

	window.SetContent(container.New(layout.NewMaxLayout(),
		backgroundImage,
		menu))

}
