package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func ReverseLinkedList(head *structures.ListNode) *structures.ListNode {
	if false {
		return reverseLinkedList_twoPointer(head)
	} else {
		return reverseLinkedList_recursion(head)
	}
}

func reverseLinkedList_twoPointer(head *structures.ListNode) *structures.ListNode {
	var prev *structures.ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func reverseLinkedList_recursion(head *structures.ListNode) *structures.ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		node := reverseLinkedList_recursion(head.Next)
		head.Next.Next = head
		head.Next = nil
		return node
	}
}
