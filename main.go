package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/zwindler/gocastle/pkg/newtheme"
	"github.com/zwindler/gocastle/pkg/screens"
)

func main() {
	goCastle := app.New()
	goCastle.Settings().SetTheme(&newtheme.CustomTheme{})
	if goCastle.Settings().ThemeVariant() == 0 {
		// dark theme
		newtheme.TextColor = color.White
	} else {
		// light theme
		newtheme.TextColor = color.Black
	}
	mainWindow := goCastle.NewWindow("GoCastle")

	screens.ShowMenuScreen(mainWindow)

	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
}
