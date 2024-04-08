package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0021 struct {
	list1 *structures.ListNode
	list2 *structures.ListNode
}

func TestMergeTwoSortedLists(t *testing.T) {
	tests := []struct {
		parameters0021
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0021{structures.Ints2List([]int{1, 2, 4}), structures.Ints2List([]int{1, 3, 4})},
			structures.Ints2List([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			parameters0021{structures.Ints2List([]int{}), structures.Ints2List([]int{})},
			structures.Ints2List([]int{}),
		},
		{
			parameters0021{structures.Ints2List([]int{}), structures.Ints2List([]int{0})},
			structures.Ints2List([]int{0}),
		},
	}

	for _, test := range tests {
		t.Run("Test MergeTwoSortedLists", func(t *testing.T) {
			// 完整輸入參數
			result := MergeTwoSortedLists(test.parameters0021.list1, test.parameters0021.list2)
			// compare 的方式需視情況調整
			if !structures.ListNodeIsEqual(result, test.ans) {
				t.Errorf("MergeTwoSortedLists(%+v) got %+v, want %+v", test.parameters0021, result, test.ans)
			}
		})
	}
}
