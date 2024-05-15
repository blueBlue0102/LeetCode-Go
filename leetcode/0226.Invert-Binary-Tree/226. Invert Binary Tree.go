package leetcode

import "github.com/blueBlue0102/LeetCode-Go/structures"

func InvertBinaryTree(root *structures.TreeNode) *structures.TreeNode {
	if root != nil {
		InvertBinaryTree(root.Left)
		InvertBinaryTree(root.Right)
		root.Left, root.Right = root.Right, root.Left
	}
	return root
}
