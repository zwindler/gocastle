package model

import "fyne.io/fyne/v2/canvas"

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
	CurrentGold       float32
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
)

func GetMaxHP(level int, baseHP int, constitution int) int {
	// 8 + 4 by level +
	// bonus point for every 3 constitution point above 10 every level
	maxHP := baseHP + (4 * (level - 1)) + (constitution-10)/3*level
	return int(maxHP)
}

func GetMaxMP(level int, baseMP int, intelligence int) int {
	// 8 + 4 by level +
	// bonus point for every 3 intelligence point above 10 every level
	maxMP := baseMP + (4 * (level - 1)) + (intelligence-10)/3*level
	return int(maxMP)
}

func DetermineBaseDamage(strength int, dexterity int) int {
	baseDamage := 4 + (strength-10)/5*2 + (dexterity-10)/5*2
	return int(baseDamage)
}
