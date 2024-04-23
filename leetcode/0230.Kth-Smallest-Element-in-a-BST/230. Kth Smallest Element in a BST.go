package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func KthSmallestElementinaBST(root *structures.TreeNode, k int) int {
	nodeStack := []*structures.TreeNode{}
	pointer := root

	for pointer != nil || len(nodeStack) > 0 {
		for pointer != nil {
			nodeStack = append(nodeStack, pointer)
			pointer = pointer.Left
		}
		pointer = nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]

		k--
		if k == 0 {
			return pointer.Val
		}

		pointer = pointer.Right
	}

	return root.Val
}
