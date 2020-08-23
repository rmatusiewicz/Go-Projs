package main

import "testing"

var tests = []struct {
	nums     []int
	n        int
	expected []int
}{
	{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
	{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
	{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	{[]int{1, 1, 1, 1}, 2, []int{1, 1, 1, 1}},
}

func TestShuffle(t *testing.T) {
	for _, tt := range tests {
		deck := Shuffle(tt.nums, tt.n)
		if len(tt.expected) != len(deck) {
			t.Errorf("Lengths wrong: %v: %v:", tt.nums, deck)
		}
		for i := range deck {
			if deck[i] != tt.expected[i] {
				t.Errorf("Wrong match: %v: %v:", tt.nums[i], deck[i])
			}
		}
	}
}
