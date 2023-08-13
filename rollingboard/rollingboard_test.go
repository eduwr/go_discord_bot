package rollingboard

import (
	"fmt"
	"testing"
)

type Dice struct {
	// Define your Dice type here or import it from the correct package
}

func TestNewRollingBoard(t *testing.T) {
	testCases := []struct {
		input    string
		expected error
	}{
		{"!roll 2d6", nil},
		{"!roll 3d10+1d4", nil},
		{"!roll invalid", fmt.Errorf("invalid dice notation")},
		{"!roll 3d1", fmt.Errorf("invalid dice size")},
	}

	for _, tc := range testCases {
		_, err := NewRollingBoard(tc.input)

		if (err == nil && tc.expected != nil) || (err != nil && tc.expected == nil) {
			t.Errorf("Expected error '%v' but got '%v'", tc.expected, err)
		}

		if err != nil && tc.expected != nil && err.Error() != tc.expected.Error() {
			t.Errorf("Expected error '%v' but got '%v'", tc.expected, err)
		}
	}
}
