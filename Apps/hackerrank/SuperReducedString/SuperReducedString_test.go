package SuperReducedString

import "testing"

var tests = []struct {
	input    string
	expected string
}{
	{"daabbccd", ""},
	{"aabb", ""},
	{"aaab", "ab"},
	{"aaabbbccaaa", "aba"},
	{"aaabccddd", "abd"},
	{"aa", ""},
	{"baab", ""},
}

func TestSuperReducedString(t *testing.T) {
	for _, test := range tests {
		actual := SuperReducedString(test.input)
		if actual != test.expected {
			t.Errorf("Error. Expected: %v , Got: %v , Input: %v", test.expected, actual, test.input)
		}
	}
}
