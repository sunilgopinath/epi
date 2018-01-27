package primitivetypes

import "errors"

var dTestCases = []struct {
	description string
	input1      int
	input2      int
	expected    int
	err         error
}{
	{
		"basic test",
		12,
		2,
		6,
		nil,
	},
	{
		"divide by zero",
		12,
		0,
		-1,
		errors.New("cannot divide by 0"),
	},
	{
		"zero numerator",
		0,
		5,
		0,
		nil,
	},
}
