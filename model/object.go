// models/objects.go

package model

import "fmt"

// ObjectStat represents a specific stat of an object.
type ObjectStat struct {
	Name     string // Stat name, e.g., "Strength", "Health", etc.
	Modifier int    // Modifier value for the stat.
}

// Object represents an object with its properties.
type Object struct {
	Name     string       // Object name.
	Category string       // Object category.
	Weight   int          // Object weight in grams
	Equipped bool         // Is object equipped
	Stats    []ObjectStat // Object stats (e.g., strength, health, etc.).
}

// Category represents a common object category.
type Category struct {
	Name        string // Category name.
	Description string // Description of the category.
}

var (
	WeaponCategory = Category{
		Name:        "Weapon",
		Description: "Weapons used for combat.",
	}

	BodyArmorCategory = Category{
		Name:        "BodyArmor",
		Description: "Gear worn to the chest.",
	}

	HeadGearCategory = Category{
		Name:        "HeadGear",
		Description: "Head gear (can be hats, helmets,...).",
	}

	PotionCategory = Category{
		Name:        "Potion",
		Description: "Magical potions with various effects.",
	}

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
	}
)

// Common categories slice.
var CategoryList = []Category{
	WeaponCategory,
	BodyArmorCategory,
	HeadGearCategory,
	PotionCategory,
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
func CreateObject(obj Object) (Object, error) {
	// Validate the category.
	if !CategoryExists(obj.Category) {
		return Object{}, fmt.Errorf("category '%s' does not exist", obj.Category)
	}

	// Create a new Object with the same properties.
	newObject := Object{
		Name:     obj.Name,
		Category: obj.Category,
		Weight:   obj.Weight,
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
