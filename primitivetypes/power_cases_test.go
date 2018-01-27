package primitivetypes

var poTestCases = []struct {
	description string
	input1      int
	input2      int
	expected    int
}{
	{
		"basic test",
		2,
		2,
		4,
	},
	{
		"zero test",
		0,
		2,
		0,
	},
	{
		"raised to zero test",
		2,
		0,
		1,
	},
}
