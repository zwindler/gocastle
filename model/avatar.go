package model

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const (
	tileSize = 32
)

type Avatar struct {
	CanvasImage *canvas.Image
	CanvasPath  string
	PosX        int
	PosY        int
}

func (subject *Avatar) MoveAvatar(futurePosX int, futurePosY int) {
	// assign new values for subject position
	subject.PosX = futurePosX
	subject.PosY = futurePosY

	subject.CanvasImage.Move(fyne.NewPos(float32(futurePosX*tileSize), float32(futurePosY*tileSize)))
}
