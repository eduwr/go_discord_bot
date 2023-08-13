package dice

import (
	"errors"
	"math/rand"
)

type Dice struct {
	size int8
}

func NewDice(size int8) (*Dice, error) {
	if size <= 1 {
		return nil, errors.New("invalid dice size")
	}

	return &Dice{size: size}, nil
}

func getRandomNbr(s int8) int8 {
	return int8(rand.Intn(int(s)))
}

func (d Dice) Roll() int8 {
	return getRandomNbr(d.size) + 1
}
