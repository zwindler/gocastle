package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/zwindler/gocastle/model"
	"github.com/zwindler/gocastle/screens"
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
