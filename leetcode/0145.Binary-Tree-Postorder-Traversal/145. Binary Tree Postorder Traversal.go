package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func BinaryTreePostorderTraversal(root *structures.TreeNode) []int {
	result := []int{}

	if root != nil {
		result = append(result, BinaryTreePostorderTraversal(root.Left)...)
		result = append(result, BinaryTreePostorderTraversal(root.Right)...)
		result = append(result, root.Val)
	}

	return result
}
