package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0203 struct {
	head *structures.ListNode
	val  int
}

func TestRemoveLinkedListElements(t *testing.T) {
	tests := []struct {
		parameters0203
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0203{head: structures.Ints2List([]int{1, 2, 6, 3, 4, 5, 6}), val: 6},
			structures.Ints2List([]int{1, 2, 3, 4, 5}),
		},
		{
			parameters0203{head: structures.Ints2List([]int{}), val: 1},
			structures.Ints2List([]int{}),
		},
		{
			parameters0203{head: structures.Ints2List([]int{7, 7, 7, 7}), val: 7},
			structures.Ints2List([]int{}),
		},
	}

	for _, test := range tests {
		t.Run("Test RemoveLinkedListElements", func(t *testing.T) {
			// 完整輸入參數
			result := RemoveLinkedListElements(test.parameters0203.head, test.parameters0203.val)
			// compare 的方式需視情況調整
			if !structures.ListNodeIsEqual(result, test.ans) {
				t.Errorf("RemoveLinkedListElements(%+v) got %+v, want %+v", test.parameters0203, result, test.ans)
			}
		})
	}
}
