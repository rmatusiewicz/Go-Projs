package main

func main() {}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func DeepestLeavesSum(root *TreeNode) int {
	leftSum, leftDepth := 0, 0
	if root.Left != nil {
		leftSum, leftDepth = dig(0, root.Left)
	}

	rightSum, rightDepth := 0, 0
	if root.Right != nil {
		rightSum, rightDepth = dig(0, root.Right)
	}

	if leftDepth > rightDepth {
		return leftSum
	} else if leftDepth == rightDepth {
		return leftSum + rightSum
	} else {
		return rightSum
	}
}

func dig(depth int, n *TreeNode) (int, int) {
	var rightMaxDepth, rightMaxVal, leftMaxDepth, leftMaxVal int
	if n.Left != nil {
		leftMaxDepth, leftMaxVal = dig(depth+1, n.Left)
	}

	if n.Right != nil {
		rightMaxDepth, rightMaxVal = dig(depth+1, n.Right)
	}

	d := 0
	v := 0
	if rightMaxDepth > leftMaxDepth {
		d = rightMaxDepth
		v = rightMaxVal
	} else if rightMaxDepth == leftMaxDepth {
		d = rightMaxDepth
		v = rightMaxVal + leftMaxVal
	} else {
		d = leftMaxDepth
		v = leftMaxVal
	}

	return d, v
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}
