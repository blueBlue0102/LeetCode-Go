package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0100 struct {
	p *structures.TreeNode
	q *structures.TreeNode
}

func TestSameTree(t *testing.T) {
	tests := []struct {
		parameters0100
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters0100{structures.Ints2TreeNode([]int{1, 2, 3}), structures.Ints2TreeNode([]int{1, 2, 3})},
			true,
		},
		{
			parameters0100{structures.Ints2TreeNode([]int{1, 2}), structures.Ints2TreeNode([]int{1, structures.NULL, 2})},
			false,
		},
		{
			parameters0100{structures.Ints2TreeNode([]int{1, 2, 1}), structures.Ints2TreeNode([]int{1, 1, 2})},
			false,
		},
	}

	for _, test := range tests {
		t.Run("Test SameTree", func(t *testing.T) {
			// 完整輸入參數
			result := SameTree(test.parameters0100.p, test.parameters0100.q)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("SameTree(%+v) got %+v, want %+v", test.parameters0100, result, test.ans)
			}
		})
	}
}
