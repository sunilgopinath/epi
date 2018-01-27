package primitivetypes

import (
	"math"
)

// Reverse returns the digits of an integer in reverse order
func Reverse(x int) int {
	a := Abs(x)
	var y int
	for a > 0 {
		y = y*10 + a%10
		a /= 10
	}
	if y < math.MinInt32 || y > math.MaxInt32 {
		return 0
	}
	if x < 0 {
		y *= -1
	}
	return y
}

// Abs returns the absolute value of an int32
func Abs(x int) int {
	if x < 0 {
		x *= -1
		return x
	}
	return x
}
