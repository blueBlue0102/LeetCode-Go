package structures

import (
	"reflect"
	"testing"
)

func TestListNodeIsEqual(t *testing.T) {
	tests := []struct {
		input1 *ListNode
		input2 *ListNode
		want   bool
	}{
		{nil, nil, true},
		{&ListNode{}, nil, false},
		{&ListNode{Val: 1, Next: &ListNode{Val: 2}}, &ListNode{Val: 1, Next: &ListNode{Val: 2}}, true},
		{&ListNode{Val: 1, Next: &ListNode{Val: 2}}, &ListNode{Val: 1, Next: &ListNode{Val: 3}}, false},
	}

	for _, test := range tests {
		t.Run("Test ListNodeIsEqual", func(t *testing.T) {
			result := ListNodeIsEqual(test.input1, test.input2)
			if result != test.want {
				t.Errorf("ListNodeIsEqual(%v, %v) = %t, want %t", test.input1, test.input2, result, test.want)
			}
		})
	}
}

func TestInts2List(t *testing.T) {
	tests := []struct {
		input []int
		ans   *ListNode
	}{
		{
			[]int{},
			nil,
		},
		{
			[]int{1, 2, 3},
			&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			[]int{10, 666, 4, 0, 0},
			&ListNode{Val: 10, Next: &ListNode{Val: 666, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0}}}}},
		},
	}

	for _, test := range tests {
		t.Run("Test Ints2List", func(t *testing.T) {
			result := Ints2List(test.input)
			if !ListNodeIsEqual(result, test.ans) {
				t.Errorf("Ints2List 轉換失敗")
			}
		})
	}
}

func TestList2Ints(t *testing.T) {
	tests := []struct {
		input *ListNode
		ans   []int
	}{
		{
			nil,
			[]int{},
		},
		{
			&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			[]int{1, 2, 3},
		},
		{
			&ListNode{Val: 10, Next: &ListNode{Val: 666, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0}}}}},
			[]int{10, 666, 4, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run("Test List2Ints", func(t *testing.T) {
			result := List2Ints(test.input)
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("List2Ints 轉換失敗 got %+v, want %+v", result, test.ans)
			}
		})
	}
}
