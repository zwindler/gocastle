package embedmaps

import (
	"testing"
)

func TestGetMapMatrixFromEmbed(t *testing.T) {
	tcs := []struct {
		path              string
		expectErr         bool
		expectOutput      [][]uint16
		expectWrongOutput bool
	}{
		{"maps/99.json", false, [][]uint16{{1, 2, 3}, {4, 5, 6}}, false},
		{"maps/99.json", false, [][]uint16{{0, 0, 0}, {4, 5, 6}}, true},
		{"dontexist", true, [][]uint16{}, false},
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
func compareMapMatrix(a, b [][]uint16) bool {
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
