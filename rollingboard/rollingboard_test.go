package rollingboard

import (
	"fmt"
	"testing"

	"github.com/eduwr/go_discord_bot/dice"
)

type Dice struct {
	// Define your Dice type here or import it from the correct package
}

func TestNewRollingBoard(t *testing.T) {
	testCases := []struct {
		input           string
		expectedError   error
		expectedDiceNum int
	}{
		{"!roll 2d6", nil, 2},
		{"!roll 3d10+1d4", nil, 4},
		{"!roll invalid", fmt.Errorf("invalid dice notation"), 0},
		{"!roll 3d1", fmt.Errorf("invalid dice size"), 0},
	}

	for _, tc := range testCases {
		board, err := NewRollingBoard(tc.input)

		if tc.expectedError != nil {
			if err == nil {
				t.Errorf("expected error: `%v` but got `%v`", tc.expectedError, err)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: `%v` for input '%s'", err, tc.input)
			}

			if len(board.dices) != tc.expectedDiceNum {
				t.Errorf("Expected %d dice(s) but got %d for input `%s`", tc.expectedDiceNum, len(board.dices), tc.input)
			}

		}

	}
}

func TestRollDices(t *testing.T) {
	d6, _ := dice.NewDice(6)
	d8, _ := dice.NewDice(8)
	d10, _ := dice.NewDice(10)

	board := &RollingBoard{
		dices: []dice.Dice{
			*d6,
			*d8,
			*d10,
		},
	}

	total := board.RollDices()

	if total < int16(len(board.dices)) {
		t.Errorf("Expected total to be greater than 0, but got %d", total)
	}
}
