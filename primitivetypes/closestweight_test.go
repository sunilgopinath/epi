package primitivetypes

import "testing"

func TestClosestIntSameBitCount(t *testing.T) {
	for _, test := range cTestsCases {
		actual, err := ClosestIntSameBitCount(test.input)
		if err != nil {
			if test.err == nil {
				t.Fatalf("Fail: %s\nExpected: %v", test.description, test.err)
			}
		}
		if actual != test.expected {
			t.Fatalf("Fail: %s\nExpected: %v\nActual: %v", test.description, test.expected, actual)
		}
	}
}
