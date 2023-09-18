package embed

import (
	"image"
	"testing"
)

func TestGetImageFromEmbed(t *testing.T) {
	tests := []struct {
		path       string
		expectErr  bool
		expectSize image.Point
	}{
		{"static/transparent_tile.png", false, image.Point{32, 32}}, // Existing file
		{"non-existent.png", true, image.Point{0, 0}},               // Non-existent file, should error
	}

	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			img, err := GetImageFromEmbed(test.path)

			if test.expectErr {
				if err == nil {
					t.Errorf("Expected an error, but none occurred")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if !test.expectErr && img != nil && img.Bounds().Size() != test.expectSize {
				t.Errorf("Image size mismatch. Expected %v, got %v", test.expectSize, img.Bounds().Size())
			}
		})
	}
}
