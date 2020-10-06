package MatchingStrings

import "testing"

var tests = []struct {
	strs     []string
	queries  []string
	expected []int32
}{
	{[]string{"ab", "ab", "abc"}, []string{"ab", "abc", "bc"}, []int32{2, 1, 0}},
	{[]string{"aba", "baba", "aba", "xzxb"}, []string{"aba", "xzxb", "ab"}, []int32{2, 1, 0}},
	{[]string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"}, []string{"abcde", "sdaklfj", "asdjf", "na", "basdn"}, []int32{1, 3, 4, 3, 2}},
}

func TestMatchingStrings(t *testing.T) {
	for _, test := range tests {
		actual := MatchingStrings(test.strs, test.queries)
		for i, x := range actual {
			if x != test.expected[i] {
				t.Errorf("Error. Expected: %v , Got: %v , Input: %v", test.expected, actual, test.strs)
			}
		}
	}
}
