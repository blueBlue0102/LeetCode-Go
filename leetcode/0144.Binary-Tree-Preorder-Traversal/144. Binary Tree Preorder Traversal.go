package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func BinaryTreePreorderTraversal(root *structures.TreeNode) []int {
	// return recur(root)
	return stack(root)
}

func recur(root *structures.TreeNode) []int {
	result := []int{}

	if root != nil {
		result = append(result, root.Val)
		result = append(result, recur(root.Left)...)
		result = append(result, recur(root.Right)...)
	}

	return result
}

func stack(root *structures.TreeNode) []int {
	result := []int{}
	pointer := root
	nodeStack := []*structures.TreeNode{}

	for pointer != nil || len(nodeStack) > 0 {
		// 移動至最左邊 node
		for pointer != nil {
			// visit node
			result = append(result, pointer.Val)

			// 將沿路看到的 node 放入 stack
			nodeStack = append(nodeStack, pointer)
			pointer = pointer.Left
		}
		// backtrack
		// 往右子樹移動
		pointer = nodeStack[len(nodeStack)-1].Right

		// pop stack
		nodeStack = nodeStack[:len(nodeStack)-1]
	}
	return result
}
