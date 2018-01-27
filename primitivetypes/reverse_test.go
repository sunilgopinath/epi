package primitivetypes

import "testing"

func TestReverse(t *testing.T) {
	for _, test := range rTestCases {
		actual := Reverse(test.input)
		if actual != test.expected {
			t.Fatalf("Fail: %s\nExpected: %v\nActual: %v", test.description, test.expected, actual)
		}
	}
}
