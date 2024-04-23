package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0145 struct {
	root *structures.TreeNode
}

func TestBinaryTreePostorderTraversal(t *testing.T) {
	tests := []struct {
		parameters0145
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0145{structures.Ints2TreeNode([]int{1, structures.NULL, 2, 3})},
			[]int{3, 2, 1},
		},
		{
			parameters0145{structures.Ints2TreeNode([]int{})},
			[]int{},
		},
		{
			parameters0145{structures.Ints2TreeNode([]int{1})},
			[]int{1},
		},
	}

	for _, test := range tests {
		t.Run("Test BinaryTreePostorderTraversal", func(t *testing.T) {
			// 完整輸入參數
			result := BinaryTreePostorderTraversal(test.parameters0145.root)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("BinaryTreePostorderTraversal(%+v) got %+v, want %+v", test.parameters0145, result, test.ans)
			}
		})
	}
}
