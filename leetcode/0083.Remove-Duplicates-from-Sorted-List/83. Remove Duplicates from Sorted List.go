package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func RemoveDuplicatesfromSortedList(head *structures.ListNode) *structures.ListNode {
	for n := head; n != nil && n.Next != nil; {
		if n.Val == n.Next.Val {
			n.Next = n.Next.Next
		} else {
			n = n.Next
		}
	}
	return head
}
