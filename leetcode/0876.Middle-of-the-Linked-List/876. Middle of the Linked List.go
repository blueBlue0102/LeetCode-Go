package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func MiddleoftheLinkedList(head *structures.ListNode) *structures.ListNode {
	p := head
	for p != nil && p.Next != nil {
		p = p.Next.Next
		head = head.Next
	}
	return head
}
