package model

import (
	"fyne.io/fyne/v2/canvas"
)

type CharacterStats struct {
	CharacterName     string
	GenderValue       string
	Avatar            Avatar
	PointsToSpend     int
	StrengthValue     int
	ConstitutionValue int
	IntelligenceValue int
	DexterityValue    int
	Level             int
	MaxHP             int
	CurrentHP         int
	MaxMP             int
	CurrentMP         int
	BaseDamage        int
	CurrentXP         int
	CurrentGold       int
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
	baseHP = 8
	baseMP = 8
)

func (player *CharacterStats) GetMaxHP() {
	// 8 + 4 by level +
	// bonus point for every 3 constitution point above 10 every level
	maxHP := baseHP + (4 * (player.Level - 1)) + (player.ConstitutionValue-10)/3*player.Level
	player.MaxHP = int(maxHP)
}

func (player *CharacterStats) GetMaxMP() {
	// 8 + 4 by level +
	// bonus point for every 3 intelligence point above 10 every level
	maxMP := baseMP + (4 * (player.Level - 1)) + (player.IntelligenceValue-10)/3*player.Level
	player.MaxMP = int(maxMP)
}

func (player *CharacterStats) DetermineBaseDamage() {
	baseDamage := 4 + (player.StrengthValue-10)/5*2 + (player.DexterityValue-10)/5*2
	player.BaseDamage = int(baseDamage)
}

// change XP player from XPAmount, could be negative, return true if leveled up
func (player *CharacterStats) ChangeXP(XPAmount int) bool {
	player.CurrentXP = player.CurrentXP + XPAmount
	// Since we change XP, check if level changes
	return player.DetermineLevel()
}

// change amount of gold of player from GoldAmount, could be negative
func (player *CharacterStats) ChangeGold(GoldAmount int) {
	// TODO: add some random elements
	player.CurrentGold = int(player.CurrentGold) + GoldAmount
}

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

				player.RefreshStats()

				return true
			}
			break
		}
	}
	return false
}

// returns true if we are going to collide with player, false instead
func (playerAvatar *Avatar) CollideWithPlayer(futurePosX int, futurePosY int) bool {
	return (playerAvatar.PosX == futurePosX && playerAvatar.PosY == futurePosY)
}

func (player *CharacterStats) RefreshStats() {
	// Max HP changes during level up, also heal player
	player.GetMaxHP()
	player.CurrentHP = player.MaxHP

	// Max MP changes during level up, also reset MP player
	player.GetMaxMP()
	player.CurrentMP = player.MaxMP

	// base damage may evolve when you can add char points
	player.DetermineBaseDamage()
}
