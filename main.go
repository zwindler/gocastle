package main

import (
	"gocastle/screens"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	mainWindow := myApp.NewWindow("GoCastle")

	screens.ShowMenuScreen(mainWindow)

	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.SetFixedSize(true)
	mainWindow.ShowAndRun()
}
