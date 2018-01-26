package primitivetypes

import "errors"

//CountBits counts the number of 1s in a binary representation of an int
func CountBits(x int) (int, error) {
	if x < 0 {
		return -1, errors.New("can't work with negative numbers")
	}

	// what we're doing here is figuring out if the least significant bit is 1
	// then shifiting right to remove it
	var n int
	for x != 0 {
		n += (x & 1)
		x >>= 1
	}
	return n, nil
}
