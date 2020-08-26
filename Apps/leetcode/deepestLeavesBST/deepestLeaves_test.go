package main

import "testing"

var tests = []struct {
	input    []int
	expected int
}{
	{[]int{1, 2, 3, 4, 5, 0, 6, 7, 0, 0, 0, 0, 8}, 15},
}

func TestDeepestLeavesSum(t *testing.T) {
	for _, tt := range tests {
		inp := ConvertArrayToBinarySearchTree(tt.input)
		result := DeepestLeavesSum(inp)
		if result != tt.expected {

		}
	}
	//[1,2,3,4,5,null,6,7,null,null,null,null,8]
	//             1
	// 	   2            3
	// 	 4     5        nil      6
	//7 nil  nil nil  nil  nil   nil 8

}
