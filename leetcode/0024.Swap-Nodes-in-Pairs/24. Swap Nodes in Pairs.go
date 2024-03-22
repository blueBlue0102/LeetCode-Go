package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func SwapNodesinPairs(head *structures.ListNode) *structures.ListNode {
	// n0(sentinelNode) -> n1(head) -> n2(n1.next) -> n3(n1.next.next)
	// swap n1 and n2
	sentinelNode := &structures.ListNode{Next: head}
	n0 := sentinelNode
	n1 := head
	for n1 != nil && n1.Next != nil {
		n3 := n1.Next.Next

		// Swapping
		n1.Next.Next = n1
		n0.Next = n1.Next
		n1.Next = n3

		// Reinitializing
		// sentinelNode -> n2 -> n1(new n0) -> n3(new n1)
		n0 = n1
		n1 = n3
	}

	return sentinelNode.Next
}
