package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0206 struct {
	head *structures.ListNode
}

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		parameters0206
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0206{head: structures.Ints2List([]int{1, 2, 3, 4, 5})},
			structures.Ints2List([]int{5, 4, 3, 2, 1}),
		},
		{
			parameters0206{head: structures.Ints2List([]int{1, 2})},
			structures.Ints2List([]int{2, 1}),
		},
		{
			parameters0206{head: structures.Ints2List([]int{1})},
			structures.Ints2List([]int{1}),
		},
		{
			parameters0206{head: structures.Ints2List([]int{})},
			structures.Ints2List([]int{}),
		},
	}

	for _, test := range tests {
		t.Run("Test ReverseLinkedList", func(t *testing.T) {
			// 完整輸入參數
			result := ReverseLinkedList(test.parameters0206.head)
			// compare 的方式需視情況調整
			if !structures.ListNodeIsEqual(result, test.ans) {
				t.Errorf("ReverseLinkedList(%+v) got %+v, want %+v", test.parameters0206, result, test.ans)
			}
		})
	}
}
