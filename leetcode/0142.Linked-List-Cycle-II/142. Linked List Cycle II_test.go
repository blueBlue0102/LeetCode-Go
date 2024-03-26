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

// 將節點新增至 tail
func connectNodeOnTail(head *structures.ListNode, node *structures.ListNode) *structures.ListNode {
	if head == nil {
		return node
	}

	n := head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = node
	return head
}

// createTestCase
//
// list: 要建立的 list
//
// pos: cycle 的起點 index 位置
//
// return 1: 所建立好的 list
//
// return 2: cycle 起點的 node
func createCycleTestCase(list []int, pos int) (*structures.ListNode, *structures.ListNode) {
	posNode := &structures.ListNode{Val: list[pos]}
	frontList := structures.Ints2List(list[0:pos])
	behindList := structures.Ints2List(list[pos+1 : len(list)-1])
	behindList = connectNodeOnTail(behindList, &structures.ListNode{Val: list[len(list)-1], Next: posNode})
	posNode.Next = behindList
	return connectNodeOnTail(frontList, posNode), posNode
}

func TestLinkedListCycleII(t *testing.T) {
	case1Head, case1Ans := createCycleTestCase([]int{3, 2, 0, -4}, 1)
	case2Head, case2Ans := createCycleTestCase([]int{1, 2}, 0)

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
			switch test.ans {
			case nil:
				if result != nil {
					t.Errorf("LinkedListCycleII(%+v) got %+v, want %+v", test.parameters0142, result, test.ans)
				}
			default:
				if !reflect.DeepEqual(result, test.ans) {
					t.Errorf("LinkedListCycleII(%+v) got %+v, want %+v", test.parameters0142, result, test.ans)
				}
			}
		})
	}
}
