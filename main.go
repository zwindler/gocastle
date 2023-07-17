package main

import (
	"gocastle/model"
	"gocastle/screens"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	goCastle := app.New()
	goCastle.Settings().SetTheme(&model.CustomTheme{})
	mainWindow := goCastle.NewWindow("GoCastle")

	screens.ShowMenuScreen(mainWindow)

	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.SetFixedSize(true)
	mainWindow.ShowAndRun()
}
