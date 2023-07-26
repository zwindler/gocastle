package screens

import (
	"gocastle/maps"
	"gocastle/model"
	"log"

	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	player     = &model.Player
	currentMap = maps.Village
)

const (
	minStat = 5
	maxStat = 20
)

// ShowNewGameScreen is the main function of the new game screen
func ShowNewGameScreen(window fyne.Window) {
	var characterAspect1 *widget.RadioGroup
	var characterAspect2 *widget.RadioGroup
	var characterAspect3 *widget.RadioGroup

	characterNameLabel := widget.NewLabel("Character's name")
	characterNameEntry := widget.NewEntry()
	// temporary, for dev
	characterNameEntry.Text = "zwindler"

	pointsToSpendLabel := widget.NewLabel("Remaining points")
	pointsToSpendValue := widget.NewLabel(fmt.Sprintf("%d", model.Player.PointsToSpend))

	strengthLabel := widget.NewLabel(fmt.Sprintf("Strength: %d", model.Player.StrengthValue))
	strengthRange := createSliderWithCallback("Strength", minStat, maxStat,
		&model.Player.StrengthValue, &model.Player.PointsToSpend,
		strengthLabel, pointsToSpendValue)

	constitutionLabel := widget.NewLabel(fmt.Sprintf("Constitution: %d", model.Player.ConstitutionValue))
	constitutionRange := createSliderWithCallback("Constitution", minStat, maxStat,
		&model.Player.ConstitutionValue, &model.Player.PointsToSpend,
		constitutionLabel, pointsToSpendValue)

	intelligenceLabel := widget.NewLabel(fmt.Sprintf("Intelligence: %d", model.Player.IntelligenceValue))
	intelligenceRange := createSliderWithCallback("Intelligence", minStat, maxStat,
		&model.Player.IntelligenceValue, &model.Player.PointsToSpend,
		intelligenceLabel, pointsToSpendValue)

	dexterityLabel := widget.NewLabel(fmt.Sprintf("Dexterity: %d", model.Player.DexterityValue))
	dexterityRange := createSliderWithCallback("Dexterity", minStat, maxStat,
		&model.Player.DexterityValue, &model.Player.PointsToSpend,
		dexterityLabel, pointsToSpendValue)

	characterGenderLabel := widget.NewLabel("Gender")
	genderRadioButton := widget.NewRadioGroup([]string{"Female", "Male", "Non-binary"}, func(selected string) {
		model.Player.GenderValue = selected
	})

	characterAspectLabel := widget.NewLabel("Aspect")
	characterAspect1 = widget.NewRadioGroup([]string{"👩‍🦰", "👨‍🦰", "🧑‍🦰", "👱‍♀️", "👱‍♂️", "👱"}, func(selected string) {
		resetRadioGroups(characterAspect2, characterAspect3)
		//model.Player.AspectValue = selected
	})

	characterAspect2 = widget.NewRadioGroup([]string{"👩‍🦱", "👨‍🦱", "🧑‍🦱", "🧕", "👳‍♂️", "👳"}, func(selected string) {
		resetRadioGroups(characterAspect1, characterAspect3)
		//model.Player.AspectValue = selected
	})

	characterAspect3 = widget.NewRadioGroup([]string{"👩‍🦳", "👨‍🦳", "🧑‍🦳", "👩‍🦲", "👨‍🦲", "🧑‍🦲"}, func(selected string) {
		resetRadioGroups(characterAspect1, characterAspect2)
		//model.Player.AspectValue = selected
	})

	backButton := widget.NewButton("Back", func() {
		ShowMenuScreen(window)
	})
	validateButton := widget.NewButton("Validate", func() {
		if characterNameEntry.Text == "" {
			content := widget.NewLabel("You still have to choose a name for you character!")
			dialog.ShowCustom("Character has no name", "Close", content, window)
		} else {
			model.Player.CharacterName = characterNameEntry.Text
			if model.Player.PointsToSpend > 0 {
				content := widget.NewLabel("You still have available characteristics point to allocate!")
				dialog.ShowCustom("Points still available", "Close", content, window)
			} else {
				if model.Player.GenderValue == "" {
					content := widget.NewLabel("Character has no gender, please choose one")
					dialog.ShowCustom("Gender not selected", "Close", content, window)
				} else {
					/*if model.Player.AspectValue == nil {
					content := widget.NewLabel("Character has no aspect, please choose one")
					dialog.ShowCustom("Aspect not selected", "Close", content, window)
					} else {*/

					// we are good to go!
					player.Avatar.PosX = currentMap.PlayerStart.X
					player.Avatar.PosY = currentMap.PlayerStart.Y
					currentMap.AddNPCs()

					// create a knife, add it to player's inventory, equip it
					// TODO rework later
					knife, err := model.CreateObject(model.HuntingKnife)
					if err != nil {
						err = fmt.Errorf("unable to add knife to inventory: %w", err)
						log.Fatalf("NewGame error: %s", err)
					}
					knifeIndex := player.AddObjectToInventory(knife)
					player.EquipItem(knifeIndex)

					player.RefreshStats(true)

					ShowGameScreen(window)
					//}
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
		characterGenderLabel, characterAspectLabel, layout.NewSpacer(), layout.NewSpacer())

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

// createSliderWithCallback is the callback function for sliders in newgame screen
func createSliderWithCallback(characteristic string, min float64, max float64,
	value *int, pointsToSpend *int,
	valueLabel, pointsToSpendLabel *widget.Label) *widget.Slider {
	slider := widget.NewSlider(min, max)
	slider.Value = float64(*value)
	slider.OnChanged = func(v float64) {
		intV := int(v)
		if (model.Player.PointsToSpend - (intV - *value)) >= 0 {
			model.Player.PointsToSpend = model.Player.PointsToSpend - (intV - *value)
			*value = intV
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

// resetRadioGroups is a helper function for resetting unselected radio groups
func resetRadioGroups(groups ...*widget.RadioGroup) {
	for _, group := range groups {
		group.SetSelected("")
	}
}
