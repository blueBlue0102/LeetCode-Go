package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0230 struct {
	root *structures.TreeNode
	k    int
}

func TestKthSmallestElementinaBST(t *testing.T) {
	tests := []struct {
		parameters0230
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0230{structures.Ints2TreeNode([]int{3, 1, 4, structures.NULL, 2}), 1},
			1,
		},
		{
			parameters0230{structures.Ints2TreeNode([]int{5, 3, 6, 2, 4, structures.NULL, structures.NULL, 1}), 3},
			3,
		},
	}

	for _, test := range tests {
		t.Run("Test KthSmallestElementinaBST", func(t *testing.T) {
			// 完整輸入參數
			result := KthSmallestElementinaBST(test.parameters0230.root, test.parameters0230.k)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("KthSmallestElementinaBST(%+v) got %+v, want %+v", test.parameters0230, result, test.ans)
			}
		})
	}
}
