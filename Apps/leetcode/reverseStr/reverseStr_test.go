package reverseStr

import "testing"

var tests = []struct {
	s        string
	k        int
	expected string
}{
	{"abcdefg", 2, "bacdfeg"},
	{"abcdefg", 8, "gfedcba"},
}

func TestReverseStr(t *testing.T) {
	for _, test := range tests {
		actual := ReverseStr(test.s, test.k)
		if actual != test.expected {
			t.Errorf("Error. Expected: %v , Got: %v , Input: %v : %v", test.expected, actual, test.s, test.k)
		}
	}
}
