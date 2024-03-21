package leetcode

import (
	"testing"
)

// 以下測試由 ChatGPT 生成後修改

func TestConstructor(t *testing.T) {
	linkedList := Constructor()
	if linkedList.head != nil {
		t.Errorf("Constructor Failed")
	}

	linkedList.head = &Node{val: 123, next: &Node{val: 456}}
	if linkedList.head.val != 123 {
		t.Errorf("Constructor Failed")
	}

	if linkedList.head.next.val != 456 {
		t.Errorf("Constructor Failed")
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		head     *Node
		getIndex int
		ans      int
	}{
		{
			head:     nil,
			getIndex: 5,
			ans:      -1,
		},
		{
			head:     &Node{val: 1, next: &Node{val: 2, next: &Node{val: 3}}},
			getIndex: 0,
			ans:      1,
		},
		{
			head:     &Node{val: 1, next: &Node{val: 2, next: &Node{val: 3}}},
			getIndex: 1,
			ans:      2,
		},
		{
			head:     &Node{val: 1, next: &Node{val: 2, next: &Node{val: 3}}},
			getIndex: 2,
			ans:      3,
		},
		{
			head:     &Node{val: 1, next: &Node{val: 2, next: &Node{val: 3}}},
			getIndex: 3,
			ans:      -1,
		},
		{
			head:     &Node{val: 1, next: &Node{val: 2, next: &Node{val: 3}}},
			getIndex: 100,
			ans:      -1,
		},
	}
	for _, test := range tests {
		t.Run("Test Get", func(t *testing.T) {
			linkedList := Constructor()
			linkedList.head = test.head

			result := linkedList.Get(test.getIndex)
			if result != test.ans {
				t.Errorf("Get(%+v) got %+v, want %+v", test.getIndex, result, test.ans)
			}
		})
	}
}

func TestAddAtHead(t *testing.T) {
	// 創建測試用的 MyLinkedList 實例
	myLinkedList := Constructor()

	// 測試 AddAtHead 方法
	myLinkedList.AddAtHead(1)
	// 驗證頭節點的值
	if myLinkedList.head == nil || myLinkedList.head.val != 1 {
		t.Errorf("After adding at head, expected head value to be 1, got: %v", myLinkedList.head)
	}

	// 再次測試 AddAtHead 方法
	myLinkedList.AddAtHead(2)
	// 驗證頭節點的值
	if myLinkedList.head == nil || myLinkedList.head.val != 2 {
		t.Errorf("After adding at head, expected head value to be 2, got: %v", myLinkedList.head)
	}
	// 驗證頭節點的 next 指向
	if myLinkedList.head.next == nil || myLinkedList.head.next.val != 1 {
		t.Errorf("After adding at head, expected head's next value to be 1, got: %v", myLinkedList.head.next)
	}
}

func TestAddAtTail(t *testing.T) {
	// 建立測試用的 MyLinkedList 實例
	myLinkedList := Constructor()

	// 測試 AddAtTail 方法
	myLinkedList.AddAtTail(1)
	// 驗證頭節點的值
	if myLinkedList.head == nil || myLinkedList.head.val != 1 {
		t.Errorf("After adding at tail, expected head value to be 1, got: %v", myLinkedList.head)
	}

	// 再次測試 AddAtTail 方法
	myLinkedList.AddAtTail(2)
	// 驗證頭節點的值
	if myLinkedList.head == nil || myLinkedList.head.val != 1 {
		t.Errorf("After adding at tail, expected head value to be 1, got: %v", myLinkedList.head)
	}
	// 驗證尾節點的值
	tailNode := myLinkedList.head
	for tailNode.next != nil {
		tailNode = tailNode.next
	}
	if tailNode.val != 2 {
		t.Errorf("After adding at tail, expected tail value to be 2, got: %v", tailNode.val)
	}

	// 再次測試 AddAtTail 方法
	myLinkedList.AddAtTail(3)
	// 驗證頭節點的值
	if myLinkedList.head == nil || myLinkedList.head.val != 1 {
		t.Errorf("After adding at tail, expected head value to be 1, got: %v", myLinkedList.head)
	}
	// 驗證尾節點的值
	tailNode = myLinkedList.head
	for tailNode.next != nil {
		tailNode = tailNode.next
	}
	if tailNode.val != 3 {
		t.Errorf("After adding at tail, expected tail value to be 3, got: %v", tailNode.val)
	}
}

func TestAddAtIndex(t *testing.T) {
	// 建立測試用的 MyLinkedList 實例
	myLinkedList := Constructor()

	// 測試將節點插入到空鏈表的情況
	myLinkedList.AddAtIndex(0, 1)
	// 驗證頭節點的值
	if myLinkedList.head == nil || myLinkedList.head.val != 1 {
		t.Errorf("After adding at index 0, expected head value to be 1, got: %v", myLinkedList.head)
	}

	// 測試將節點插入到鏈表末尾的情況
	myLinkedList.AddAtIndex(1, 2)
	// 驗證尾節點的值
	tailNode := myLinkedList.head
	for tailNode.next != nil {
		tailNode = tailNode.next
	}
	if tailNode.val != 2 {
		t.Errorf("After adding at index 1, expected tail value to be 2, got: %v", tailNode.val)
	}

	// 測試將節點插入到鏈表中間的情況
	myLinkedList.AddAtIndex(1, 3)
	// 驗證中間節點的值
	middleNode := myLinkedList.head
	for i := 0; i < 1; i++ {
		middleNode = middleNode.next
	}
	if middleNode.val != 3 {
		t.Errorf("After adding at index 1, expected middle value to be 3, got: %v", middleNode.val)
	}
	if middleNode.next.val != 2 {
		t.Errorf("After adding at index 1, expected tail value to be 2, got: %v", middleNode.next.val)
	}

	// 測試將節點插入到超出鏈表長度的情況
	myLinkedList.AddAtIndex(5, 4)
	// 驗證鏈表長度是否與預期相符
	length := 0
	node := myLinkedList.head
	for node != nil {
		length++
		node = node.next
	}
	if length != 3 {
		t.Errorf("After adding at index 5, expected list length to be 3, got: %d", length)
	}
}

func TestDeleteAtIndex(t *testing.T) {
	// 建立測試用的 MyLinkedList 實例
	myLinkedList := Constructor()

	// 測試從空鏈表中刪除節點的情況
	myLinkedList.DeleteAtIndex(0)
	// 驗證鏈表是否為空
	if myLinkedList.head != nil {
		t.Errorf("After deleting from empty list, expected head to be nil, got: %v", myLinkedList.head)
	}

	// 將一些節點新增到鏈表中
	myLinkedList.head = &Node{val: 1}
	myLinkedList.head.next = &Node{val: 2}
	myLinkedList.head.next.next = &Node{val: 3}
	myLinkedList.head.next.next.next = &Node{val: 4}
	myLinkedList.head.next.next.next.next = &Node{val: 5}

	// 測試刪除頭節點的情況
	// 1,2,3,4,5 => 2,3,4,5
	myLinkedList.DeleteAtIndex(0)
	// 驗證頭節點的值
	if myLinkedList.head.val != 2 {
		t.Errorf("After deleting head, expected head value to be 2, got: %v", myLinkedList.head.val)
	}

	// 測試刪除中間節點的情況
	// 2,3,4,5 => 2,4,5
	myLinkedList.DeleteAtIndex(1)
	// 驗證頭節點的值
	if myLinkedList.head.val != 2 {
		t.Errorf("After deleting at index 1, expected head value to be 2, got: %v", myLinkedList.head.val)
	}
	// 驗證中間節點的值
	if myLinkedList.head.next.val != 4 {
		t.Errorf("After deleting at index 1, expected middle value to be 4, got: %v", myLinkedList.head.next.val)
	}

	// 測試刪除尾節點的情況
	// 2,4,5 => 2,4
	myLinkedList.DeleteAtIndex(2)
	// 驗證頭節點的值
	if myLinkedList.head.val != 2 {
		t.Errorf("After deleting at index 2, expected head value to be 2, got: %v", myLinkedList.head.val)
	}
	// 驗證尾節點的值
	if myLinkedList.head.next.val != 4 {
		t.Errorf("After deleting at index 2, expected tail value to be 4, got: %v", myLinkedList.head.next.val)
	}
	// 驗證尾節點的下一個節點是否為空
	if myLinkedList.head.next.next != nil {
		t.Errorf("After deleting at index 2, expected next to be nil, got: %v", myLinkedList.head.next)
	}

	// 測試全部刪除的情況
	// 2,4 => []
	myLinkedList.DeleteAtIndex(0)
	myLinkedList.DeleteAtIndex(0)
	if myLinkedList.head != nil {
		t.Errorf("After deleting all node, expected head to be nil, got: %v", myLinkedList.head)
	}
}
