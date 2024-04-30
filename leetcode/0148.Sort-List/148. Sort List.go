package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func SortList(head *structures.ListNode) *structures.ListNode {
	// return BruteForce(head)
	return TopDownMergeSort(head)
}

// FAILED LEETCODE SUBMISSION
func BruteForce(head *structures.ListNode) *structures.ListNode {
	sentinelNode := &structures.ListNode{Next: head}
	target := head
	for target != nil {
		p := target
		for p != nil {
			if p.Val < target.Val {
				target.Val, p.Val = p.Val, target.Val
			}
			p = p.Next
		}
		target = target.Next
	}

	return sentinelNode.Next
}

// 取得 Linked List 的中位節點，並且會將中節點前一個節點的 next 設為 null
func MergeSort_GetMid(head *structures.ListNode) *structures.ListNode {
	prevMid, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prevMid = head
		head = head.Next
	}
	prevMid.Next = nil
	return head
}

func MergeSort_Merge(a *structures.ListNode, b *structures.ListNode) *structures.ListNode {
	result := &structures.ListNode{}
	n := result
	for a != nil && b != nil {
		if a.Val <= b.Val {
			n.Next = a
			a = a.Next
		} else {
			n.Next = b
			b = b.Next
		}
		n = n.Next
	}

	if a == nil {
		n.Next = b
	} else if b == nil {
		n.Next = a
	}

	return result.Next
}

func TopDownMergeSort(head *structures.ListNode) *structures.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := MergeSort_GetMid(head)
	left := TopDownMergeSort(head)
	right := TopDownMergeSort(mid)
	return MergeSort_Merge(left, right)
}

// func BottomUpMergeSort(head *structures.ListNode) *structures.ListNode {

// }
