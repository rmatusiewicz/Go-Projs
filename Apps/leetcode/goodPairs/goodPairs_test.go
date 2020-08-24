package main

import "testing"

var tests = []struct {
	nums     []int
	expected int
}{
	{[]int{1, 2, 3, 1, 1, 3}, 4},
	{[]int{1, 1, 1, 1}, 6},
	{[]int{1, 2, 3}, 0},
	{[]int{1, 2, 1, 2, 1, 2}, 6},
}

func TestNumIdenticalPairs(t *testing.T) {
	for _, test := range tests {
		result := NumIdenticalPairs(test.nums)
		if test.expected != result {
			t.Errorf("Failure %d : %d", test.expected, result)
		}
	}
}
