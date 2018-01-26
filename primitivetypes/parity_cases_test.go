package primitivetypes

import "errors"

var pTestCases = []struct {
	description string
	input       int
	expected    int
	err         error
}{
	{
		"basic 1 test ",
		5,
		0,
		nil,
	},
	{
		"basic 0 test",
		4,
		1,
		nil,
	},
	{
		"large number test",
		10000000,
		0,
		nil,
	},
	{
		"negative test",
		-1,
		-1,
		errors.New("cannot accept negative numbers"),
	},
}
