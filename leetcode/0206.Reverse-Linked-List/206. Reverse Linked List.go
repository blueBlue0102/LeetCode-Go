package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func ReverseLinkedList(head *structures.ListNode) *structures.ListNode {
	var prev *structures.ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
