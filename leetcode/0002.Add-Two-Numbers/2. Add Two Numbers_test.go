package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0002 struct {
	l1 *structures.ListNode
	l2 *structures.ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		parameters0002
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0002{structures.Ints2List([]int{2, 4, 3}), structures.Ints2List([]int{5, 6, 4})},
			structures.Ints2List([]int{7, 0, 8}),
		},
		{
			parameters0002{structures.Ints2List([]int{0}), structures.Ints2List([]int{0})},
			structures.Ints2List([]int{0}),
		},
		{
			parameters0002{structures.Ints2List([]int{9, 9, 9, 9, 9, 9, 9}), structures.Ints2List([]int{9, 9, 9, 9})},
			structures.Ints2List([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for _, test := range tests {
		t.Run("Test AddTwoNumbers", func(t *testing.T) {
			// 完整輸入參數
			result := AddTwoNumbers(test.parameters0002.l1, test.parameters0002.l2)
			// compare 的方式需視情況調整
			if !structures.ListNodeIsEqual(result, test.ans) {
				t.Errorf("AddTwoNumbers(%+v) got %+v, want %+v", test.parameters0002, result, test.ans)
			}
		})
	}
}
