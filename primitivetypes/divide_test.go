package primitivetypes

import (
	"testing"
)

func TestDivide(t *testing.T) {
	for _, test := range dTestCases {
		actual, err := Divide(test.input1, test.input2)
		if err != nil {
			if test.err == nil {
				t.Fatalf("Fail: %s\nExpected: %v\nActual: %v", test.description, test.expected, err)
			}
		}
		if actual != test.expected {
			t.Fatalf("Fail: %s\nExpected: %v\nActual: %v", test.description, test.expected, actual)
		}
	}
}
