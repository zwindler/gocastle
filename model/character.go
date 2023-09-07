// models/character.go

package model

import (
	"fmt"

	"github.com/zwindler/gocastle/pkg/hp"
	"github.com/zwindler/gocastle/pkg/mp"
)

type CharacterStats struct {
	// Character personalization
	CharacterName string
	GenderValue   string

	// Avatar
	Avatar Avatar

	// Main characteristics
	PointsToSpend     int
	StrengthValue     int
	ConstitutionValue int
	IntelligenceValue int
	DexterityValue    int

	// XP + levels
	CurrentXP int
	Level     int

	// Secondary characteristics
	// Those characteristics depend on main chars, level and gear
	HP             hp.HP
	MP             *mp.MP
	PhysicalDamage int
	Armor          int

	// Inventory
	Inventory       []*Object
	CurrentGold     int
	InventoryWeight int // weight of all player's inventory in grams
	EquippedWeight  int // same thing for equipped items only
}

var (
	PlayerAvatar = Avatar{}
	Player       = CharacterStats{
		// temporary, for dev
		CharacterName: "zwindler",
		PointsToSpend: 0,
		// end temporary for dev
		Avatar: PlayerAvatar,
		// PointsToSpend:     10,
		StrengthValue:     10,
		ConstitutionValue: 10,
		IntelligenceValue: 10,
		DexterityValue:    10,
		Level:             1,
		HP:                hp.New(baseHP),
		MP:                mp.New(baseMP),
	}
	xpTable = []int{
		0, // Level 1
		100,
		300,
		600,
		1000, // Level 5
		1500,
		2100,
		2800,
		3600,
		4500, // Level 10
	}

	// basic base secondary characteristics.
	baseHP             = 8
	baseMP             = 8
	basePhysicalDamage = 2
)

// DeterminePhysicalDamage changes physicalDamage stat depending on str, dex and gear.
func (player *CharacterStats) DeterminePhysicalDamage() {
	damage := basePhysicalDamage + (player.StrengthValue-10)/5*2 + (player.DexterityValue-10)/5*2

	// search in inventory items modifying the physicalDamage
	for _, item := range player.Inventory {
		if item.Equipped {
			for _, stat := range item.Stats {
				if stat.Name == "physicalDamage" {
					damage += stat.Modifier
				}
			}
		}
	}

	player.PhysicalDamage = damage
}

// ChangeXP changes XP player from XPAmount, could be negative, return true if leveled up.
func (player *CharacterStats) ChangeXP(xpAmount int) bool {
	player.CurrentXP += xpAmount
	// Since we change XP, check if level changes
	return player.DetermineLevel()
}

// ChangeGold changes amount of gold of player from GoldAmount, could be negative.
func (player *CharacterStats) ChangeGold(goldAmount int) {
	// TODO: add some random elements
	player.CurrentGold += goldAmount
	player.ComputeWeight()
}

// DetermineLevel check player currentXP and increase level if necessary
// You can't loose levels even if you lost XP (by design). Returns true if
// player leveled up.
func (player *CharacterStats) DetermineLevel() bool {
	for i, requiredXP := range xpTable {
		if player.CurrentXP >= requiredXP {
			// we are still above threshold, continue

			continue
		} else {
			// we are bellow next threshold, that's our level
			if i > player.Level {
				// only change level if it's greater than current
				// there could be effects removing XP but I don't want to affect level

				// increase PointsToSpend by 2 per new levels
				// it shouldn't be more than 2 points each time when the game is finished/balanced
				player.PointsToSpend += 2 * (i - player.Level)

				// set new level
				player.Level = i
				player.RefreshStats(true)

				return true
			}
			break
		}
	}
	return false
}

// CollideWithPlayer returns true if we are going to collide with player, false instead.
func (subject *Avatar) CollideWithPlayer(futurePosX, futurePosY int) bool {
	return (subject.Coord.X == futurePosX && subject.Coord.Y == futurePosY)
}

// RefreshStats is used when characters stats are modified, which in turn
// changes basic stats for player. If heal is true, reset HP/MP to 100%max.
func (player *CharacterStats) RefreshStats(heal bool) {
	// Max HP changes during level up
	player.HP.Compute(player.Level, baseHP, player.ConstitutionValue)
	// player.GetMaxHP()
	// Max MP changes during level up, also reset MP player
	player.MP.Compute(player.Level, baseMP, player.IntelligenceValue)
	// base damage may evolve when you can add char points
	player.DeterminePhysicalDamage()

	if heal {
		player.HP.Current.Set(player.HP.Max.Get())
		player.MP.Current.Set(player.MP.Max.Get())
	}
}

// AddObjectToInventory adds an object to the player's inventory.
func (player *CharacterStats) AddObjectToInventory(obj *Object, equip bool) {
	player.Inventory = append(player.Inventory, obj)
	obj.InInventory = true
	obj.Equipped = equip

	player.ComputeWeight()
}

// RemoveObjectFromInventory removes an object from the player's inventory by its index.
func (player *CharacterStats) RemoveObjectFromInventory(index int) {
	if index >= 0 && index < len(player.Inventory) {
		player.Inventory = append(player.Inventory[:index], player.Inventory[index+1:]...)
	}
	player.ComputeWeight()
}

// EquipItem equips an item in the player's inventory.
// If the item is already equipped or the category doesn't exist, it returns an error.
// If another item of the same category is equipped, un-equip it.
func (player *CharacterStats) EquipItem(item *Object) error {
	if !CategoryExists(item.Category) {
		return fmt.Errorf("category '%s' does not exist", item.Category)
	}
	if item.Equipped {
		return fmt.Errorf("item '%s' is already equipped", item.Name)
	}

	// Check if there is already an equipped item in the same category
	// if true, un-equip it
	for i, otherItem := range player.Inventory {
		if otherItem.Equipped && otherItem.Category == item.Category {
			// Un-equip the already equipped item
			player.Inventory[i].Equipped = false
			break
		}
	}

	// Equip the selected item
	item.Equipped = true
	player.ComputeWeight()
	return nil
}

// UnequipItem un-equips an item in the player's inventory.
func (player *CharacterStats) UnequipItem(item *Object) {
	item.Equipped = false
	player.ComputeWeight()
}

// ComputeWeight computes the total weight of the player's inventory and equipped items weight.
func (player *CharacterStats) ComputeWeight() {
	totalWeight := 0
	equippedWeight := 0
	for _, item := range player.Inventory {
		totalWeight += item.Weight
		if item.Equipped {
			equippedWeight += item.Weight
		}
	}

	// add gold to inventory weight (1g / piece)
	player.InventoryWeight = totalWeight + player.CurrentGold
	player.EquippedWeight = equippedWeight
}

func (player *CharacterStats) DeduceGenderFromAspect(index int) {
	switch index % 3 {
	case 0:
		player.GenderValue = "Female"
	case 1:
		player.GenderValue = "Non Binary"
	default:
		player.GenderValue = "Male"
	}
}
