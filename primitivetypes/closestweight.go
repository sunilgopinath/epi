package primitivetypes

import "errors"

// NumDigits of bits in the input
const NumDigits = 63

// ClosestIntSameBitCount returns the closest min(|y-x|) of the same bit count
func ClosestIntSameBitCount(x int64) (int64, error) {

	// the approach we are taking here is to swap the right most
	// digits which differ
	// ex: 1011 we would swap digit 1 and 2 (reading left to right, 0 indexed)
	var i uint
	for i = 0; i < NumDigits-1; i++ {
		if (x>>i)&1 != (x >> (i + 1) & 1) {
			// swap digits using bit mask
			x ^= (1 << i) | 1<<(i+1)
			return x, nil
		}
	}
	return 0, errors.New("input is all 0s or all 1s")
}
