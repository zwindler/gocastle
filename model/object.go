// models/objects.go

package model

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"github.com/zwindler/gocastle/pkg/embed"
)

// ObjectStat represents a specific stat of an object.
type ObjectStat struct {
	Name     string // Stat name, e.g., "Strength", "Health", etc.
	Modifier int    // Modifier value for the stat.
}

// Object represents an object with its properties.
type Object struct {
	Name        string        // Object name.
	Category    string        // Object category.
	Weight      int           // Object weight in grams
	InInventory bool          // Is Object in inventory
	Equipped    bool          // Is Object equipped
	Coord       Coord         // Object position
	Stats       []ObjectStat  // Object stats (e.g., strength, health, etc.).
	CanvasImage *canvas.Image // Object image
	CanvasPath  string        // Image path for Object
}

// Category represents a common object category.
type Category struct {
	Name        string // Category name.
	Description string // Description of the category.
}

var (
	CategoryList []Category

	HuntingKnife = Object{
		Name:     "Hunting knife",
		Category: "Weapon",
		Weight:   200,
		Stats: []ObjectStat{
			{
				Name:     "physicalDamage",
				Modifier: 2,
			},
		},
		CanvasPath: "static/knife.png",
	}

	BluntSword = Object{
		Name:     "Blunt sword",
		Category: "Weapon",
		Weight:   1500,
		Stats: []ObjectStat{
			{
				Name:     "physicalDamage",
				Modifier: 4,
			},
		},
		CanvasPath: "static/sword.png",
	}

	BareHands = Object{
		Name:     "Bare Hands",
		Category: "Weapon",
		Weight:   0,
	}
)

// GenerateCategories creates all the categories based on the provided names and descriptions.
func GenerateCategories(names, descriptions []string) {
	if len(names) != len(descriptions) {
		panic("number of names and descriptions should be the same")
	}

	CategoryList = make([]Category, len(names))
	for i := range names {
		CategoryList[i] = Category{
			Name:        names[i],
			Description: descriptions[i],
		}
	}
}

// InitializeCategories initializes the categories with their names and descriptions.
func InitializeCategories() {
	GenerateCategories(
		[]string{
			"Overgarment",
			"Body Armor",
			"Weapon",
			"Right Ring",
			"Belt",
			"Belt Item",
			"Boots",
			"Head Gear",
			"Neckwear",
			"Shield",
			"Greaves",
			"Gauntlets",
			"Left Ring",
			"Bracers",
		},
		[]string{
			"Outer garments like cloaks or capes.",
			"Gear worn to the chest.",
			"Weapons used for combat.",
			"A ring worn on the right hand.",
			"A belt worn around the waist.",
			"Consumables that are easily accessible in combat.",
			"Footwear.",
			"Head gear (can be hats, helmets,...).",
			"Items worn around the neck.",
			"Shields used for defense.",
			"Protective armor for legs",
			"Protective gloves for hands.",
			"A ring worn on the left hand.",
			"Arm protectors.",
		},
	)
}

// CategoryExists checks if the given category exists in the CommonCategories slice.
func CategoryExists(categoryName string) bool {
	for _, cat := range CategoryList {
		if cat.Name == categoryName {
			return true
		}
	}
	return false
}

// CreateObject creates a copy of the given object and returns it.
// It also validates the category before creating the object.
func CreateObject(obj Object, coord Coord) (Object, error) {
	// Validate the category.
	if !CategoryExists(obj.Category) {
		return Object{}, fmt.Errorf("category '%s' does not exist", obj.Category)
	}

	// Create a new Object with the same properties.
	newObject := Object{
		Name:        obj.Name,
		Category:    obj.Category,
		Weight:      obj.Weight,
		CanvasPath:  obj.CanvasPath,
		CanvasImage: canvas.NewImageFromImage(embed.GetImageFromEmbed(obj.CanvasPath)),
		Coord:       coord,
		Equipped:    false,
		InInventory: false,
	}

	// Copy the ObjectStat slice.
	for _, stat := range obj.Stats {
		newObject.Stats = append(newObject.Stats, ObjectStat{
			Name:     stat.Name,
			Modifier: stat.Modifier,
		})
	}

	return newObject, nil
}

// DrawObject displays an object's image on the mapContainer.
func (subject *Object) DrawObject(mapContainer *fyne.Container) {
	// don't put object in container if object is in inventory
	if !subject.InInventory {
		subject.CanvasImage.FillMode = canvas.ImageFillOriginal
		subject.CanvasImage.Resize(fyneTileSize)

		subject.MoveObject(subject.Coord.X, subject.Coord.Y)

		mapContainer.Add(subject.CanvasImage)
	}
}

// MoveObject moves object's coordinates and updates image position on map.
func (subject *Object) MoveObject(futurePosX, futurePosY int) {
	// assign new values for subject position
	subject.Coord.X = futurePosX
	subject.Coord.Y = futurePosY

	subject.CanvasImage.Move(fyne.NewPos(float32(futurePosX*tileSize), float32(futurePosY*tileSize)))
}

// RefreshObject allows to refresh Object Image in case it was removed (save/load).
func (subject *Object) RefreshObject() {
	subject.CanvasImage = canvas.NewImageFromImage(embed.GetImageFromEmbed(subject.CanvasPath))
}
