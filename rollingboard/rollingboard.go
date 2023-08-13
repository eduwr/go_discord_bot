package rollingboard

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/eduwr/go_discord_bot/dice"
)

type RollingBoard struct {
	dices []dice.Dice
}

func NewRollingBoard(diceNotation string) (*RollingBoard, error) {
	regex := regexp.MustCompile(`!roll\s+([0-9]*d[0-9]+(\+[0-9]*d[0-9]+)*)`)
	match := regex.FindStringSubmatch(diceNotation)
	if len(match) < 2 {
		return nil, fmt.Errorf("invalid dice notation")
	}

	diceExpression := match[1]
	parts := strings.Split(diceExpression, "+")

	dices := []dice.Dice{}

	for _, part := range parts {
		subParts := strings.Split(part, "d")
		if len(subParts) == 2 {
			count, _ := strconv.Atoi(subParts[0])
			sides, _ := strconv.Atoi(subParts[1])

			for i := 0; i < count; i++ {
				d, err := dice.NewDice(int8(sides))

				if err != nil {
					return nil, err
				}
				dices = append(dices, *d)
			}
		}
	}

	return &RollingBoard{
		dices: dices,
	}, nil
}

func RollBoard(diceNotation string) (string, error) {
	regex := regexp.MustCompile(`!roll\s+([0-9]*d[0-9]+(\+[0-9]*d[0-9]+)*)`)
	match := regex.FindStringSubmatch(diceNotation)
	if len(match) < 2 {
		return "", fmt.Errorf("invalid dice notation")
	}

	diceExpression := match[1]
	parts := strings.Split(diceExpression, "+")
	total := 0

	for _, part := range parts {
		subParts := strings.Split(part, "d")
		if len(subParts) == 2 {
			count, _ := strconv.Atoi(subParts[0])
			sides, _ := strconv.Atoi(subParts[1])

			for i := 0; i < count; i++ {
				roll := 1 + rand.Intn(sides)
				total += roll
			}
		}
	}

	return strconv.Itoa(total), nil
}
