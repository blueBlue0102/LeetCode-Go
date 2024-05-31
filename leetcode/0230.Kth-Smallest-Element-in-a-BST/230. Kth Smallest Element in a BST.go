package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func KthSmallestElementinaBST(root *structures.TreeNode, k int) int {
	// return Iterative(root,k)
	return Recursion(root, k)
}

func Iterative(root *structures.TreeNode, k int) int {
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

func InOrderDFS(root *structures.TreeNode, k int, result *[]int) {
	if len(*result) == k {
		return
	}
	if root != nil {
		InOrderDFS(root.Left, k, result)
		(*result) = append((*result), root.Val)
		InOrderDFS(root.Right, k, result)
	}
}

func Recursion(root *structures.TreeNode, k int) int {
	nodeValArray := make([]int, 0, k+1)
	InOrderDFS(root, k, &nodeValArray)
	return nodeValArray[k-1]
}
