package newtheme

import (
	"testing"

	"fyne.io/fyne/v2/theme"
)

func TestCustomTheme_Size_WithCustomValues(t *testing.T) {
	customTheme := CustomTheme{}
	sizeName := theme.SizeNameInnerPadding

	resultSize := customTheme.Size(sizeName)

	if resultSize != float32(8) {
		t.Errorf("CustomTheme.Size(%v) returned %v, expected %v", sizeName, resultSize, float32(8))
	}
}
