package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func LinkedListCycle(head *structures.ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
