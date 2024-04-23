package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0094 struct {
	root *structures.TreeNode
}

func TestBinaryTreeInorderTraversal(t *testing.T) {
	tests := []struct {
		parameters0094
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0094{structures.Ints2TreeNode([]int{1, structures.NULL, 2, 3})},
			[]int{1, 3, 2},
		},
		{
			parameters0094{structures.Ints2TreeNode([]int{})},
			[]int{},
		},
		{
			parameters0094{structures.Ints2TreeNode([]int{1})},
			[]int{1},
		},
		{
			parameters0094{structures.Ints2TreeNode([]int{1, 2, 3, structures.NULL, 4, structures.NULL, structures.NULL, 5, 6, structures.NULL, 7})},
			[]int{2, 5, 7, 4, 6, 1, 3},
		},
	}

	for _, test := range tests {
		t.Run("Test BinaryTreeInorderTraversal", func(t *testing.T) {
			// 完整輸入參數
			result := BinaryTreeInorderTraversal(test.parameters0094.root)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("BinaryTreeInorderTraversal(%+v) got %+v, want %+v", test.parameters0094, result, test.ans)
			}
		})
	}
}
