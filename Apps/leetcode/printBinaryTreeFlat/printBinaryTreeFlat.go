package printBinaryTreeFlat

func main() {
	var testTree = &Node{
		{value, 1},
		{left, &Node{
			{value, 2},
			{left, nil},
			{right, nil},
		}},
		{right, nil},
	}
}

func printBinaryTreeFlat(root *Node) {
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
}

type Node struct {
	value int
	right *Node
	left  *Node
}
