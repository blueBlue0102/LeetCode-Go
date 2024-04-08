package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters2487 struct {
	head *structures.ListNode
}

func TestRemoveNodesFromLinkedList(t *testing.T) {
	tests := []struct {
		parameters2487
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters2487{structures.Ints2List([]int{5, 2, 13, 3, 8})},
			structures.Ints2List([]int{13, 8}),
		},
		{
			parameters2487{structures.Ints2List([]int{1, 1, 1, 1})},
			structures.Ints2List([]int{1, 1, 1, 1}),
		},
	}

	for _, test := range tests {
		t.Run("Test RemoveNodesFromLinkedList", func(t *testing.T) {
			// 完整輸入參數
			result := RemoveNodesFromLinkedList(test.parameters2487.head)
			// compare 的方式需視情況調整
			if !structures.ListNodeIsEqual(result, test.ans) {
				t.Errorf("RemoveNodesFromLinkedList(%+v) got %+v, want %+v", test.parameters2487, result, test.ans)
			}
		})
	}
}
