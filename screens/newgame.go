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
	genderValue       string // TODO maybe change this to allow only the 3 types?
	aspectValue       string // TODO maybe change this to allow only the allowed types?
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

	characterGenderLabel := widget.NewLabel("Gender")
	genderRadioButton := widget.NewRadioGroup([]string{"Female", "Male", "Non-binary"}, func(selected string) {})
	emptyLabel := widget.NewLabel(" ")

	characterAspectLabel := widget.NewLabel("Aspect")
	characterAspect1 = widget.NewRadioGroup([]string{"👩‍🦰", "👨‍🦰", "🧑‍🦰", "👱‍♀️", "👱‍♂️", "👱"}, func(selected string) {
		resetRadioGroups(characterAspect2, characterAspect3)
		aspectValue = selected
	})

	characterAspect2 = widget.NewRadioGroup([]string{"👩‍🦱", "👨‍🦱", "🧑‍🦱", "🧕", "👳‍♂️", "👳"}, func(selected string) {
		resetRadioGroups(characterAspect1, characterAspect3)
		aspectValue = selected
	})

	characterAspect3 = widget.NewRadioGroup([]string{"👩‍🦳", "👨‍🦳", "🧑‍🦳", "👩‍🦲", "👨‍🦲", "🧑‍🦲"}, func(selected string) {
		resetRadioGroups(characterAspect1, characterAspect2)
		aspectValue = selected
	})

	backButton := widget.NewButton("Back", func() {
		ShowMenuScreen(window)
	})
	validateButton := widget.NewButton("Validate", func() {
		if characterNameEntry.Text == "" {
			content := widget.NewLabel("You still have to choose a name for you character!")
			dialog.ShowCustom("Character has no name", "Close", content, window)
		} else {
			if pointsToSpend > 0 {
				content := widget.NewLabel("You still have available characteristics point to allocate!")
				dialog.ShowCustom("Points still available", "Close", content, window)
			} else {
				if genderRadioButton.Selected == "" {
					content := widget.NewLabel("Character has no gender, please choose one")
					dialog.ShowCustom("Gender not selected", "Close", content, window)
				} else {
					if aspectValue == "" {
						content := widget.NewLabel("Character has no aspect, please choose one")
						dialog.ShowCustom("Aspect not selected", "Close", content, window)
					}
				}
			}
		}
	})

	firstLine := container.New(layout.NewFormLayout(),
		characterNameLabel,
		characterNameEntry,
	)

	slidersLine := container.New(layout.NewGridLayout(5),
		pointsToSpendLabel, strengthLabel, constitutionLabel, intelligenceLabel, dexterityLabel,
		pointsToSpendValue, strengthRange, constitutionRange, intelligenceRange, dexterityRange)

	characterGenderAspectLabelLine := container.New(layout.NewGridLayout(4),
		characterGenderLabel, characterAspectLabel, emptyLabel, emptyLabel)

	characterGenderAspectLine := container.New(layout.NewGridLayout(4),
		genderRadioButton, characterAspect1, characterAspect2, characterAspect3)

	lastLine := container.NewHBox(
		backButton,
		validateButton,
	)

	content := container.NewVBox(
		firstLine,
		slidersLine,
		characterGenderAspectLabelLine,
		characterGenderAspectLine,
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
