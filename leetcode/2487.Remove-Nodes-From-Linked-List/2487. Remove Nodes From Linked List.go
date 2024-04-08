package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func RemoveNodesFromLinkedList(head *structures.ListNode) *structures.ListNode {
	watchNode := head
	tail := head

	for watchNode != nil {
		p := watchNode.Next
		for {
			if p == nil {
				tail = watchNode
				watchNode = watchNode.Next
				break
			} else if p.Val > watchNode.Val {
				if watchNode == head {
					head = p
				}
				tail.Next = p
				watchNode = p
				break
			} else {
				p = p.Next
			}
		}
	}

	return head
}
