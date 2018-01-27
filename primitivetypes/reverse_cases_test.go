package primitivetypes

import "math"

var rTestCases = []struct {
	description string
	input       int
	expected    int
}{
	{
		"basic test",
		123,
		321,
	},
	{
		"negative test",
		-123,
		-321,
	},
	{
		"overflow test",
		math.MinInt32,
		0,
	},
	{
		"overflow test max",
		math.MaxInt32,
		0,
	},
}
