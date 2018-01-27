package primitivetypes

// Power returns the first number exponential wrt to the second
func Power(x float32, y int) float32 {
	var r float32
	r = 1.0
	p := y
	if y < 0 {
		p *= -1
		x = 1.0 / x
	}
	/*
	  What we need to do is instead of just doing multiplication as many
	  times as the y value, ie x*x*x*x... y times, we can use a property of
	  when y is even to square the value

	*/
	for p != 0 {
		if p&1 != 0 { // or p%2 (test or even)
			r *= x
		}
		x *= x
		p >>= 1 // basically shifts down p
	}
	return r
}
