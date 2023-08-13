package dice

import (
	"errors"
	"math/rand"
)

type dice struct {
	size int8
}

func NewDice(size int8) (*dice, error) {
	if size <= 1 {
		return nil, errors.New("invalid dice size")
	}

	return &dice{size: size}, nil
}

func getRandomNbr(s int8) int8 {
	return int8(rand.Intn(int(s)))
}

func (d dice) Roll() int8 {
	r := getRandomNbr(d.size + 1)
	for r == 0 {
		r = getRandomNbr(d.size + 1)
	}
	return r
}
