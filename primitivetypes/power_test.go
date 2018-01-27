package primitivetypes

import "testing"

func TestPower(t *testing.T) {
	for _, test := range poTestCases {
		actual := Power(test.input1, test.input2)
		if actual != test.expected {
			t.Fatalf("Fail: %s\nExpected: %v\nActual: %v", test.description, test.expected, actual)
		}
	}
}
