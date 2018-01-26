package primitivetypes

import "testing"

func TestMultiply(t *testing.T) {
	for _, test := range mTestCases {
		actual := Multiply(test.input1, test.input2)
		if actual != test.expected {
			t.Fatalf("Fail: %s\nExpected: %v\nActual: %v", test.description, test.expected, actual)
		}
	}
}
