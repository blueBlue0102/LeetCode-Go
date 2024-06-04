package leetcode

import "github.com/blueBlue0102/LeetCode-Go/structures"

func ConstructBinaryTreefromPreorderandInorderTraversal(preorder []int, inorder []int) *structures.TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	for i := range inorder {
		inorderMap[inorder[i]] = i
	}
	var buildTree func(preLeft int, inLeft int, inRight int) *structures.TreeNode
	buildTree = func(preLeft int, inLeft int, inRight int) *structures.TreeNode {
		rootVal := preorder[preLeft]
		root := &structures.TreeNode{Val: rootVal}
		rootIndexAtInorder := inorderMap[rootVal]
		leftNodeAmount := rootIndexAtInorder - inLeft
		rightNodeAmount := inRight - rootIndexAtInorder
		if leftNodeAmount > 0 {
			root.Left = buildTree(preLeft+1, inLeft, rootIndexAtInorder-1)
		}
		if rightNodeAmount > 0 {
			root.Right = buildTree(preLeft+leftNodeAmount+1, rootIndexAtInorder+1, inRight)
		}
		return root
	}

	return buildTree(0, 0, len(inorder)-1)
}
