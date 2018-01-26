package primitivetypes

// Multiply does not use arithmetic ops to multiply
func Multiply(x, y int64) int64 {
	if y == 0 {
		return 0
	}
	return add(x, Multiply(x, y-1))
}

func add(x, y int64) int64 {
	/* for ex:
	    1001 +
	    1101
	   -------
	   10110

	   Explanation:
	   1. To get 0 in LSB we do XOR as two 1s will give us 0
	   2. If that is the case we need to add the carry and shift (ripple)
	   3. Sum is carried through x

	*/

	for y != 0 {
		carry := x & y
		x ^= y
		y = carry << 1
	}
	return x
}
