package screens

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/zwindler/gocastle/model"
	"github.com/zwindler/gocastle/pkg/embed"
)

const (
	minStat = 5
	maxStat = 20
)

// ShowNewGameScreen is the main function of the new game screen.
func ShowNewGameScreen(window fyne.Window) {
	var (
		characterAspect1 *widget.RadioGroup
		characterAspect2 *widget.RadioGroup
		characterAspect3 *widget.RadioGroup
	)

	characterNameLabel := widget.NewLabelWithStyle("Character's name", 0, fyne.TextStyle{Bold: true, Italic: true})
	characterNameEntry := widget.NewEntry()
	// temporary, for dev
	characterNameEntry.Text = "zwindler"

	CharacteristicsLabel := widget.NewLabelWithStyle("Characteristics", 0, fyne.TextStyle{Bold: true, Italic: true})
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

	aspectIconPath := [][]string{
		{
			"static/red_haired_woman.png", "static/red_haired_person.png", "static/red_haired_man.png",
			"static/blond_haired_woman.png", "static/blond_haired_person.png", "static/blond_haired_man.png",
		},
		{
			"static/dark_haired_woman.png", "static/dark_haired_person.png", "static/dark_haired_man.png",
			"static/scarf_woman.png", "static/turban_person.png", "static/turban_man.png",
		},
		{
			"static/bald_woman.png", "static/bald_person.png", "static/bald_man.png",
			"static/white_haired_woman.png", "static/white_haired_person.png", "static/white_haired_man.png",
		},
	}

	characterAspectLabel := widget.NewLabelWithStyle("Aspect", 0, fyne.TextStyle{Bold: true, Italic: true})
	characterAspect1 = widget.NewRadioGroup([]string{"1", "2", "3", "4", "5", "6"}, func(selected string) {
		resetRadioGroups(characterAspect2, characterAspect3)
		if selected != "" {
			index, _ := strconv.Atoi(selected)
			// TODO deal with error
			player.Avatar.CanvasPath = aspectIconPath[0][index-1]
			player.DeduceGenderFromAspect(index)
		}
	})

	characterAspect2 = widget.NewRadioGroup([]string{"7", "8", "9", "10", "11", "12"}, func(selected string) {
		resetRadioGroups(characterAspect1, characterAspect3)
		if selected != "" {
			index, _ := strconv.Atoi(selected)
			// TODO deal with error
			player.Avatar.CanvasPath = aspectIconPath[1][index-7]
			player.DeduceGenderFromAspect(index)
		}
	})

	characterAspect3 = widget.NewRadioGroup([]string{"13", "14", "15", "16", "17", "18"}, func(selected string) {
		resetRadioGroups(characterAspect1, characterAspect2)
		if selected != "" {
			index, _ := strconv.Atoi(selected)
			// TODO deal with error
			player.Avatar.CanvasPath = aspectIconPath[2][index-13]
			player.DeduceGenderFromAspect(index)
		}
	})
	characterAspect3.Resize(fyne.NewSize(10, 10))

	backButton := widget.NewButton("Back", func() {
		ShowMenuScreen(window)
	})
	validateButton := widget.NewButton("Validate", func() {
		if characterNameEntry.Text == "" {
			content := widget.NewLabel("You still have to choose a name for you character!")
			dialog.ShowCustom("Character has no name", "Close", content, window)
		} else {
			player.CharacterName = characterNameEntry.Text
			if player.PointsToSpend > 0 {
				content := widget.NewLabel("You still have available characteristics point to allocate!")
				dialog.ShowCustom("Points still available", "Close", content, window)
			} else {
				if player.Avatar.CanvasPath == "" {
					content := widget.NewLabel("Character has no aspect, please choose one")
					dialog.ShowCustom("Aspect not selected", "Close", content, window)
				} else {
					// we are good to go!
					// initialise game objects for the first time
					initGame(window, true)
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

	var characterAspectTable []*fyne.Container
	for column := 0; column < 3; column++ {
		characterAspectTable = append(characterAspectTable, container.NewWithoutLayout())
		for row := 0; row < 6; row++ {
			image := canvas.NewImageFromImage(embed.GetImageFromEmbed(aspectIconPath[column][row]))
			image.FillMode = canvas.ImageFillOriginal
			image.Resize(fyneTileSize)
			currentPos := fyne.NewPos(0, float32(row)*38)
			image.Move(currentPos)
			characterAspectTable[column].Add(image)
		}
	}

	characterGenderAspectLine := container.New(layout.NewGridLayout(6),
		characterAspect1, characterAspectTable[0],
		characterAspect2, characterAspectTable[1],
		characterAspect3, characterAspectTable[2])

	lastLine := container.NewBorder(nil, nil, backButton, validateButton, nil)

	content := container.NewVBox(
		firstLine,
		CharacteristicsLabel,
		slidersLine,
		characterAspectLabel,
		characterGenderAspectLine,
	)

	mainContent = container.NewBorder(nil, lastLine, nil, nil, content)

	window.SetContent(mainContent)
}

// createSliderWithCallback is the callback function for sliders in newgame screen.
// _ parameter is pointsToSpend because we don't need it here.
func createSliderWithCallback(characteristic string, min, max float64, //nolint:unparam // TODO: min is a constant
	value, _ *int,
	valueLabel, pointsToSpendLabel *widget.Label,
) *widget.Slider {
	slider := widget.NewSlider(min, max)
	slider.Value = float64(*value)
	slider.OnChanged = func(v float64) {
		intV := int(v)
		if (model.Player.PointsToSpend - (intV - *value)) >= 0 {
			model.Player.PointsToSpend -= (intV - *value)
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

// resetRadioGroups is a helper function for resetting unselected radio groups.
func resetRadioGroups(groups ...*widget.RadioGroup) {
	for _, group := range groups {
		group.SetSelected("")
	}
}
