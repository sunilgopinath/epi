package primitivetypes

var poTestCases = []struct {
	description string
	input1      float32
	input2      int
	expected    float32
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
	{
		"raised to negative test",
		2,
		-1,
		0.5,
	},
	{
		"longer test",
		10,
		10,
		10000000000,
	},
}
