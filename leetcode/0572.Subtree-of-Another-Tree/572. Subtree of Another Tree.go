package leetcode

import "github.com/blueBlue0102/LeetCode-Go/structures"

func SubtreeofAnotherTree(root *structures.TreeNode, subRoot *structures.TreeNode) bool {
	if IsSameTree(root, subRoot) {
		return true
	} else if root != nil && subRoot != nil && (SubtreeofAnotherTree(root.Left, subRoot) || SubtreeofAnotherTree(root.Right, subRoot)) {
		return true
	}
	return false
}

func IsSameTree(p *structures.TreeNode, q *structures.TreeNode) bool {
	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		if !IsSameTree(p.Left, q.Left) || !IsSameTree(p.Right, q.Right) {
			return false
		}
	}
	return true
}
