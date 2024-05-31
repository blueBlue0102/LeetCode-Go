package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0098 struct {
	root *structures.TreeNode
}

func TestValidateBinarySearchTree(t *testing.T) {
	tests := []struct {
		parameters0098
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters0098{structures.Ints2TreeNode([]int{2, 1, 3})},
			true,
		},
		{
			parameters0098{structures.Ints2TreeNode([]int{5, 1, 4, structures.NULL, structures.NULL, 3, 6})},
			false,
		},
		{
			parameters0098{structures.Ints2TreeNode([]int{2, 2, 2})},
			false,
		},
		{
			parameters0098{structures.Ints2TreeNode([]int{5, 4, 6, structures.NULL, structures.NULL, 3, 7})},
			false,
		},
	}

	for _, test := range tests {
		t.Run("Test ValidateBinarySearchTree", func(t *testing.T) {
			// 完整輸入參數
			result := ValidateBinarySearchTree(test.parameters0098.root)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("ValidateBinarySearchTree(%+v) got %+v, want %+v", test.parameters0098, result, test.ans)
			}
		})
	}
}
