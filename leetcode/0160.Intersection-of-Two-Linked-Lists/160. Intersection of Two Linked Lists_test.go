package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// 填入 function input type
type parameters0160 struct {
	headA *structures.ListNode
	headB *structures.ListNode
}

func TestIntersectionofTwoLinkedLists(t *testing.T) {
	case1Ans := structures.Ints2List([]int{8, 4, 5})
	case1HeadA := structures.ConnectNodeOnTail(structures.Ints2List([]int{4, 1}), case1Ans)
	case1HeadB := structures.ConnectNodeOnTail(structures.Ints2List([]int{5, 6, 1}), case1Ans)

	case2Ans := structures.Ints2List([]int{2, 4})
	case2HeadA := structures.ConnectNodeOnTail(structures.Ints2List([]int{1, 9, 1}), case2Ans)
	case2HeadB := structures.ConnectNodeOnTail(structures.Ints2List([]int{3}), case2Ans)

	tests := []struct {
		parameters0160
		// 填入 function output type
		ans *structures.ListNode
	}{
		// 填入 test case
		{
			parameters0160{headA: case1HeadA, headB: case1HeadB},
			case1Ans,
		},
		{
			parameters0160{headA: case2HeadA, headB: case2HeadB},
			case2Ans,
		},
		{
			parameters0160{headA: structures.Ints2List([]int{2, 6, 4}), headB: structures.Ints2List([]int{1, 5})},
			nil,
		},
	}

	for _, test := range tests {
		t.Run("Test IntersectionofTwoLinkedLists", func(t *testing.T) {
			// 完整輸入參數
			result := IntersectionofTwoLinkedLists(test.parameters0160.headA, test.parameters0160.headB)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("IntersectionofTwoLinkedLists(%+v) got %+v, want %+v", test.parameters0160, result, test.ans)
			}
		})
	}
}
