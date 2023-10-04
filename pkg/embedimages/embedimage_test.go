package embedimages

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
