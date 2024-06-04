package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0105 struct {
	preorder []int
	inorder  []int
}

func TestConstructBinaryTreefromPreorderandInorderTraversal(t *testing.T) {
	tests := []struct {
		parameters0105
		// 填入 function output type
		ans *structures.TreeNode
	}{
		// 填入 test case
		{
			parameters0105{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
			structures.Ints2TreeNode([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}),
		},
		{
			parameters0105{[]int{-1}, []int{-1}},
			structures.Ints2TreeNode([]int{-1}),
		},
	}

	for _, test := range tests {
		t.Run("Test ConstructBinaryTreefromPreorderandInorderTraversal", func(t *testing.T) {
			// 完整輸入參數
			result := ConstructBinaryTreefromPreorderandInorderTraversal(test.parameters0105.preorder, test.parameters0105.inorder)
			// compare 的方式需視情況調整
			if !structures.TreeNodeIsEqual(result, test.ans) {
				t.Errorf("ConstructBinaryTreefromPreorderandInorderTraversal(%+v) got %+v, want %+v", test.parameters0105, result, test.ans)
			}
		})
	}
}
