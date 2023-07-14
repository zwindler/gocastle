package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	pointsToSpend     float64
	strengthValue     float64
	constitutionValue float64
	intelligenceValue float64
	dexterityValue    float64
)

func ShowNewGameScreen(window fyne.Window) {
	var characterAspect1 *widget.RadioGroup
	var characterAspect2 *widget.RadioGroup
	var characterAspect3 *widget.RadioGroup

	// set initial defaults
	pointsToSpend = 10
	strengthValue = 10
	constitutionValue = 10
	intelligenceValue = 10
	dexterityValue = 10

	characterNameLabel := widget.NewLabel("Character's name")
	characterNameEntry := widget.NewEntry()

	pointsToSpendLabel := widget.NewLabel("Remaining points")
	pointsToSpendValue := widget.NewLabel("10")

	strengthLabel := widget.NewLabel("Strength: 10")
	strengthRange := createSliderWithCallback("Strength", 5, 35,
		10, &strengthValue, &pointsToSpend,
		strengthLabel, pointsToSpendValue)

	constitutionLabel := widget.NewLabel("Constitution: 10")
	constitutionRange := createSliderWithCallback("Constitution", 5, 35,
		10, &constitutionValue, &pointsToSpend,
		constitutionLabel, pointsToSpendValue)

	intelligenceLabel := widget.NewLabel("Intelligence: 10")
	intelligenceRange := createSliderWithCallback("Intelligence", 5, 35,
		10, &intelligenceValue, &pointsToSpend,
		intelligenceLabel, pointsToSpendValue)

	dexterityLabel := widget.NewLabel("Dexterity: 10")
	dexterityRange := createSliderWithCallback("Dexterity", 5, 35,
		10, &dexterityValue, &pointsToSpend,
		dexterityLabel, pointsToSpendValue)

	backButton := widget.NewButton("Back", func() {
		ShowMenuScreen(window)
	})
	validateButton := widget.NewButton("Validate", func() {
		if pointsToSpend > 0 {
			content := widget.NewLabel("You still have available characteristics point to allocate!")
			dialog.ShowCustom("Points still available", "Close", content, window)
		}
	})

	characterGenderLabel := widget.NewLabel("Gender")
	genderRadioButton := widget.NewRadioGroup([]string{"Female", "Male", "Non-binary"}, func(selected string) {})

	characterAspect1 = widget.NewRadioGroup([]string{"👩‍🦰", "👨‍🦰", "🧑‍🦰", "👱‍♀️", "👱‍♂️", "👱"}, func(selected string) {
		resetRadioGroups(characterAspect2, characterAspect3)
		fmt.Println("Character Aspect 1:", selected)
	})

	characterAspect2 = widget.NewRadioGroup([]string{"👩‍🦱", "👨‍🦱", "🧑‍🦱", "🧕", "👳‍♂️", "👳"}, func(selected string) {
		resetRadioGroups(characterAspect1, characterAspect3)
		fmt.Println("Character Aspect 2:", selected)
	})

	characterAspect3 = widget.NewRadioGroup([]string{"👩‍🦳", "👨‍🦳", "🧑‍🦳", "👩‍🦲", "👨‍🦲", "🧑‍🦲"}, func(selected string) {
		resetRadioGroups(characterAspect1, characterAspect2)
		fmt.Println("Character Aspect 3:", selected)
	})

	firstLine := container.New(layout.NewFormLayout(),
		characterNameLabel,
		characterNameEntry,
	)

	slidersLine := container.New(layout.NewGridLayout(5),
		pointsToSpendLabel, strengthLabel, constitutionLabel, intelligenceLabel, dexterityLabel,
		pointsToSpendValue, strengthRange, constitutionRange, intelligenceRange, dexterityRange)

	characterGenderBox := container.New(layout.NewVBoxLayout(),
		characterGenderLabel,
		genderRadioButton,
	)

	characterAspectLine := container.New(layout.NewGridLayout(4),
		characterGenderBox, characterAspect1, characterAspect2, characterAspect3)

	lastLine := container.NewHBox(
		backButton,
		validateButton,
	)

	content := container.NewVBox(
		firstLine,
		slidersLine,
		characterAspectLine,
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

func resetRadioGroups(groups ...*widget.RadioGroup) {
	for _, group := range groups {
		group.SetSelected("")
	}
}
