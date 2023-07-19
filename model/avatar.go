package model

import "fyne.io/fyne/v2/canvas"

type Avatar struct {
	CanvasImage *canvas.Image
	CanvasPath  string
	PosX        int
	PosY        int
}
