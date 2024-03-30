package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func AddTwoNumbers(l1 *structures.ListNode, l2 *structures.ListNode) *structures.ListNode {
	result := &structures.ListNode{}
	curr := result
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		carry = 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		sum %= 10

		curr.Next = &structures.ListNode{Val: sum}
		curr = curr.Next
	}

	return result.Next
}
