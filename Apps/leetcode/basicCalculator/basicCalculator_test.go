package basicCalculator

import "testing"

var tests = []struct {
	input    string
	expected int
}{
	{"10+20", 30},
	// {"100+200-300", 0},
	// {"2-1 + 2", 3},
	// {"(1+(4+5+2)-3)+(6+8)", 23},
	// {"1", 1},
	// {"1 - 2", -1},
	// {},
}

func TestBasicCalculator(t *testing.T) {
	for _, test := range tests {
		actual := BasicCalculator(test.input)
		if actual != test.expected {
			t.Errorf("Error. Expected: %v , Got: %v , Input: %v", test.expected, actual, test.input)
		}
	}
}
