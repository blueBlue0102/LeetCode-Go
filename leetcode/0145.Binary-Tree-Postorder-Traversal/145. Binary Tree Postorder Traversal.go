package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func BinaryTreePostorderTraversal(root *structures.TreeNode) []int {
	// return recur(root)
	return stack(root)
}

func recur(root *structures.TreeNode) []int {
	result := []int{}

	if root != nil {
		result = append(result, recur(root.Left)...)
		result = append(result, recur(root.Right)...)
		result = append(result, root.Val)
	}

	return result
}

func stack(root *structures.TreeNode) []int {
	result := []int{}
	pointer := root
	var lastVisited *structures.TreeNode = nil
	nodeStack := []*structures.TreeNode{}

	for pointer != nil && lastVisited != root {
		// 移動至最左 node
		for pointer != nil && pointer != lastVisited {
			// 將沿路的 node 放入 stack
			nodeStack = append(nodeStack, pointer)
			pointer = pointer.Left
		}
		// backtrack
		pointer = nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]

		// 如果右子樹存在，繼續往右子樹探尋
		if pointer.Right == nil || pointer.Right == lastVisited {
			// visit node
			result = append(result, pointer.Val)
			// mark lastVisited
			lastVisited = pointer
		} else {
			// traverse the node's right subtree
			nodeStack = append(nodeStack, pointer)
			pointer = pointer.Right
		}
	}

	return result
}
