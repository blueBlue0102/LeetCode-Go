package leetcode

func GraphValidTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	seenNodes := make([]bool, n)
	seenCount := 0
	edgeMap := map[int][]int{}
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		if edgeMap[n1] == nil {
			edgeMap[n1] = make([]int, 0)
		}
		if edgeMap[n2] == nil {
			edgeMap[n2] = make([]int, 0)
		}
		edgeMap[n1] = append(edgeMap[n1], n2)
		edgeMap[n2] = append(edgeMap[n2], n1)
	}

	var dfs func(node int, prevNode int) bool
	dfs = func(node int, prevNode int) bool {
		if seenNodes[node] {
			return false
		}
		seenNodes[node] = true
		seenCount++
		for _, newNode := range edgeMap[node] {
			if newNode == prevNode {
				continue
			}
			if !dfs(newNode, node) {
				return false
			}
		}
		return true
	}

	if dfs(0, 0) && seenCount == n {
		return true
	}
	return false
}
