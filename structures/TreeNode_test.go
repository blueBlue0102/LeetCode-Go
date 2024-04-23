package structures

import (
	"testing"
)

func TestInts2TreeNode(t *testing.T) {
	tests := []struct {
		input    []int
		expected *TreeNode
	}{
		{
			input:    []int{1, 2, 3, NULL, 4, NULL, 5},
			expected: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}},
		},
		{
			input:    []int{},
			expected: nil,
		},
	}

	for _, test := range tests {
		actual := Ints2TreeNode(test.input)
		if !TreeNodeIsEqual(actual, test.expected) {
			t.Errorf("For input %v, expected tree %v, but got %v", test.input, test.expected, actual)
		}
	}
}

func TestTreeNodeIsEqual(t *testing.T) {
	tree1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	tree2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	tree3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}

	if !TreeNodeIsEqual(tree1, tree2) {
		t.Errorf("Expected trees to be equal, but they're not")
	}

	if TreeNodeIsEqual(tree1, tree3) {
		t.Errorf("Expected trees to be different, but they're equal")
	}
}
