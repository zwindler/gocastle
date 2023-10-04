package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/zwindler/gocastle/model"
)

// showLevelUpScreen is the main function for the level-up screen.
func showLevelUpScreen() *fyne.Container {
	pointsToSpendLabel := widget.NewLabel("Remaining points")
	pointsToSpendValue := widget.NewLabel(fmt.Sprintf("%d", model.Player.PointsToSpend))

	currentStrength := model.Player.StrengthValue
	currentConstitution := model.Player.ConstitutionValue
	currentIntelligence := model.Player.IntelligenceValue
	currentDexterity := model.Player.DexterityValue

	strengthLabel := widget.NewLabel(fmt.Sprintf("Strength: %d", model.Player.StrengthValue))
	strengthRange := createSliderLevelUpWithCallback("Strength", minStat, maxStat,
		&model.Player.StrengthValue, currentStrength, strengthLabel, pointsToSpendValue)

	constitutionLabel := widget.NewLabel(fmt.Sprintf("Constitution: %d", model.Player.ConstitutionValue))
	constitutionRange := createSliderLevelUpWithCallback("Constitution", minStat, maxStat,
		&model.Player.ConstitutionValue, currentConstitution, constitutionLabel, pointsToSpendValue)

	intelligenceLabel := widget.NewLabel(fmt.Sprintf("Intelligence: %d", model.Player.IntelligenceValue))
	intelligenceRange := createSliderLevelUpWithCallback("Intelligence", minStat, maxStat,
		&model.Player.IntelligenceValue, currentIntelligence, intelligenceLabel, pointsToSpendValue)

	dexterityLabel := widget.NewLabel(fmt.Sprintf("Dexterity: %d", model.Player.DexterityValue))
	dexterityRange := createSliderLevelUpWithCallback("Dexterity", minStat, maxStat,
		&model.Player.DexterityValue, currentDexterity, dexterityLabel, pointsToSpendValue)

	return container.New(layout.NewGridLayout(5),
		pointsToSpendLabel, strengthLabel, constitutionLabel, intelligenceLabel, dexterityLabel,
		pointsToSpendValue, strengthRange, constitutionRange, intelligenceRange, dexterityRange)
}

// createSliderLevelUpWithCallback is the callback function for characteristics sliders.
// _ parameter is pointsToSpend because we don't need it here.
func createSliderLevelUpWithCallback(characteristic string, min, max float64, //nolint:unparam // TODO: min is a constant
	value *int, currentValue int, valueLabel, pointsToSpendLabel *widget.Label,
) *widget.Slider {
	slider := widget.NewSlider(min, max)
	slider.Value = float64(*value)
	slider.OnChanged = func(v float64) {
		intV := int(v)
		if (model.Player.PointsToSpend - (intV - *value)) >= 0 {
			// player still has enough point to spend to make this modification
			// however, this could mean that player wants to remove points allocated
			// to characteristics from previous levels, which we don't want

			// we only allow modification if new value is greater or equal than current value
			if intV >= currentValue {
				model.Player.PointsToSpend -= (intV - *value)
				*value = intV
			}
		} else {
			slider.Value = float64(*value)
			slider.Refresh()
		}
		valueLabel.SetText(fmt.Sprintf("%s: %d", characteristic, *value))
		pointsToSpendLabel.SetText(fmt.Sprintf("%d", model.Player.PointsToSpend))
		valueLabel.Refresh()
	}
	return slider
}
