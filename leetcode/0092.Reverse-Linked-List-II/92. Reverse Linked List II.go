package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func ReverseLinkedListII(head *structures.ListNode, left int, right int) *structures.ListNode {
	// 設計 sentinel node 來解決當 left=1 時，head 需要被更新，但當 left!=1 head 不需要被更新的問題
	sentinelNode := &structures.ListNode{Next: head}
	currIndex := 0
	curr := sentinelNode

	// find left node
	// left node 的定義是要反轉 node 的前一個 node，不能與要反轉的 node 重疊
	// 所以如果 left=1 的話，leftNode=sentinelNode
	// rigth node 也同理，若 right=len(list)，則 rightNode=nil
	for currIndex < left-1 {
		currIndex++
		curr = curr.Next
	}
	leftNode := curr

	// move curr to the node that start reverse
	currIndex++
	curr = curr.Next
	// 紀錄 newTail node，以便在反轉完後進行連接
	newTail := curr

	// start reverse
	var prev *structures.ListNode = nil
	for currIndex <= right {
		next := curr.Next
		curr.Next = prev
		prev = curr // prev is newHead in the end
		curr = next // curr is rightNode in the end
		currIndex++
	}

	// 將反轉後的 list 與未反轉的 node 連接
	leftNode.Next = prev
	newTail.Next = curr

	return sentinelNode.Next
}
