package primitivetypes

import (
	"errors"
)

// Divide returns the quotient without using arithmentic divide
func Divide(x, y int) (int, error) {
	if x == 0 {
		return 0, nil
	}
	if y == 0 {
		return -1, errors.New("Cannot divide by 0")
	}
	var r int
	for x >= y {
		r++
		x -= y
	}
	return r, nil
}
