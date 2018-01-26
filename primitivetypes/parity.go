package primitivetypes

import (
	"errors"
)

// Parity returns 1 or 0 depending on if the number of 1s is even or odd
func Parity(x int) (int, error) {
	if x < 0 {
		return -1, errors.New("cannot accept negative numbers")
	}
	var c int
	for x != 0 {
		c ^= (x & 1)
		x >>= 1
	}
	return c, nil
}
