package validPalindrome

import "testing"

var tests = []struct {
	input    string
	expected bool
}{
	{"aba", true},
	{"aaa", true},
	{"a1A", true},
	{"a", true},
	{"abb", false},
	{"123321", true},
	{"", true},
	{"abaaba", true},
	{"abacaba", true},
	{"A man, a plan, a canal: Panama", true},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range tests {
		result := IsPalindrome(test.input)
		if test.expected != result {
			t.Errorf("Error: %t : %t", test.expected, result)
		}
	}
}
