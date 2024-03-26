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

// 將節點新增至 tail
func ConnectNodeOnTail(head *ListNode, node *ListNode) *ListNode {
	if head == nil {
		return node
	}

	n := head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = node
	return head
}

// createTestCase
//
// list: 要建立的 list
//
// pos: cycle 的起點 index 位置
//
// return 1: 所建立好的 list
//
// return 2: cycle 起點的 node
func CreateCycleTestCase(list []int, pos int) (*ListNode, *ListNode) {
	posNode := &ListNode{Val: list[pos]}
	frontList := Ints2List(list[0:pos])
	behindList := Ints2List(list[pos+1 : len(list)-1])
	behindList = ConnectNodeOnTail(behindList, &ListNode{Val: list[len(list)-1], Next: posNode})
	posNode.Next = behindList
	return ConnectNodeOnTail(frontList, posNode), posNode
}
