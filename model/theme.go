package model

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var TextColor color.Color

type CustomTheme struct{}

func (t CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (t CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t CustomTheme) Size(style fyne.ThemeSizeName) float32 {
	switch style {
	case theme.SizeNameInnerPadding:
		return 8
	case theme.SizeNamePadding:
		return 6
	case theme.SizeNameText:
		return 16
	default:
		return theme.DefaultTheme().Size(style)
	}
}
