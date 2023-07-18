package model

type CharacterStats struct {
	CharacterName     string
	GenderValue       string
	AspectValue       string
	PointsToSpend     uint
	StrengthValue     uint
	ConstitutionValue uint
	IntelligenceValue uint
	DexterityValue    uint
	Level             uint
	MaxHP             uint
	CurrentHP         uint
}

var Player = CharacterStats{
	CharacterName:     "",
	GenderValue:       "Female",
	AspectValue:       ":-)",
	PointsToSpend:     10,
	StrengthValue:     10,
	ConstitutionValue: 10,
	IntelligenceValue: 10,
	DexterityValue:    10,
	Level:             1,
	MaxHP:             10,
	CurrentHP:         10,
}

func GetMaxHP(level uint, baseHP uint, constitution uint) uint {
	maxHP := baseHP + (4 * (level - 1)) + (constitution-10)/5*level
	return uint(maxHP)
}
