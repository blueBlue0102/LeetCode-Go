package structures

import (
	"testing"
)

func TestStack(t *testing.T) {
	// 創建一個新的堆疊
	stack := NewStack()

	// 測試堆疊是否為空
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty, but it's not")
	}

	// 將一些元素推入堆疊
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// 測試堆疊的長度是否正確
	expectedLen := 3
	if stack.Len() != expectedLen {
		t.Errorf("Expected stack length to be %d, but got %d", expectedLen, stack.Len())
	}

	// 測試推入元素後是否堆疊仍然為空
	if stack.IsEmpty() {
		t.Errorf("Expected stack not to be empty, but it is")
	}

	// 測試彈出元素的順序是否正確
	expectedPopOrder := []int{3, 2, 1}
	for _, expected := range expectedPopOrder {
		popped := stack.Pop()
		if popped != expected {
			t.Errorf("Expected %d to be popped, but got %d", expected, popped)
		}
	}

	// 測試彈出所有元素後堆疊是否為空
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty after popping all elements, but it's not")
	}
}
