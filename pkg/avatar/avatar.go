package avatar

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"github.com/zwindler/gocastle/pkg/coord"
	"github.com/zwindler/gocastle/pkg/embedimages"
)

type Avatar struct {
	CanvasImage          *canvas.Image
	CanvasPath           string
	Coord                coord.Coord
	ObjectInMapContainer *fyne.CanvasObject
}

// New create a new Avatar.
func New(avatar Avatar) Avatar {
	// Now this function is useless, but it will be useful in the future
	return avatar
}

// Spawn create a copy of an Avatar on given x,y coordinates.
func Spawn(avatar Avatar, coord coord.Coord) Avatar {
	img, _ := embedimages.GetImageFromEmbed(avatar.CanvasPath)

	return Avatar{
		CanvasPath:  avatar.CanvasPath,
		CanvasImage: canvas.NewImageFromImage(img),
		Coord:       coord,
	}
}

// Refresh allows to refresh Avatar Image in case it was removed (save/load).
func (subject *Avatar) Refresh() {
	img, _ := embedimages.GetImageFromEmbed(subject.CanvasPath)

	subject.CanvasImage = canvas.NewImageFromImage(img)
}

// Move moves avatar's coordinates and updates image position on map.
func (subject *Avatar) Move(mapContainer *fyne.Container, futurePosX, futurePosY int) {
	// assign new values for subject position
	subject.Coord.X, subject.Coord.Y = futurePosX, futurePosY

	subject.CanvasImage.Move(fyne.NewPos(float32(futurePosX*coord.TileSize), float32(futurePosY*coord.TileSize)))

	// remove/re-add Avatar from mapContainer to redraw it on top
	mapContainer.Remove(*subject.ObjectInMapContainer)
	mapContainer.Add(*subject.ObjectInMapContainer)
}

// Draw displays an avatar's image on the mapContainer.
func (subject *Avatar) Draw(mapContainer *fyne.Container) {
	subject.CanvasImage.FillMode = canvas.ImageFillOriginal
	subject.CanvasImage.Resize(coord.FyneTileSize)

	mapContainer.Add(subject.CanvasImage)

	// determine Object in fyne.container.Objects slice
	avatarInMapContainer := &mapContainer.Objects[len(mapContainer.Objects)-1]
	subject.ObjectInMapContainer = avatarInMapContainer

	subject.Move(mapContainer, subject.Coord.X, subject.Coord.Y)
}

// DistanceFromAvatar computes the distance between 2 Avatars.
func (subject *Avatar) DistanceFromAvatar(subject2 *Avatar) float64 {
	dx := float64(subject.Coord.X - subject2.Coord.X)
	dy := float64(subject.Coord.Y - subject2.Coord.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// MoveTowardsAvatar is a trivial pathfinding algorithm for NPCs.
func (subject *Avatar) MoveTowardsAvatar(subject2 *Avatar) (int, int) {
	// Calculate the distance between the Avatar and the other Avatar in the x and y directions
	deltaX := subject2.Coord.X - subject.Coord.X
	deltaY := subject2.Coord.Y - subject.Coord.Y

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
	return subject.Coord.X + moveX, subject.Coord.Y + moveY
}

// CollideWithPlayer returns true if we are going to collide with player, false instead.
func (subject *Avatar) CollideWithPlayer(futurePosX, futurePosY int) bool {
	return (subject.Coord.X == futurePosX && subject.Coord.Y == futurePosY)
}
