package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func BinaryTreePreorderTraversal(root *structures.TreeNode) []int {
	result := []int{}

	if root != nil {
		result = append(result, root.Val)
		result = append(result, BinaryTreePreorderTraversal(root.Left)...)
		result = append(result, BinaryTreePreorderTraversal(root.Right)...)
	}

	return result
}
