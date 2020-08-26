package main

func ConvertArrayToBinarySearchTree(input []int) *TreeNode {
	root := &TreeNode{Val: input[0]}
	assignmentCounter := 0
	iterator := 1
	current := root
	var assignChildren []*TreeNode

	for iterator <= len(input) {
		//assignLeft
		current.Left = &TreeNode{Val: input[iterator]}
		iterator = iterator + 1
		assignChildren = append(assignChildren, current.Left)
		if iterator > len(input) {
			break
		}

		//assignRight
		current.Right = &TreeNode{Val: input[iterator]}
		iterator = iterator + 1
		assignChildren = append(assignChildren, current.Right)
		if iterator > len(input) {
			break
		}
		current = assignChildren[assignmentCounter]
		assignmentCounter++
	}
	return root
}

// func addChild(node *TreeNode, val int) {
// 	node.Left = &TreeNode{Val: val}
// 	childrenCounter = childrenCounter + 1
// 	if childrenCounter >= len(input) {
// 		break
// 	}
// }
