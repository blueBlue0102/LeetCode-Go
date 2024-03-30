package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0092 struct {
	head  *structures.ListNode
	left  int
	right int
}

func TestReverseLinkedListII(t *testing.T) {
	tests := []struct {
		parameters0092
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0092{structures.Ints2List([]int{1, 2, 3, 4, 5}), 2, 4},
			structures.Ints2List([]int{1, 4, 3, 2, 5}),
		},
		{
			parameters0092{structures.Ints2List([]int{5}), 1, 1},
			structures.Ints2List([]int{5}),
		},
		{
			parameters0092{structures.Ints2List([]int{1, 2, 3, 4, 5}), 1, 5},
			structures.Ints2List([]int{5, 4, 3, 2, 1}),
		},
		{
			parameters0092{structures.Ints2List([]int{1, 2, 3, 4, 5}), 3, 5},
			structures.Ints2List([]int{1, 2, 5, 4, 3}),
		},
		{
			parameters0092{structures.Ints2List([]int{1, 2, 3, 4, 5}), 1, 1},
			structures.Ints2List([]int{1, 2, 3, 4, 5}),
		},
		{
			parameters0092{structures.Ints2List([]int{1, 2, 3, 4, 5}), 3, 3},
			structures.Ints2List([]int{1, 2, 3, 4, 5}),
		},
		{
			parameters0092{structures.Ints2List([]int{1, 2, 3, 4, 5}), 5, 5},
			structures.Ints2List([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		t.Run("Test ReverseLinkedListII", func(t *testing.T) {
			// 完整輸入參數
			result := ReverseLinkedListII(test.parameters0092.head, test.parameters0092.left, test.parameters0092.right)
			// compare 的方式需視情況調整
			if !structures.ListNodeIsEqual(result, test.ans) {
				t.Errorf("ReverseLinkedListII(%+v) got %+v, want %+v", test.parameters0092, result, test.ans)
			}
		})
	}
}
