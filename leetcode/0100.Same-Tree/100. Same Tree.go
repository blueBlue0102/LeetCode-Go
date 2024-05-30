package leetcode

import "github.com/blueBlue0102/LeetCode-Go/structures"

func SameTree(p *structures.TreeNode, q *structures.TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		if !SameTree(p.Left, q.Left) || !SameTree(p.Right, q.Right) {
			return false
		}
	}
	return true
}
