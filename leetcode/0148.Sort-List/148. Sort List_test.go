package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0148 struct {
	head *structures.ListNode
}

func TestSortList(t *testing.T) {
	tests := []struct {
		parameters0148
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0148{structures.Ints2List([]int{4, 2, 1, 3})},
			structures.Ints2List([]int{1, 2, 3, 4}),
		},
		{
			parameters0148{structures.Ints2List([]int{-1, 5, 3, 4, 0})},
			structures.Ints2List([]int{-1, 0, 3, 4, 5}),
		},
		{
			parameters0148{structures.Ints2List([]int{})},
			structures.Ints2List([]int{}),
		},
	}

	for _, test := range tests {
		t.Run("Test SortList", func(t *testing.T) {
			// 完整輸入參數
			result := SortList(test.parameters0148.head)
			// compare 的方式需視情況調整
			if !structures.ListNodeIsEqual(result, test.ans) {
				t.Errorf("SortList(%+v) got %+v, want %+v", test.parameters0148, result, test.ans)
			}
		})
	}
}
