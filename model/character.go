package model

type CharacterStats struct {
	CharacterName     string
	GenderValue       string
	AspectValue       string
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
}

var (
	Player = CharacterStats{
		// TODO remove
		CharacterName: "zwindler",
		GenderValue:   "Female",
		AspectValue:   ":-)",
		// endTODO
		PointsToSpend:     10,
		StrengthValue:     10,
		ConstitutionValue: 10,
		IntelligenceValue: 10,
		DexterityValue:    10,
		Level:             1,
		MaxHP:             8,
		CurrentHP:         8,
		MaxMP:             10,
		CurrentMP:         10,
	}
	PlayerDefaultHP = 10
)

func GetMaxHP(level int, baseHP int, constitution int) int {
	maxHP := baseHP + (4 * (level - 1)) + (constitution-10)/5*level
	return int(maxHP)
}
