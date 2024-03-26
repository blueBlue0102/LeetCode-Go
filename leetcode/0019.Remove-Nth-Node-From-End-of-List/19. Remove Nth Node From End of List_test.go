package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0019 struct {
	head *structures.ListNode
	n    int
}

func TestRemoveNthNodeFromEndofList(t *testing.T) {
	tests := []struct {
		parameters0019
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0019{head: structures.Ints2List([]int{1, 2, 3, 4, 5}), n: 2},
			structures.Ints2List([]int{1, 2, 3, 5}),
		},
		{
			parameters0019{head: structures.Ints2List([]int{1}), n: 1},
			structures.Ints2List([]int{}),
		},
		{
			parameters0019{head: structures.Ints2List([]int{1, 2}), n: 1},
			structures.Ints2List([]int{1}),
		},
	}

	for _, test := range tests {
		t.Run("Test RemoveNthNodeFromEndofList", func(t *testing.T) {
			// 完整輸入參數
			result := RemoveNthNodeFromEndofList(test.parameters0019.head, test.parameters0019.n)
			// compare 的方式需視情況調整
			if !structures.ListNodeIsEqual(result, test.ans) {
				t.Errorf("RemoveNthNodeFromEndofList(%+v) got %+v, want %+v", test.parameters0019, result, test.ans)
			}
		})
	}
}
