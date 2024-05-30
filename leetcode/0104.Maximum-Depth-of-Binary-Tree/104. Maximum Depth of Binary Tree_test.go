package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0104 struct {
	root *structures.TreeNode
}

func TestMaximumDepthofBinaryTree(t *testing.T) {
	tests := []struct {
		parameters0104
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0104{structures.Ints2TreeNode([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7})},
			3,
		},
		{
			parameters0104{structures.Ints2TreeNode([]int{1, structures.NULL, 2})},
			2,
		},
	}

	for _, test := range tests {
		t.Run("Test MaximumDepthofBinaryTree", func(t *testing.T) {
			// 完整輸入參數
			result := MaximumDepthofBinaryTree(test.parameters0104.root)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("MaximumDepthofBinaryTree(%+v) got %+v, want %+v", test.parameters0104, result, test.ans)
			}
		})
	}
}
