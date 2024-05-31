package leetcode

import (
	"math"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func ValidateBinarySearchTree(root *structures.TreeNode) bool {
	// 題目 node 數量 >= 1，所以不考慮 root == nil
	return IsBST(root, math.MinInt, math.MaxInt)
}

func IsBST(root *structures.TreeNode, leftBound int, rightBound int) bool {
	if root != nil {
		if root.Left != nil && (root.Left.Val >= root.Val || root.Left.Val <= leftBound) {
			return false
		}
		if root.Right != nil && (root.Right.Val <= root.Val || root.Right.Val >= rightBound) {
			return false
		}
		if !IsBST(root.Left, leftBound, root.Val) || !IsBST(root.Right, root.Val, rightBound) {
			return false
		}
	}
	return true
}
