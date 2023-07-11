package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	PointsToSpend     float64
	StrengthValue     float64
	ConstitutionValue float64
	IntelligenceValue float64
	DexterityValue    float64
)

func ShowNewGameScreen(window fyne.Window) {
	PointsToSpend = 10
	StrengthValue = 10
	ConstitutionValue = 10
	IntelligenceValue = 10
	DexterityValue = 10

	characterNameLabel := widget.NewLabel("Character's name")
	characterNameEntry := widget.NewEntry()

	pointsToSpendLabel := widget.NewLabel("Points to spend")
	pointsToSpendValue := widget.NewLabel("10")

	strengthLabel := widget.NewLabel("Strength: 10")
	strengthRange := createSliderWithCallback("Strength", 5, 30,
		10, &StrengthValue, &PointsToSpend,
		strengthLabel, pointsToSpendValue)

	constitutionLabel := widget.NewLabel("Constitution: 10")
	constitutionRange := createSliderWithCallback("Constitution", 5, 30,
		10, &ConstitutionValue, &PointsToSpend,
		constitutionLabel, pointsToSpendValue)

	intelligenceLabel := widget.NewLabel("Intelligence: 10")
	intelligenceRange := createSliderWithCallback("Intelligence", 5, 30,
		10, &IntelligenceValue, &PointsToSpend,
		intelligenceLabel, pointsToSpendValue)

	dexterityLabel := widget.NewLabel("Dexterity: 10")
	dexterityRange := createSliderWithCallback("Dexterity", 5, 30,
		10, &DexterityValue, &PointsToSpend,
		dexterityLabel, pointsToSpendValue)

	backButton := widget.NewButton("Back", func() {
		ShowMenuScreen(window)
	})
	validateButton := widget.NewButton("Validate", func() {
	})

	firstLine := container.New(layout.NewFormLayout(),
		characterNameLabel,
		characterNameEntry,
	)

	slidersLine := container.New(layout.NewGridLayout(5),
		pointsToSpendLabel, strengthLabel, constitutionLabel, intelligenceLabel, dexterityLabel,
		pointsToSpendValue, strengthRange, constitutionRange, intelligenceRange, dexterityRange)

	lastLine := container.NewHBox(
		backButton,
		validateButton,
	)

	content := container.NewVBox(
		firstLine,
		slidersLine,
		lastLine,
	)

	window.SetContent(content)
}

func createSliderWithCallback(characteristic string, min float64, max float64,
	defaultValue float64, value *float64, pointsToSpend *float64,
	valueLabel, pointsToSpendLabel *widget.Label) *widget.Slider {
	slider := widget.NewSlider(min, max)
	slider.Value = defaultValue
	slider.OnChanged = func(v float64) {
		if (*pointsToSpend - (v - *value)) >= 0 {
			*pointsToSpend = *pointsToSpend - (v - *value)
			*value = v
		} else {
			slider.Value = *value
			slider.Refresh()
		}
		valueLabel.SetText(fmt.Sprintf("%s: %.0f", characteristic, *value))
		pointsToSpendLabel.SetText(fmt.Sprintf("%.0f", *pointsToSpend))
		valueLabel.Refresh()
	}
	return slider
}
