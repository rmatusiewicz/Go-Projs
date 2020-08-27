package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	if root.Left != nil {
		invertTree(root.Left)
	}

	if root.Right != nil {
		invertTree(root.Right)
	}

	return root
}
