package leetcode

type MyLinkedList struct {
	head *Node
}

type Node struct {
	val  int
	next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (list *MyLinkedList) Get(index int) int {
	node := list.head
	count := 0
	for node != nil {
		if count == index {
			return node.val
		} else {
			count++
			node = node.next
		}
	}
	return -1
}

func (list *MyLinkedList) AddAtHead(val int) {
	node := &Node{val: val, next: list.head}
	list.head = node
}

func (list *MyLinkedList) AddAtTail(val int) {
	if list.head == nil {
		list.AddAtHead(val)
		return
	}

	node := list.head
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{val: val}
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		list.AddAtHead(val)
		return
	}
	if list.head == nil {
		return
	}
	// 目標最終要指向第 index-1 的那個 node
	node := list.head

	// 試圖將 node 指向 list[index-1]，若做不到表示 index 超過 length
	for count := 1; count != index; count++ {
		if node.next != nil {
			node = node.next
		} else {
			return
		}
	}

	// 開始新增節點
	if node.next == nil {
		node.next = &Node{val: val}
	} else {
		newNode := &Node{val: val, next: node.next}
		node.next = newNode
	}
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if list.head == nil {
		return
	}
	// 砍頭更新 list.head 即可
	if index == 0 {
		list.head = list.head.next
		return
	}

	node := list.head
	for count := 1; count < index; count++ {
		if node.next == nil {
			return
		} else {
			node = node.next
		}
	}

	// 要刪除的目標不存在
	if node.next == nil {
		return
	}

	if node.next.next == nil {
		node.next = nil
	} else {
		node.next = node.next.next
	}
}
