package main

import (
	"gocastle/model"
	"gocastle/screens"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	goCastle := app.New()
	goCastle.Settings().SetTheme(&model.CustomTheme{})
	if goCastle.Settings().ThemeVariant() == 0 {
		// dark theme
		model.TextColor = color.White
	} else {
		// light theme
		model.TextColor = color.Black
	}
	mainWindow := goCastle.NewWindow("GoCastle")

	screens.ShowMenuScreen(mainWindow)

	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
}
