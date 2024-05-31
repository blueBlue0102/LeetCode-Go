package leetcode

import "github.com/blueBlue0102/LeetCode-Go/structures"

func BinaryTreeLevelOrderTraversal(root *structures.TreeNode) [][]int {
	result := [][]int{}
	cal(root, &result, 1)
	return result
}

func cal(root *structures.TreeNode, result *[][]int, level int) {
	if root != nil {
		if len(*result) < level {
			*result = append(*result, []int{})
		}
		(*result)[level-1] = append((*result)[level-1], root.Val)
		cal(root.Left, result, level+1)
		cal(root.Right, result, level+1)
	}
}
