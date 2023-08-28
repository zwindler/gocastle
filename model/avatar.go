package model

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"github.com/zwindler/gocastle/utils"
)

const (
	tileSize = 32
)

var fyneTileSize = fyne.NewSize(tileSize, tileSize)

type Avatar struct {
	CanvasImage          *canvas.Image
	CanvasPath           string
	PosX                 int
	PosY                 int
	ObjectInMapContainer *fyne.CanvasObject
}

// CreateAvatar create a copy of an Avatar on given x,y coordinates.
func CreateAvatar(avatar Avatar, x, y int) Avatar {
	return Avatar{
		CanvasPath:  avatar.CanvasPath,
		CanvasImage: canvas.NewImageFromImage(utils.GetImageFromEmbed(avatar.CanvasPath)),
		PosX:        x,
		PosY:        y,
	}
}

// MoveAvatar moves avatar's coordinates and updates image position on map.
func (subject *Avatar) MoveAvatar(mapContainer *fyne.Container, futurePosX, futurePosY int) {
	// assign new values for subject position
	subject.PosX = futurePosX
	subject.PosY = futurePosY

	subject.CanvasImage.Move(fyne.NewPos(float32(futurePosX*tileSize), float32(futurePosY*tileSize)))

	// remove/re-add Avatar from mapContainer to redraw it on top
	mapContainer.Remove(*subject.ObjectInMapContainer)
	mapContainer.Add(*subject.ObjectInMapContainer)
}

// DrawAvatar displays an avatar's image on the mapContainer.
func (subject *Avatar) DrawAvatar(mapContainer *fyne.Container) {
	subject.CanvasImage.FillMode = canvas.ImageFillOriginal
	subject.CanvasImage.Resize(fyneTileSize)

	mapContainer.Add(subject.CanvasImage)

	// determine Object in fyne.container.Objects slice
	avatarInMapContainer := &mapContainer.Objects[len(mapContainer.Objects)-1]
	subject.ObjectInMapContainer = avatarInMapContainer

	subject.MoveAvatar(mapContainer, subject.PosX, subject.PosY)
}

// DistanceFromAvatar computes the distance between 2 Avatars.
func (subject *Avatar) DistanceFromAvatar(subject2 *Avatar) float64 {
	dx := float64(subject.PosX - subject2.PosX)
	dy := float64(subject.PosY - subject2.PosY)
	return math.Sqrt(dx*dx + dy*dy)
}

// MoveAvatarTowardsAvatar is a trivial pathfinding algorithm for NPCs.
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
