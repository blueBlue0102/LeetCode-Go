package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0226 struct {
	root *structures.TreeNode
}

func TestInvertBinaryTree(t *testing.T) {
	tests := []struct {
		parameters0226
		// 填入 function output type
		ans *structures.TreeNode
	}{
		// 填入 test case
		{
			parameters0226{structures.Ints2TreeNode([]int{4, 2, 7, 1, 3, 6, 9})},
			structures.Ints2TreeNode([]int{4, 7, 2, 9, 6, 3, 1}),
		},
		{
			parameters0226{structures.Ints2TreeNode([]int{2, 1, 3})},
			structures.Ints2TreeNode([]int{2, 3, 1}),
		},
		{
			parameters0226{structures.Ints2TreeNode([]int{})},
			structures.Ints2TreeNode([]int{}),
		},
	}

	for _, test := range tests {
		t.Run("Test InvertBinaryTree", func(t *testing.T) {
			// 完整輸入參數
			result := InvertBinaryTree(test.parameters0226.root)
			// compare 的方式需視情況調整
			if !structures.TreeNodeIsEqual(result, test.ans) {
				t.Errorf("InvertBinaryTree(%+v) got %+v, want %+v", test.parameters0226, result, test.ans)
			}
		})
	}
}
