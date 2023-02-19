package grains

import (
	"errors"
)

const (
	maxNumSquares = 64
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > maxNumSquares {
		return 0, errors.New("invalid number of squares.")
	}
	return 1 << (number - 1), nil
}

func Total() uint64 {
	return 1<<maxNumSquares - 1
}
