package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0102 struct {
	root *structures.TreeNode
}

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {
	tests := []struct {
		parameters0102
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters0102{structures.Ints2TreeNode([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7})},
			[][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			parameters0102{structures.Ints2TreeNode([]int{1})},
			[][]int{{1}},
		},
		{
			parameters0102{structures.Ints2TreeNode([]int{})},
			[][]int{},
		},
	}

	for _, test := range tests {
		t.Run("Test BinaryTreeLevelOrderTraversal", func(t *testing.T) {
			// 完整輸入參數
			result := BinaryTreeLevelOrderTraversal(test.parameters0102.root)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("BinaryTreeLevelOrderTraversal(%+v) got %+v, want %+v", test.parameters0102, result, test.ans)
			}
		})
	}
}
