package primitivetypes

import "testing"

func TestCountBits(t *testing.T) {
	for _, test := range testCases {
		actual, err := CountBits(test.input)
		if actual != test.expected {
			t.Fatalf("Fail: %s\nExpected: %v\nActual: %v", test.description, test.expected, actual)
		}
		if test.err != nil {
			if err == nil {
				t.Fatalf("Fail: %s\nExpected: %v\nActual: %v", test.description, test.expected, actual)
			}
		}
		t.Logf("Pass: %s", test.description)
	}

}

// for _, tc := range testCases {
//   if actual := Age(tc.seconds, tc.planet); actual != tc.expected {
//     t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
//   }
//   t.Logf("PASS: %s", tc.description)
// }
