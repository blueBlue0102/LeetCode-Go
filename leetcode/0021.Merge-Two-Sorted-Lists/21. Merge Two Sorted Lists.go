package leetcode

import (
	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func MergeTwoSortedLists(list1 *structures.ListNode, list2 *structures.ListNode) *structures.ListNode {
	type Method string
	const (
		TwoPointer Method = "雙指針"
		Recur      Method = "遞迴"
	)

	method := Recur

	switch method {
	case TwoPointer:
		return twoPointer(list1, list2)
	case Recur:
		return recur(list1, list2)
	default:
		return twoPointer(list1, list2)
	}
}

func twoPointer(list1 *structures.ListNode, list2 *structures.ListNode) *structures.ListNode {
	// sentinel node 的作用在於可以不必考慮 head 節點該如何被建立的問題
	// 假如 list1 和 list2 都是 nil 的話，head 也會是 nil
	// 若不透過 sentinel node 的寫法，就是要在開頭初始化時額外判斷這件事

	ans := &structures.ListNode{Next: nil}
	ansTail := ans

	for list1 != nil && list2 != nil {
		node := &structures.ListNode{}
		ansTail.Next = node

		if list1.Val <= list2.Val {
			node.Val = list1.Val
			list1 = list1.Next
		} else {
			node.Val = list2.Val
			list2 = list2.Next
		}

		ansTail = ansTail.Next
	}

	if list1 == nil {
		ansTail.Next = list2
	} else if list2 == nil {
		ansTail.Next = list1
	}

	return ans.Next
}

func recur(list1 *structures.ListNode, list2 *structures.ListNode) *structures.ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = recur(list1.Next, list2)
		return list1
	} else {
		list2.Next = recur(list1, list2.Next)
		return list2
	}
}
