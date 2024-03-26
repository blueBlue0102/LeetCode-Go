package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func RemoveNthNodeFromEndofList(head *structures.ListNode, n int) *structures.ListNode {
	sentinelNode := &structures.ListNode{Next: head}
	leftNode := sentinelNode
	rightNode := sentinelNode

	// 使 rightNode 領先 leftNode n 個 node
	for i := 0; i < n; i++ {
		rightNode = rightNode.Next
	}

	// leftNode 和 rightNode 同時前進，直到 rightNode 碰到 tail
	for rightNode.Next != nil {
		leftNode = leftNode.Next
		rightNode = rightNode.Next
	}

	// 移除節點
	leftNode.Next = leftNode.Next.Next

	return sentinelNode.Next
}
