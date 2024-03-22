package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0024 struct {
	head *structures.ListNode
}

func TestSwapNodesinPairs(t *testing.T) {
	tests := []struct {
		parameters0024
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0024{head: structures.Ints2List([]int{1, 2, 3, 4})},
			structures.Ints2List([]int{2, 1, 4, 3}),
		},
		{
			parameters0024{head: structures.Ints2List([]int{})},
			structures.Ints2List([]int{}),
		},
		{
			parameters0024{head: structures.Ints2List([]int{1})},
			structures.Ints2List([]int{1}),
		},
		{
			parameters0024{head: structures.Ints2List([]int{1, 2, 3})},
			structures.Ints2List([]int{2, 1, 3}),
		},
		{
			parameters0024{head: structures.Ints2List([]int{1, 2, 3, 4, 5})},
			structures.Ints2List([]int{2, 1, 4, 3, 5}),
		},
	}

	for _, test := range tests {
		t.Run("Test SwapNodesinPairs", func(t *testing.T) {
			// 完整輸入參數
			result := SwapNodesinPairs(test.parameters0024.head)
			// compare 的方式需視情況調整
			if !structures.ListNodeIsEqual(result, test.ans) {
				t.Errorf("SwapNodesinPairs(%+v) got %+v, want %+v", test.parameters0024, result, test.ans)
			}
		})
	}
}
