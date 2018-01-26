package primitivetypes

var mTestCases = []struct {
	description string
	input1      int64
	input2      int64
	expected    int64
}{
	{
		"basic test",
		5,
		3,
		15,
	},
	{
		"multiply by zero",
		0,
		3,
		0,
	},
	{
		"large numbers",
		10000000,
		3,
		30000000,
	},
}
