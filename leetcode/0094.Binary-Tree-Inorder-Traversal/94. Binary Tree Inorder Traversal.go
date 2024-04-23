package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func BinaryTreeInorderTraversal(root *structures.TreeNode) []int {
	result := []int{}

	if root != nil {
		result = append(result, BinaryTreeInorderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, BinaryTreeInorderTraversal(root.Right)...)
	}

	return result
}
