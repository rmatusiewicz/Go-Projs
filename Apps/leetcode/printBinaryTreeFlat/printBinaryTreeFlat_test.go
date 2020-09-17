package printBinaryTreeFlat

import "testing"

var tests = []struct {
	input *Node
}{

	{&Node{
		value: 1,
		left: &Node{
			value: 2,
			left: &Node{
				value: 4,
				left:  nil,
				right: nil,
			},
			right: &Node{
				value: 5,
				left:  nil,
				right: nil,
			},
		},
		right: &Node{
			value: 3,
			left: &Node{
				value: 6,
				left:  nil,
				right: nil,
			},
			right: &Node{
				value: 7,
				left:  nil,
				right: nil,
			},
		},
	}},
}

func TestPrintBinaryTreeFlat(t *testing.T) {
	for _, test := range tests {
		PrintBinaryTreeFlat(test.input)
	}
}
