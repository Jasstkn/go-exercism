package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	Cows int // number of cows
}

func (p SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", p.Cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	switch {
	case cows == 0:
		return 0, errors.New("division by zero")
	case err == ErrScaleMalfunction && fodder > 0:
		fodder *= 2
	case fodder < 0 && (err == ErrScaleMalfunction || err == nil):
		return 0, errors.New("negative fodder")
	case err != nil:
		return 0, err
	case cows < 0:
		return 0, &SillyNephewError{Cows: cows}
	}
	return fodder / float64(cows), nil

}
