package leetcode

import (
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/structures"
)

func TestLowestCommonAncestorofaBinarySearchTree(t *testing.T) {
	// 構建測試用的二叉搜索樹
	root := &structures.TreeNode{Val: 6}
	root.Left = &structures.TreeNode{Val: 2}
	root.Right = &structures.TreeNode{Val: 8}
	root.Left.Left = &structures.TreeNode{Val: 0}
	root.Left.Right = &structures.TreeNode{Val: 4}
	root.Right.Left = &structures.TreeNode{Val: 7}
	root.Right.Right = &structures.TreeNode{Val: 9}
	root.Left.Right.Left = &structures.TreeNode{Val: 3}
	root.Left.Right.Right = &structures.TreeNode{Val: 5}

	tests := []struct {
		p, q, expected *structures.TreeNode
	}{
		{root.Left, root.Right, root},                      // 節點 2 和 8 的 LCA 是 6
		{root.Left, root.Left.Right, root.Left},            // 節點 2 和 4 的 LCA 是 2
		{root.Left.Left, root.Left.Right.Right, root.Left}, // 節點 0 和 5 的 LCA 是 2
		{root.Right.Left, root.Right.Right, root.Right},    // 節點 7 和 9 的 LCA 是 8
	}

	for _, test := range tests {
		result := LowestCommonAncestorofaBinarySearchTree(root, test.p, test.q)
		if !structures.TreeNodeIsEqual(result, test.expected) {
			t.Errorf("For nodes %d and %d, expected LCA is %d, but got %d", test.p.Val, test.q.Val, test.expected.Val, result.Val)
		}
	}
}
