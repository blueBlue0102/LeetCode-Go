package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visitedNode := map[int]*Node{}

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		var newNode *Node
		if visitedNode[node.Val] == nil {
			newNode = &Node{Val: node.Val, Neighbors: make([]*Node, 0, len(node.Neighbors))}
			visitedNode[node.Val] = newNode
			for _, neighbor := range node.Neighbors {
				newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor))
			}
		} else {
			newNode = visitedNode[node.Val]
		}
		return newNode
	}
	return dfs(node)
}
