package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, nCows int) (float64, error) {
	fodderAmount, errFodder := fc.FodderAmount(nCows)
	fatteningFactor, errFactor := fc.FatteningFactor()
	if errFodder != nil {
		return 0, errFodder
	}
	if errFactor != nil {
		return 0, errFactor
	}
	return fodderAmount * fatteningFactor / float64(nCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, nCows int) (float64, error) {
	if nCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, nCows)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	nCows    int
	errorMSG string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprint(err.nCows) + " cows are invalid: " + err.errorMSG
}

func ValidateNumberOfCows(nCows int) error {
	switch {
	case nCows < 0:
		return &InvalidCowsError{
			nCows:    nCows,
			errorMSG: "there are no negative cows",
		}
	case nCows == 0:
		return &InvalidCowsError{
			nCows:    0,
			errorMSG: "no cows don't need food",
		}
	default:
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
