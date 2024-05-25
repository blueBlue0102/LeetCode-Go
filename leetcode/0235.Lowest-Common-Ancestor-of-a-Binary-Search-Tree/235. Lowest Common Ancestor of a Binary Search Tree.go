package leetcode

import "github.com/blueBlue0102/LeetCode-Go/structures"

func LowestCommonAncestorofaBinarySearchTree(root, p, q *structures.TreeNode) *structures.TreeNode {
	cursor := root
	for cursor != nil {
		if cursor.Val < p.Val && cursor.Val < q.Val {
			cursor = cursor.Right
		} else if cursor.Val > p.Val && cursor.Val > q.Val {
			cursor = cursor.Left
		} else {
			return cursor
		}
	}
	return nil
}
