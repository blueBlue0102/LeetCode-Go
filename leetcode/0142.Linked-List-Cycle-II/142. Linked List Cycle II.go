package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func LinkedListCycleII(head *structures.ListNode) *structures.ListNode {
	if head == nil {
		return nil
	}

	hare := head
	tortoise := head

	// 試圖在環中相遇
	for doWhile := true; doWhile; doWhile = hare != tortoise {
		if hare.Next != nil && hare.Next.Next != nil {
			hare = hare.Next.Next
			tortoise = tortoise.Next
		} else {
			return nil
		}
	}

	// 成功相遇，將烏龜放回原點，兔子留在原地且每次只走一步
	// 再次相遇時，該點即為環的起點
	tortoise = head
	for hare != tortoise {
		hare = hare.Next
		tortoise = tortoise.Next
	}

	return hare
}
