package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

// 檢查兩個 ListNode 是否相等
func ListNodeIsEqual(l1 *ListNode, l2 *ListNode) bool {
	// 如果都為 nil 則相等
	if l1 == nil && l2 == nil {
		return true
	}
	// 如果其中一個為 nil，則不相等
	if l1 == nil || l2 == nil {
		return false
	}
	// 如果值不相等，則不相等
	if l1.Val != l2.Val {
		return false
	}
	// 遞迴比較下一個節點
	return ListNodeIsEqual(l1.Next, l2.Next)
}

func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	dummyNode := &ListNode{}
	target := dummyNode
	for _, value := range nums {
		target.Next = &ListNode{Val: value}
		target = target.Next
	}
	return dummyNode.Next
}

func List2Ints(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
