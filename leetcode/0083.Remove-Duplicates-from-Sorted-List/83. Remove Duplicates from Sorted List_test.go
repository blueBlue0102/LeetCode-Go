package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0083 struct {
	head *structures.ListNode
}

func TestRemoveDuplicatesfromSortedList(t *testing.T) {
	tests := []struct {
		parameters0083
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0083{structures.Ints2List([]int{1, 1, 2})},
			structures.Ints2List([]int{1, 2}),
		},
		{
			parameters0083{structures.Ints2List([]int{1, 1, 2, 3, 3})},
			structures.Ints2List([]int{1, 2, 3}),
		},
	}

	for _, test := range tests {
		t.Run("Test RemoveDuplicatesfromSortedList", func(t *testing.T) {
			// 完整輸入參數
			result := RemoveDuplicatesfromSortedList(test.parameters0083.head)
			// compare 的方式需視情況調整
			if !structures.ListNodeIsEqual(result, test.ans) {
				t.Errorf("RemoveDuplicatesfromSortedList(%+v) got %+v, want %+v", test.parameters0083, result, test.ans)
			}
		})
	}
}
