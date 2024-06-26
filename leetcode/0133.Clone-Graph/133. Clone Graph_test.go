package leetcode

import (
	"testing"
)

// 將輸入的鄰接表轉換為圖的節點
func createGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	// 創建所有節點並存儲在 map 中
	nodeMap := make(map[int]*Node)
	for i := 1; i <= len(adjList); i++ {
		nodeMap[i] = &Node{Val: i}
	}

	// 根據鄰接表更新每個節點的鄰居列表
	for i, neighbors := range adjList {
		node := nodeMap[i+1]
		for _, neighbor := range neighbors {
			node.Neighbors = append(node.Neighbors, nodeMap[neighbor])
		}
	}

	// 返回圖的起始節點，這裡返回第一個節點
	return nodeMap[1]
}

// 檢查兩個節點是否相等的函數
func nodesEqual(node1, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	if len(node1.Neighbors) != len(node2.Neighbors) {
		return false
	}

	visited := make(map[*Node]*Node)
	return dfsEqual(node1, node2, visited)
}

// 使用 DFS 比較兩個節點的鄰居列表
func dfsEqual(node1, node2 *Node, visited map[*Node]*Node) bool {
	visited[node1] = node2

	for i, neighbor1 := range node1.Neighbors {
		neighbor2 := node2.Neighbors[i]
		if neighbor2 != visited[neighbor1] {
			if neighbor1.Val != neighbor2.Val || !dfsEqual(neighbor1, neighbor2, visited) {
				return false
			}
		}
	}
	return true
}

// 填入 function input type
type parameters0133 struct {
	node *Node
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		parameters0133
		// 填入 function output type
		ans *Node
	}{
		// 填入 test case
		{
			parameters0133{createGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})},
			createGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
		},
		{
			parameters0133{createGraph([][]int{{}})},
			createGraph([][]int{{}}),
		},
		{
			parameters0133{createGraph([][]int{})},
			createGraph([][]int{}),
		},
	}

	for _, test := range tests {
		t.Run("Test CloneGraph", func(t *testing.T) {
			// 完整輸入參數
			result := CloneGraph(test.parameters0133.node)
			// compare 的方式需視情況調整
			if !nodesEqual(result, test.ans) {
				t.Errorf("CloneGraph(%+v) got %+v, want %+v", test.parameters0133, result, test.ans)
			}
		})
	}
}
