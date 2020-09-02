package maxVowels

import "testing"

var tests = []struct {
	input    string
	k        int
	expected int
}{

	{"abcdeeeefg", 4, 4},
	{"ab", 4, 1},
	{"abababab", 4, 4},
	{"aaaaaa", 4, 4},
	{"weallloveyou", 7, 4},
	{"qempburycnhrvvccr", 13, 2},
}

func TestmaxVowels(t *testing.T) {
	for _, test := range tests {
		actual := maxVowels(test.input, test.k)
		if actual != test.expected {
			t.Errorf("Error. Expected: %v , Got: %v , Input: %v", test.expected, actual, test.input)
		}
	}
}
