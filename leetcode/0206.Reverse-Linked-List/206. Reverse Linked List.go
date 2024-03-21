package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func ReverseLinkedList(head *structures.ListNode) *structures.ListNode {
	var prev *structures.ListNode = nil
	node := head

	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	if prev == nil {
		return head
	} else {
		return prev
	}
}
