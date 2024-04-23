package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func BinaryTreeInorderTraversal(root *structures.TreeNode) []int {
	// return recur(root)
	return stack(root)
}

func recur(root *structures.TreeNode) []int {
	result := []int{}

	if root != nil {
		result = append(result, recur(root.Left)...)
		result = append(result, root.Val)
		result = append(result, recur(root.Right)...)
	}

	return result
}

func stack(root *structures.TreeNode) []int {
	result := []int{}
	nodeStack := []*structures.TreeNode{}
	pointer := root
	for pointer != nil || len(nodeStack) > 0 {
		// 移動 pointer 去最左邊
		for pointer != nil {
			// 沿路看到的 node 都放入 stack
			nodeStack = append(nodeStack, pointer)
			pointer = pointer.Left
		}
		// pointer 已變成 null，進行一次 backtrack
		pointer = nodeStack[len(nodeStack)-1]

		// visit node
		result = append(result, pointer.Val)

		// 往右子樹移動
		pointer = pointer.Right

		// 該 node 已遍歷完成，pop stack
		nodeStack = nodeStack[:len(nodeStack)-1]
	}
	return result
}
