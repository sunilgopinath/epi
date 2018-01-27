package primitivetypes

// Power returns the first number exponential wrt to the second
func Power(x, y int) int {
	if y == 0 {
		return 1
	}
	c := 1
	for c < y {
		x *= x
		c++
	}
	return x
}
