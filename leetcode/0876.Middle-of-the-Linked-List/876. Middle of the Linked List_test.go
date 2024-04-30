package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0876 struct {
	head *structures.ListNode
}

func TestMiddleoftheLinkedList(t *testing.T) {
	tests := []struct {
		parameters0876
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0876{structures.Ints2List([]int{1, 2, 3, 4, 5})},
			structures.Ints2List([]int{3, 4, 5}),
		},
		{
			parameters0876{structures.Ints2List([]int{1, 2, 3, 4, 5, 6})},
			structures.Ints2List([]int{4, 5, 6}),
		},
	}

	for _, test := range tests {
		t.Run("Test MiddleoftheLinkedList", func(t *testing.T) {
			// 完整輸入參數
			result := MiddleoftheLinkedList(test.parameters0876.head)
			// compare 的方式需視情況調整
			if !structures.ListNodeIsEqual(result, test.ans) {
				t.Errorf("MiddleoftheLinkedList(%+v) got %+v, want %+v", test.parameters0876, result, test.ans)
			}
		})
	}
}
