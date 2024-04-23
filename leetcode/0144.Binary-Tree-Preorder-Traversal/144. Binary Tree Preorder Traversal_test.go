package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0144 struct {
	root *structures.TreeNode
}

func TestBinaryTreePreorderTraversal(t *testing.T) {
	tests := []struct {
		parameters0144
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0144{structures.Ints2TreeNode([]int{1, structures.NULL, 2, 3})},
			[]int{1, 2, 3},
		},
		{
			parameters0144{structures.Ints2TreeNode([]int{})},
			[]int{},
		},
		{
			parameters0144{structures.Ints2TreeNode([]int{1})},
			[]int{1},
		},
		{
			parameters0144{structures.Ints2TreeNode([]int{1, 2, 3, structures.NULL, 4, structures.NULL, structures.NULL, 5, 6, structures.NULL, 7})},
			[]int{1, 2, 4, 5, 7, 6, 3},
		},
	}

	for _, test := range tests {
		t.Run("Test BinaryTreePreorderTraversal", func(t *testing.T) {
			// 完整輸入參數
			result := BinaryTreePreorderTraversal(test.parameters0144.root)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("BinaryTreePreorderTraversal(%+v) got %+v, want %+v", test.parameters0144, result, test.ans)
			}
		})
	}
}
