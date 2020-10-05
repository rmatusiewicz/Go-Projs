package MinimumBribes

import "testing"

var tests = []struct {
	input    []int32
	expected int
}{
	{[]int32{2, 1, 5, 3, 4}, 3},
	//1, 2, 3, 4, 5

	{[]int32{2, 1, 5, 3, 4}, 3},
	{[]int32{2, 5, 1, 3, 4}, -1},
}

func TestMinimumBribes(t *testing.T) {
	for _, test := range tests {
		actual := MinimumBribes(test.input)
		if actual != test.expected {
			t.Errorf("Error. Expected: %v , Got: %v , Input: %v", test.expected, actual, test.input)
		}
	}
}
