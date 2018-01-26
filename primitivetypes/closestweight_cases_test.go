package primitivetypes

import "errors"

var cTestsCases = []struct {
	description string
	input       int64
	expected    int64
	err         error
}{
	{
		"basic test",
		6,
		5,
		nil,
	},
	{
		"all zeros",
		0,
		0,
		errors.New("input is all 0s or all 1s"),
	},
	{
		"very large",
		9223372036854775806,
		9223372036854775805,
		errors.New("input is all 0s or all 1s"),
	},
}
