package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0572 struct {
	root    *structures.TreeNode
	subRoot *structures.TreeNode
}

func TestSubtreeofAnotherTree(t *testing.T) {
	tests := []struct {
		parameters0572
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters0572{structures.Ints2TreeNode([]int{3, 4, 5, 1, 2}), structures.Ints2TreeNode([]int{4, 1, 2})},
			true,
		},
		{
			parameters0572{structures.Ints2TreeNode([]int{3, 4, 5, 1, 2, structures.NULL, structures.NULL, structures.NULL, structures.NULL, 0}), structures.Ints2TreeNode([]int{4, 1, 2})},
			false,
		},
	}

	for _, test := range tests {
		t.Run("Test SubtreeofAnotherTree", func(t *testing.T) {
			// 完整輸入參數
			result := SubtreeofAnotherTree(test.parameters0572.root, test.parameters0572.subRoot)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("SubtreeofAnotherTree(%+v) got %+v, want %+v", test.parameters0572, result, test.ans)
			}
		})
	}
}
