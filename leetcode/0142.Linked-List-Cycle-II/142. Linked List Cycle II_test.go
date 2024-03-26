package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0142 struct {
	head *structures.ListNode
}

func TestLinkedListCycleII(t *testing.T) {
	case1Head, case1Ans := structures.CreateCycleTestCase([]int{3, 2, 0, -4}, 1)
	case2Head, case2Ans := structures.CreateCycleTestCase([]int{1, 2}, 0)

	tests := []struct {
		parameters0142
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0142{head: case1Head},
			case1Ans,
		},
		{
			parameters0142{head: case2Head},
			case2Ans,
		},
		{
			parameters0142{head: structures.Ints2List([]int{1})},
			nil,
		},
	}

	for _, test := range tests {
		t.Run("Test LinkedListCycleII", func(t *testing.T) {
			// 完整輸入參數
			result := LinkedListCycleII(test.parameters0142.head)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("LinkedListCycleII(%+v) got %+v, want %+v", test.parameters0142, result, test.ans)
			}
		})
	}
}
