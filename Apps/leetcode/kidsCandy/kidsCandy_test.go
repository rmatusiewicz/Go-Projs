package main

import "testing"

var tests = []struct {
	candies  []int
	extra    int
	expected []bool
}{
	{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
	{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
	{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	{[]int{1, 1, 1}, 3, []bool{true, true, true}},
}

func TestKidsWithCandies(t *testing.T) {
	for _, test := range tests {
		x := kidsWithCandies(test.candies, test.extra)
		if !checkResults(x, test.expected) {
			t.Errorf("Error: %v: %v: %v ", test.candies, test.extra, x)
		}
	}
}

func checkResults(x []bool, y []bool) bool {
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
