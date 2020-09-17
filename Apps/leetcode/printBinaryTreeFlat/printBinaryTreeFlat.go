package printBinaryTreeFlat

import (
	"fmt"
	"strconv"
)

func PrintBinaryTreeFlat(root *Node) {
	var children []*Node
	childrenIndex := 0
	children = append(children, root)
	for len(children) != childrenIndex {
		current := children[childrenIndex]

		if current.left != nil {
			children = append(children, current.left)
		}

		if current.right != nil {
			children = append(children, current.right)
		}

		childrenIndex++
	}

	binIndex := 0
	numToProcess := 1
	for binIndex < len(children) {
		for i := 0; i < numToProcess; i++ {
			if binIndex+i < len(children) {
				fmt.Printf(strconv.Itoa(children[binIndex+i].value))
			} else {
				break
			}
		}
		fmt.Printf("\n")
		binIndex = binIndex + numToProcess
		numToProcess = 2 * numToProcess
	}
}

type Node struct {
	value int
	right *Node
	left  *Node
}
