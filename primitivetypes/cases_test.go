package primitivetypes

import "errors"

var testCases = []struct {
	description string
	input       int
	expected    int
	err         error
}{
	{
		"basic test",
		5,
		2,
		nil,
	},
	{
		"test edge 0",
		0,
		0,
		nil,
	},
	{
		"test negative numbers",
		-1,
		-1,
		errors.New("cannot work with negative numbers"),
	},
}
