package main

import "testing"

var countTests = []struct {
	input    string
	expected int
}{
	{"Hello", 1},
	{"Hello there", 2},
	{"Hello there, how are you?", 5},
	{"", 0},
	{" hello ", 1},
	{"    ", 0},
	{"   , , , ,        a, eaefa", 6},
}

func TestCountSegments(t *testing.T) {
	for _, tt := range countTests {
		actual := CountSegs(tt.input)
		if tt.expected != actual {
			t.Errorf("Input: %s, Expected: %v, Actual: %v", tt.input, tt.expected, actual)
		}
	}
}
