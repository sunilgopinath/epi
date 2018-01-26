package primitivetypes

import "testing"

func TestParity(t *testing.T) {
	for _, test := range pTestCases {
		actual, err := Parity(test.input)
		if actual != test.expected {
			t.Fatalf("Fail: %s\nExpected: %v\nActual: %v", test.description, test.expected, actual)
		}
		if err != nil {
			if test.err == nil {
				t.Fatalf("Fail: %s\nExpected: %v\nActual: %v", test.description, test.expected, err)
			}
		}
	}
}
