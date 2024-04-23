package structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 自定義的 NULL，實際上是一個 64 位元最小的負數
//
// 當使用 `Ints2TreeNode` 時，用於表示空的葉子
const NULL = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func TreeNodeIsEqual(a *TreeNode, b *TreeNode) bool {
	if b == nil && a == nil {
		return true
	}

	if b == nil || a == nil || b.Val != a.Val {
		return false
	}

	return TreeNodeIsEqual(a.Left, b.Left) && TreeNodeIsEqual(a.Right, b.Right)
}
