package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

// func 直覺寫法(headA, headB *structures.ListNode) *structures.ListNode {
// 	a := headA
// 	b := headB
// 	swapA := true
// 	swapB := true

// 	for a != nil && b != nil {
// 		// 若相同則表示找到交集了
// 		if a == b {
// 			return a
// 		}

// 		// 檢查 a 是否可以跳到 List B
// 		if a.Next == nil && swapA {
// 			a = headB
// 			swapA = false
// 		} else {
// 			a = a.Next
// 		}
// 		// 檢查 b 是否可以跳到 List A
// 		if b.Next == nil && swapB {
// 			b = headA
// 			swapB = false
// 		} else {
// 			b = b.Next
// 		}
// 	}
// 	return nil
// }

func 最佳寫法(headA, headB *structures.ListNode) *structures.ListNode {
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

func IntersectionofTwoLinkedLists(headA, headB *structures.ListNode) *structures.ListNode {
	// return 直覺寫法(headA, headB)
	return 最佳寫法(headA, headB)
}
