package leetcode

import "github.com/blueBlue0102/LeetCode-Go/structures"

func MaximumDepthofBinaryTree(root *structures.TreeNode) int {
	height := 0
	if root != nil {
		height++
		height += max(MaximumDepthofBinaryTree(root.Left), MaximumDepthofBinaryTree(root.Right))
	}
	return height
}
