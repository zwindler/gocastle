package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowNewGameScreen(window fyne.Window) {
	characterNameLabel := widget.NewLabel("Character's name")
	characterNameEntry := widget.NewEntry()

	pointsToSpendLabel := widget.NewLabel("Points to spend")
	pointsToSpendValue := widget.NewLabel("10")

	strengthLabel := widget.NewLabel("Strength: 10")
	strengthRange := widget.NewSlider(5, 50)
	strengthRange.Value = 10
	strengthRange.OnChanged = func(v float64) {
		strengthLabel.SetText(fmt.Sprintf("Strength: %.0f", v))
		strengthLabel.Refresh()
	}

	constitutionLabel := widget.NewLabel("Constitution: 10")
	constitutionRange := widget.NewSlider(5, 50)
	constitutionRange.Value = 10
	constitutionRange.OnChanged = func(v float64) {
		constitutionLabel.SetText(fmt.Sprintf("Constitution: %.0f", v))
		constitutionLabel.Refresh()
	}

	intelligenceLabel := widget.NewLabel("Intelligence: 10")
	intelligenceRange := widget.NewSlider(5, 50)
	intelligenceRange.Value = 10
	intelligenceRange.OnChanged = func(v float64) {
		intelligenceLabel.SetText(fmt.Sprintf("Intelligence: %.0f", v))
		intelligenceLabel.Refresh()
	}

	dexterityLabel := widget.NewLabel("Dexterity: 10")
	dexterityRange := widget.NewSlider(5, 50)
	dexterityRange.Value = 10
	dexterityRange.OnChanged = func(v float64) {
		dexterityLabel.SetText(fmt.Sprintf("Dexterity: %.0f", v))
		dexterityLabel.Refresh()
	}

	backButton := widget.NewButton("Back", func() {
		ShowMenuScreen(window)
	})
	validateButton := widget.NewButton("Validate", func() {
	})

	firstLine := container.NewHBox(
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
