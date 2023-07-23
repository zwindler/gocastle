package model

import (
	"fmt"

	"fyne.io/fyne/v2/canvas"
)

type CharacterStats struct {
	// Character personalisation
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

	// Secondary charateristics
	// Those characteristics depend on main chars, level and gear
	MaxHP          int
	CurrentHP      int
	MaxMP          int
	CurrentMP      int
	PhysicalDamage int
	Armor          int

	// Inventory
	CurrentGold int
	Inventory   []Object
}

var (
	PlayerAvatar = Avatar{
		CanvasImage: canvas.NewImageFromFile("./static/warrior.png"),
	}
	Player = CharacterStats{
		// temporary, for dev
		CharacterName: "zwindler",
		GenderValue:   "Female",
		PointsToSpend: 0,
		// end temporary for dev
		Avatar: PlayerAvatar,
		//PointsToSpend:     10,
		StrengthValue:     10,
		ConstitutionValue: 10,
		IntelligenceValue: 10,
		DexterityValue:    10,
		Level:             1,
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

	// basic base secondary characteristics
	baseHP             = 8
	baseMP             = 8
	basePhysicalDamage = 2
)

// GetMaxHP changes player maxHP depending of player's level and constitution
func (player *CharacterStats) GetMaxHP() {
	// 8 + 4 by level +
	// bonus point for every 3 constitution point above 10 every level
	maxHP := baseHP + (4 * (player.Level - 1)) + (player.ConstitutionValue-10)/3*player.Level
	player.MaxHP = int(maxHP)
}

// GetMaxMP changes player maxMP depending of player's level and intelligence
func (player *CharacterStats) GetMaxMP() {
	// 8 + 4 by level +
	// bonus point for every 3 intelligence point above 10 every level
	maxMP := baseMP + (4 * (player.Level - 1)) + (player.IntelligenceValue-10)/3*player.Level
	player.MaxMP = int(maxMP)
}

// DeterminePhysicalDamage changes physicalDamage stat depending on str, dex and gear
func (player *CharacterStats) DeterminePhysicalDamage() {
	damage := basePhysicalDamage + (player.StrengthValue-10)/5*2 + (player.DexterityValue-10)/5*2

	// search in inventory items modifying the physicalDamage
	for _, item := range player.Inventory {
		for _, stat := range item.Stats {
			if stat.Name == "physicalDamage" {
				damage += stat.Modifier
			}
		}
	}

	fmt.Printf("damage %d\n", damage)
	player.PhysicalDamage = int(damage)
}

// ChangeXP changes XP player from XPAmount, could be negative, return true if leveled up
func (player *CharacterStats) ChangeXP(XPAmount int) bool {
	player.CurrentXP = player.CurrentXP + XPAmount
	// Since we change XP, check if level changes
	return player.DetermineLevel()
}

// ChangeGold changes amount of gold of player from GoldAmount, could be negative
func (player *CharacterStats) ChangeGold(GoldAmount int) {
	// TODO: add some random elements
	player.CurrentGold = int(player.CurrentGold) + GoldAmount
}

// DetermineLevel check player currentXP and increase level if necessary
// You can't loose levels even if you lost XP (by design). Returns true if
// player leveled up
func (player *CharacterStats) DetermineLevel() bool {
	for i, requiredXP := range xpTable {
		//fmt.Printf("Current XP %d", player.CurrentXP)
		if player.CurrentXP >= requiredXP {
			// we are still above threshold, continue
			//fmt.Printf("%d continue", i)
			continue
		} else {
			//fmt.Printf("%d stop", i)
			// we are bellow next threshold, that's our level
			if i > player.Level {
				// only change level if it's greater than current
				// there could be effects removing XP but I don't want to affect level

				// increase PointsToSpend by 2 per new levels
				// it shouldn't be more than 2 points each time when the game is finished/balanced
				player.PointsToSpend = player.PointsToSpend + 2*(i-player.Level)

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

// CollideWithPlayer returns true if we are going to collide with player, false instead
func (playerAvatar *Avatar) CollideWithPlayer(futurePosX int, futurePosY int) bool {
	return (playerAvatar.PosX == futurePosX && playerAvatar.PosY == futurePosY)
}

// RefreshStats is used when characters stats are modified, which in turn
// changes basic stats for player. If heal is true, reset HP/MP to 100%max
func (player *CharacterStats) RefreshStats(heal bool) {
	// Max HP changes during level up
	player.GetMaxHP()
	// Max MP changes during level up, also reset MP player
	player.GetMaxMP()
	// base damage may evolve when you can add char points
	player.DeterminePhysicalDamage()

	if heal {
		player.CurrentHP = player.MaxHP
		player.CurrentMP = player.MaxMP
	}
}

// AddObjectToInventory adds an object to the player's inventory.
func (player *CharacterStats) AddObjectToInventory(obj Object) {
	player.Inventory = append(player.Inventory, obj)
}

// RemoveObjectFromInventory removes an object from the player's inventory by its index.
func (player *CharacterStats) RemoveObjectFromInventory(index int) {
	if index >= 0 && index < len(player.Inventory) {
		player.Inventory = append(player.Inventory[:index], player.Inventory[index+1:]...)
	}
}
