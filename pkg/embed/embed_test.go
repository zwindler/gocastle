package embed

import (
	"image"
	"testing"
)

func TestGetImageFromEmbed(t *testing.T) {
	tcs := []struct {
		path       string
		expectErr  bool
		expectSize image.Point
	}{
		{"static/transparent_tile.png", false, image.Point{32, 32}}, // Existing file
		{"non-existent.png", true, image.Point{0, 0}},               // Non-existent file, should error
	}

	for _, tc := range tcs {
		t.Run(tc.path, func(t *testing.T) {
			img, err := GetImageFromEmbed(tc.path)

			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected an error, but none occurred")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if !tc.expectErr && img != nil && img.Bounds().Size() != tc.expectSize {
				t.Errorf("Image size mismatch. Expected %v, got %v", tc.expectSize, img.Bounds().Size())
			}
		})
	}
}

func TestGetMapMatrixFromEmbed(t *testing.T) {
	tcs := []struct {
		path              string
		expectErr         bool
		expectOutput      [][]int
		expectWrongOutput bool
	}{
		{"maps/99.json", false, [][]int{{1, 2, 3}, {4, 5, 6}}, false},
		{"maps/99.json", false, [][]int{{0, 0, 0}, {4, 5, 6}}, true},
		{"dontexist", true, [][]int{}, false},
	}

	for _, tc := range tcs {
		currentMap, err := GetMapMatrixFromEmbed(tc.path)

		if tc.expectErr {
			if err == nil {
				t.Errorf("Expected an error, but none occurred")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !compareMapMatrix(currentMap, tc.expectOutput) && !tc.expectWrongOutput {
				t.Errorf("MapMatrix mismatch. Expected %v, got %v", tc.expectOutput, currentMap)
			}
		}
	}
}

// compareMapMatrix is a function that checks that 2 MapMatrix are identical.
func compareMapMatrix(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for row := range a {
		if len(a[row]) != len(b[row]) {
			return false
		}
		for column := range a[row] {
			if a[row][column] != b[row][column] {
				return false
			}
		}
	}
	return true
}
