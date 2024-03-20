package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func RemoveLinkedListElements(head *structures.ListNode, val int) *structures.ListNode {
	dummyNode := &structures.ListNode{Next: head}
	prevNode := dummyNode

	for prevNode.Next != nil {
		if prevNode.Next.Val == val {
			if prevNode.Next.Next == nil {
				prevNode.Next = nil
			} else {
				prevNode.Next = prevNode.Next.Next
			}
		} else {
			prevNode = prevNode.Next
		}
	}

	return dummyNode.Next
}
