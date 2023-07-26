package model

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

const (
	tileSize = 32
)

var (
	fyneTileSize = fyne.NewSize(tileSize, tileSize)
)

type Avatar struct {
	CanvasImage *canvas.Image
	CanvasPath  string
	PosX        int
	PosY        int
}

// CreateAvatar create a copy of an Avatar on given x,y coordinates
func CreateAvatar(avatar Avatar, x, y int) Avatar {
	return Avatar{
		CanvasImage: canvas.NewImageFromFile(avatar.CanvasPath),
		PosX:        x,
		PosY:        y,
	}
}

// MoveAvatar moves avatar's coordinates and updates image position on map
func (subject *Avatar) MoveAvatar(futurePosX int, futurePosY int) {
	// assign new values for subject position
	subject.PosX = futurePosX
	subject.PosY = futurePosY

	subject.CanvasImage.Move(fyne.NewPos(float32(futurePosX*tileSize), float32(futurePosY*tileSize)))
}

// DrawAvatar displays an avatar's image on the mapContainer
func (subject *Avatar) DrawAvatar(mapContainer *fyne.Container) {
	subject.CanvasImage.FillMode = canvas.ImageFillOriginal
	subject.CanvasImage.Resize(fyneTileSize)

	subject.MoveAvatar(subject.PosX, subject.PosY)

	mapContainer.Add(subject.CanvasImage)
}

// DistanceFromAvatar computes the distance between 2 Avatars
func (subject *Avatar) DistanceFromAvatar(subject2 *Avatar) float64 {
	dx := float64(subject.PosX - subject2.PosX)
	dy := float64(subject.PosY - subject2.PosY)
	return math.Sqrt(dx*dx + dy*dy)
}

// MoveAvatarTowardsAvatar is a trivial pathfinding algorithm for NPCs
func (subject *Avatar) MoveAvatarTowardsAvatar(subject2 *Avatar) (int, int) {
	// Calculate the distance between the Avatar and the other Avatar in the x and y directions
	deltaX := subject2.PosX - subject.PosX
	deltaY := subject2.PosY - subject.PosY

	moveX := 0
	moveY := 0

	if deltaX > 0 {
		moveX = 1
	} else if deltaX < 0 {
		moveX = -1
	}

	if deltaY > 0 {
		moveY = 1
	} else if deltaY < 0 {
		moveY = -1
	}

	// Update new Avatar's position to move one step closer to the other Avatar
	return subject.PosX + moveX, subject.PosY + moveY
}
